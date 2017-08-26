package gluasocket_smtp

import (
	"github.com/yuin/gopher-lua"
)

func messageFn(l *lua.LState) int {
	l.RaiseError("socket.smtp.message(mesgt) not implemented yet") // TODO
	return 0
}
