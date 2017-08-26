# LuaSocket for GopherLua

A native Go implementation of [LuaSocket](https://github.com/diegonehab/luasocket) for the [GopherLua](https://github.com/yuin/gopher-lua) VM.

## Using

### Loading Modules

```go
import (
	"github.com/BixData/glua-socket"
	"github.com/BixData/glua-socket/http"
	"github.com/BixData/glua-socket/mime"
	"github.com/BixData/glua-socket/url"
	"github.com/yuin/gopher-lua"
)

import (
	"github.com/BixData/glua-socket"
	"github.com/yuin/gopher-lua"
)

// Bring up a GopherLua VM
luaState := lua.NewState()
defer luaState.Close()

// Load LuaSocket modules, implemented entirely in Go
luaState.PreloadModule("socket", gluasocket.Loader)
luaState.PreloadModule("socket.http", gluasocket_http.Loader)
luaState.PreloadModule("socket.url", gluasocket_url.Loader)
luaState.PreloadModule("mime", gluasocket_mime.Loader)
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
