package gluasocket_smtp

import (
	"github.com/yuin/gopher-lua"
)

func sendFn(l *lua.LState) int {
	l.RaiseError("socket.smtp.send(from,rcpt,source,user,password,server,port,domain,step,create) not implemented yet") // TODO
	return 0
}
