# LuaSocket for GopherLua

A native Go implementation of [LuaSocket](https://github.com/diegonehab/luasocket) for the [GopherLua](https://github.com/yuin/gopher-lua) VM.

## Using

### Get system time

```go
import (
	"github.com/BixData/glua-socket"
	"github.com/yuin/gopher-lua"
)

// Bring up a GopherLua VM
luaState := lua.NewState()
defer luaState.Close()

// Load this LuaSocket module, implemented entirely in Go
luaState.PreloadModule("socket", gluasocket.Loader)

// Run some Lua that makes use of gettime()
luaState.DoString("return require 'socket'.gettime()")

// Read the returned time
lv := luaState.Get(-1)
retval, ok := lv.(lua.LNumber)
gettimeValue := float64(retval)
```
