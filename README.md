# LuaSocket for GopherLua

A native Go implementation of [LuaSocket](https://github.com/diegonehab/luasocket) for the [GopherLua](https://github.com/yuin/gopher-lua) VM.

## Using

### Loading Modules

```go
import (
	"github.com/BixData/glua-socket/http"
	"github.com/BixData/glua-socket/ltn12"
	"github.com/BixData/glua-socket/mime"
	"github.com/BixData/glua-socket/smtp"
	"github.com/BixData/glua-socket/socket"
	"github.com/BixData/glua-socket/url"
	"github.com/yuin/gopher-lua"
)

// Bring up a GopherLua VM
luaState := lua.NewState()
defer luaState.Close()

// Preload modules
luaState.PreloadModule("ltn12", gluasocket_ltn12.Loader)
luaState.PreloadModule("mime", gluasocket_mime.Loader)
luaState.PreloadModule("socket", gluasocket_socket.Loader)
luaState.PreloadModule("socket.http", gluasocket_http.Loader)
luaState.PreloadModule("socket.smtp", gluasocket_smtp.Loader)
luaState.PreloadModule("socket.url", gluasocket_url.Loader)
```

### Get system time

Run some Lua that makes use of `socket.gettime()`:

```go
luaState.DoString("return require 'socket'.gettime()")

// Read the returned time
lv := luaState.Get(-1)
retval, ok := lv.(lua.LNumber)
gettimeValue := float64(retval)
```

## Testing

The original LuaSocket Lua-based unit tests are used and wrapped in Go unit test functions. Tests that perform `os.exit()` are modified to perform `error()` instead so that errors are made detectable.
