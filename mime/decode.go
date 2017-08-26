package gluasocket_mime

import (
	"github.com/yuin/gopher-lua"
)

func decodeFn(l *lua.LState) int {
	l.RaiseError("mime.decode(encoding) not implemented yet") // TODO
	return 0
}
