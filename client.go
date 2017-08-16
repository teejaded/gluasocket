package gluasocket

import (
	"net"
	"time"

	"github.com/yuin/gopher-lua"
)

type Client struct {
	Conn    net.Conn
	Timeout time.Duration
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
	//client := checkClient(l)
	//pattern := l.CheckString(2)
	//prefix := "" // l.CheckString(3)
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
