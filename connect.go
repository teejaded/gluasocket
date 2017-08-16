package gluasocket

import (
	"fmt"
	"net"

	"github.com/yuin/gopher-lua"
)

func connectFn(l *lua.LState) int {
	hostname := l.ToString(1)
	port := l.ToInt(2)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		l.RaiseError(fmt.Sprintf("Failed connecting to %s:%d: %v", hostname, port, err))
		return 0
	}

	client := &Client{Conn: conn}
	ud := l.NewUserData()
	ud.Value = client
	l.SetMetatable(ud, l.GetTypeMetatable("client"))
	l.Push(ud)
	return 1
}
