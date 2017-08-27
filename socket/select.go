package gluasocket_socket

import (
	"time"

	"github.com/yuin/gopher-lua"
)

func selectFn(l *lua.LState) int {

	// Read arguments
	recvt := l.Get(1)
	sendt := l.Get(2)
	timeout := l.Get(3)

	// Handle select(nil, nil, timeout)
	if recvt.Type() == lua.LTNil && sendt.Type() == lua.LTNil {
		timeoutVal, ok := timeout.(lua.LNumber)
		if !ok {
			l.RaiseError("Malformed timeout in call to socket.select(?,?,timeout)")
			return 0
		}
		timeoutDuration := time.Duration(timeoutVal * 1.0e9)
		time.Sleep(timeoutDuration)
		l.Push(lua.LString("timeout"))
		return 1
	}

	// TODO Handle socket select
	l.RaiseError("socket.select(recvt,sendt,timeout) not implemented yet")
	return 0
}
