package gluasocket_socketsmtp_test

import (
	"bytes"
	"fmt"
	"net"
	"net/mail"
	"testing"
	"time"

	"github.com/mhale/smtpd"
	"github.com/stretchr/testify/assert"
	"github.com/teejaded/gluasocket"
	"github.com/yuin/gopher-lua"
)

func TestSmtpSend(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	gluasocket.Preload(L)

	rSubject := ""
	rFrom := ""
	mailHandler := func(origin net.Addr, from string, to []string, data []byte) {
		msg, _ := mail.ReadMessage(bytes.NewReader(data))
		rSubject = msg.Header.Get("Subject")
		rFrom = from
	}
	go smtpd.ListenAndServe("127.0.0.1:25252", mailHandler, "smtpTest", "")

	subject := "test subject"
	from := "sender@domain.com"
	script := fmt.Sprintf(`
	smtp = require 'socket.smtp'
	-- gopher-lua coroutine workaround
	function step(src, snk)
		local chunk = src()
		local ret, snk_err = snk(chunk, src_err)
		if chunk and ret then return 1
		else return nil, snk_err end
	end
	smtp.send{
		server="%s",
		port="%d",
		from="<%s>",
		rcpt="<r@a.b>",
		source=smtp.message{headers={subject="%s"}},
		step=step,
	}`, "127.0.0.1", 25252, from, subject)

	assert.NoError(L.DoString(script))
	time.Sleep(20 * time.Millisecond)
	assert.Equal(subject, rSubject)
	assert.Equal(from, rFrom)
}
