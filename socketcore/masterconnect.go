package gluasocket_socketcore

import (
	"bufio"
	"fmt"
	"net"

	"github.com/yuin/gopher-lua"
)

func masterConnectMethod(L *lua.LState) int {
	master, ud := checkMaster(L)
	hostname := L.ToString(2)
	port := L.ToInt(3)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	reader := bufio.NewReader(conn)
	master.Client = &Client{Conn: conn, Reader: reader, Timeout: master.Timeout}
	L.SetMetatable(ud, L.GetTypeMetatable(CLIENT_TYPENAME))
	L.Push(ud)
	return 1
}
