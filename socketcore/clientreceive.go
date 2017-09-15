package gluasocket_socketcore

import (
	"bytes"
	"io"
	"net"
	"time"

	"github.com/yuin/gopher-lua"
)

func clientReceiveMethod(L *lua.LState) int {
	client := checkClient(L)
	luaPattern := L.Get(2)
	//luaPrefix := "" // TODO l.CheckString(3)

	// Read a number of bytes from the socket
	if luaPattern.Type() == lua.LTNumber {
		if client.Timeout == 0 {
			client.Conn.SetDeadline(time.Time{})
		} else {
			client.Conn.SetDeadline(time.Now().Add(client.Timeout))
		}
		var buf bytes.Buffer
		for i := 0; i < L.ToInt(2); i++ {
			byte, err := client.Reader.ReadByte()
			if err == io.EOF {
				break
			}
			if err != nil {
				errstr := err.Error()
				if err, ok := err.(net.Error); ok && err.Timeout() {
					errstr = "timeout"
				}
				L.Push(lua.LNil)
				L.Push(lua.LString(errstr))
				return 2
			}
			buf.WriteByte(byte)
		}
		L.Push(lua.LString(string(buf.Bytes())))
		return 1
	}

	// Read a line of text from the socket. Line separators are not returned.
	if luaPattern.Type() == lua.LTString && luaPattern.String() == "*l" {
		var buf bytes.Buffer
		for {
			if client.Timeout == 0 {
				client.Conn.SetDeadline(time.Time{})
			} else {
				client.Conn.SetDeadline(time.Now().Add(client.Timeout))
			}
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

	// Read until the connection is closed
	if luaPattern.Type() == lua.LTString && luaPattern.String() == "*a" {
		if client.Timeout == 0 {
			client.Conn.SetDeadline(time.Time{})
		} else {
			client.Conn.SetDeadline(time.Now().Add(client.Timeout))
		}
		var buf bytes.Buffer
		for {
			byte, err := client.Reader.ReadByte()
			if err == io.EOF {
				break
			}
			if err != nil {
				errstr := err.Error()
				if err, ok := err.(net.Error); ok && err.Timeout() {
					errstr = "timeout"
				}
				L.Push(lua.LNil)
				L.Push(lua.LString(errstr))
				return 2
			}
			buf.WriteByte(byte)
		}
		L.Push(lua.LString(string(buf.Bytes())))
		return 1
	}

	L.RaiseError("client:receive() not implemented yet")
	return 0
}
