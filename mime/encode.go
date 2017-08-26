package gluasocket_mime

import (
	"github.com/yuin/gopher-lua"
)

func encodeFn(l *lua.LState) int {
	l.RaiseError("mime.encode(encoding, mode) not implemented yet") // TODO
	return 0
}
