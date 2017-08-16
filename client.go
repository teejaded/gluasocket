package gluasocket

import (
	"bufio"
	"bytes"
	"net"
	"time"

	"github.com/yuin/gopher-lua"
)

type Client struct {
	Conn    net.Conn
	Timeout time.Duration
	Reader  *bufio.Reader
}

var clientMethods = map[string]lua.LGFunction{
	"receive":    clientReceive,
	"settimeout": clientSetTimeout,
}

// ----------------------------------------------------------------------------

func checkClient(l *lua.LState) *Client {
	ud := l.CheckUserData(1)
	if v, ok := ud.Value.(*Client); ok {
		return v
	}
	l.ArgError(1, "client expected")
	return nil
}

// ----------------------------------------------------------------------------

func clientReceive(l *lua.LState) int {
	client := checkClient(l)
	luaPattern := l.Get(2)
	//luaPrefix := "" // l.CheckString(3)

	if luaPattern.Type() == lua.LTString {
		pattern, ok := luaPattern.(lua.LString)
		if !ok {
			l.Push(lua.LNil)
			l.Push(lua.LString("Malformed pattern argument to socket:receive(pattern,...)"))
			return 2
		}
		// Read a line of text from the socket. Line separators are not returned.
		if pattern == "*l" {
			client.Conn.SetReadDeadline(time.Now().Add(client.Timeout))
			var buf bytes.Buffer
			for {
				line, isPrefix, err := client.Reader.ReadLine()
				if err != nil {
					errstr := err.Error()
					if err, ok := err.(net.Error); ok && err.Timeout() {
						errstr = "timeout"
					}
					l.Push(lua.LNil)
					l.Push(lua.LString(errstr))
					return 2
				}
				buf.Write(line)
				if !isPrefix {
					break
				}
			}
			l.Push(lua.LString(string(buf.Bytes())))
			return 1
		}
	}

	l.RaiseError("client:receive() not implemented yet")
	return 0
}

// ----------------------------------------------------------------------------

func clientSetTimeout(l *lua.LState) int {
	client := checkClient(l)
	timeout := l.CheckNumber(2)
	client.Timeout = time.Duration(timeout * 1.0e9)
	return 0
}
