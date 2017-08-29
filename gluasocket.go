package gluasocket

import (
	"github.com/BixData/gluasocket/headers"
	"github.com/BixData/gluasocket/http"
	"github.com/BixData/gluasocket/ltn12"
	"github.com/BixData/gluasocket/mime"
	"github.com/BixData/gluasocket/mimecore"
	"github.com/BixData/gluasocket/smtp"
	"github.com/BixData/gluasocket/socket"
	"github.com/BixData/gluasocket/tp"
	"github.com/BixData/gluasocket/url"
	"github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	L.PreloadModule("ltn12", gluasocket_ltn12.Loader)
	L.PreloadModule("mime.core", gluasocket_mimecore.Loader)
	L.PreloadModule("mime", gluasocket_mime.Loader)
	L.PreloadModule("socket", gluasocket_socket.Loader)
	L.PreloadModule("socket.headers", gluasocket_headers.Loader)
	L.PreloadModule("socket.http", gluasocket_http.Loader)
	L.PreloadModule("socket.smtp", gluasocket_smtp.Loader)
	L.PreloadModule("socket.tp", gluasocket_tp.Loader)
	L.PreloadModule("socket.url", gluasocket_url.Loader)
}
