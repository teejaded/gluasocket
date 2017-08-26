package gluasocket_mime

import (
	"github.com/yuin/gopher-lua"
)

func normalizeFn(l *lua.LState) int {
	l.RaiseError("mime.normalize(marker) not implemented yet") // TODO
	return 0
}
