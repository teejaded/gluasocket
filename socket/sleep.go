package gluasocket_socket

import (
	"time"

	"github.com/yuin/gopher-lua"
)

func sleepFn(l *lua.LState) int {

	// Read arguments
	timeout := l.Get(1)

	// Handle
	timeoutVal, ok := timeout.(lua.LNumber)
	if !ok {
		l.RaiseError("Malformed timeout in call to socket.sleep(time)")
		return 0
	}
	timeoutDuration := time.Duration(timeoutVal * 1.0e9)
	time.Sleep(timeoutDuration)

	return 0
}
