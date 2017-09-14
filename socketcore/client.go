package gluasocket_socketcore

import (
	"bufio"
	"net"
	"time"

	"github.com/yuin/gopher-lua"
)

const (
	CLIENT_TYPENAME = "tcp{client}"
)

type Client struct {
	Conn    net.Conn
	Timeout time.Duration
	Reader  *bufio.Reader
}

var clientMethods = map[string]lua.LGFunction{
	"close":      clientCloseFn,
	"receive":    clientReceiveFn,
	"settimeout": clientSetTimeoutFn,
}

// ----------------------------------------------------------------------------

func checkClient(L *lua.LState) *Client {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Client); ok {
		return v
	}
	L.ArgError(1, "client expected")
	return nil
}
