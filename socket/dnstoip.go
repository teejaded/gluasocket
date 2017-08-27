package gluasocket_socket

import (
	"github.com/yuin/gopher-lua"
)

func dnsToIp(l *lua.LState) int {
	l.RaiseError("socket.dns.toip(address) not implemented yet")
	return 0
}
