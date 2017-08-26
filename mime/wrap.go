package gluasocket_mime

import (
	"github.com/yuin/gopher-lua"
)

func wrapFn(l *lua.LState) int {
	l.RaiseError("mime.wrap(encoding, length) not implemented yet") // TODO
	return 0
}
