package gluasocket

import (
	"bufio"
	"fmt"
	"net"

	"github.com/yuin/gopher-lua"
)

func connectFn(l *lua.LState) int {
	hostname := l.ToString(1)
	port := l.ToInt(2)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}

	reader := bufio.NewReader(conn)
	client := &Client{Conn: conn, Reader: reader}
	ud := l.NewUserData()
	ud.Value = client
	l.SetMetatable(ud, l.GetTypeMetatable("client"))
	l.Push(ud)
	return 1
}
