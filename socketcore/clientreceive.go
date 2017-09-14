package gluasocket_socketcore

import (
	"bytes"
	"net"
	"time"

	"github.com/yuin/gopher-lua"
)

func clientReceiveFn(L *lua.LState) int {
	client := checkClient(L)
	luaPattern := L.Get(2)
	//luaPrefix := "" // l.CheckString(3)

	if luaPattern.Type() == lua.LTString {
		pattern, ok := luaPattern.(lua.LString)
		if !ok {
			L.Push(lua.LNil)
			L.Push(lua.LString("Malformed pattern argument to socket:receive(pattern,...)"))
			return 2
		}
		// Read a line of text from the socket. Line separators are not returned.
		if pattern == "*l" {
			client.Conn.SetReadDeadline(time.Now().Add(client.Timeout))
			var buf bytes.Buffer
			for {
				line, isPrefix, err := client.Reader.ReadLine()
				if err != nil {
					errstr := err.Error()
					if err, ok := err.(net.Error); ok && err.Timeout() {
						errstr = "timeout"
					}
					L.Push(lua.LNil)
					L.Push(lua.LString(errstr))
					return 2
				}
				buf.Write(line)
				if !isPrefix {
					break
				}
			}
			L.Push(lua.LString(string(buf.Bytes())))
			return 1
		}
	}

	L.RaiseError("client:receive() not implemented yet")
	return 0
}
