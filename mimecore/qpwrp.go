package gluasocket_mimecore

import (
	"github.com/yuin/gopher-lua"
)

func qpwrpFn(l *lua.LState) int {
	l.RaiseError("mime.qpwrp(n,B,length) not implemented yet") // TODO
	return 0
}
