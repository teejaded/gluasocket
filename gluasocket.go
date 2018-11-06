package gluasocket

import (
	"github.com/teejaded/gluasocket/ltn12"
	"github.com/teejaded/gluasocket/mime"
	"github.com/teejaded/gluasocket/mimecore"
	"github.com/teejaded/gluasocket/socket"
	"github.com/teejaded/gluasocket/socketcore"
	"github.com/teejaded/gluasocket/socketexcept"
	"github.com/teejaded/gluasocket/socketftp"
	"github.com/teejaded/gluasocket/socketheaders"
	"github.com/teejaded/gluasocket/sockethttp"
	"github.com/teejaded/gluasocket/socketsmtp"
	"github.com/teejaded/gluasocket/sockettp"
	"github.com/teejaded/gluasocket/socketurl"
	"github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	L.PreloadModule("ltn12", gluasocket_ltn12.Loader)
	L.PreloadModule("mime.core", gluasocket_mimecore.Loader)
	L.PreloadModule("mime", gluasocket_mime.Loader)
	L.PreloadModule("socket", gluasocket_socket.Loader)
	L.PreloadModule("socket.core", gluasocket_socketcore.Loader)
	L.PreloadModule("socket.except", gluasocket_socketexcept.Loader)
	L.PreloadModule("socket.ftp", gluasocket_socketftp.Loader)
	L.PreloadModule("socket.headers", gluasocket_socketheaders.Loader)
	L.PreloadModule("socket.http", gluasocket_sockethttp.Loader)
	L.PreloadModule("socket.smtp", gluasocket_socketsmtp.Loader)
	L.PreloadModule("socket.tp", gluasocket_sockettp.Loader)
	L.PreloadModule("socket.url", gluasocket_socketurl.Loader)
}
