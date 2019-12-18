package gluasocket_socketcore

import (
	"bufio"

	"github.com/yuin/gopher-lua"
)

func masterAcceptMethod(L *lua.LState) int {
	master, ud := checkMaster(L)
	conn, err := master.Listener.Accept()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(conn)
	client := &Client{Conn: conn, Reader: reader, Timeout: master.Timeout}
	ud.Value = client
	L.SetMetatable(ud, L.GetTypeMetatable(CLIENT_TYPENAME))
	L.Push(ud)
	return 1
}
