package gluasocket_socketcore

import (
	"github.com/yuin/gopher-lua"
)

func dnsGetHostName(l *lua.LState) int {
	l.RaiseError("socket.dns.gethostname() not implemented yet")
	return 0
}
