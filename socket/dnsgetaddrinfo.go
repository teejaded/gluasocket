package gluasocket_socket

import (
	"github.com/yuin/gopher-lua"
)

func dnsGetAddrInfo(l *lua.LState) int {
	l.RaiseError("socket.dns.getaddrinfo(address) not implemented yet")
	return 0
}
