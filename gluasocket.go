package gluasocket

import (
	"github.com/nubix-io/gluasocket/ltn12"
	"github.com/nubix-io/gluasocket/mime"
	"github.com/nubix-io/gluasocket/mimecore"
	"github.com/nubix-io/gluasocket/socket"
	"github.com/nubix-io/gluasocket/socketcore"
	"github.com/nubix-io/gluasocket/socketexcept"
	"github.com/nubix-io/gluasocket/socketftp"
	"github.com/nubix-io/gluasocket/socketheaders"
	"github.com/nubix-io/gluasocket/sockethttp"
	"github.com/nubix-io/gluasocket/socketsmtp"
	"github.com/nubix-io/gluasocket/sockettp"
	"github.com/nubix-io/gluasocket/socketurl"
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
