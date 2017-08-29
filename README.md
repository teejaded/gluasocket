# LuaSocket for GopherLua

A native Go implementation of [LuaSocket](https://github.com/diegonehab/luasocket) for the [GopherLua](https://github.com/yuin/gopher-lua) VM.

## Using

### Loading Modules

```go
import (
	"github.com/BixData/gluasocket"
)

// Bring up a GopherLua VM
L := lua.NewState()
defer L.Close()

// Preload modules
gluasocket.Preload(L)
```

### Get system time

Run some Lua that makes use of `socket.gettime()`:

```go
L.DoString("return require 'socket'.gettime()")

// Read the returned time
lv := L.Get(-1)
retval, ok := lv.(lua.LNumber)
gettimeValue := float64(retval)
```

## Testing

The original LuaSocket Lua-based unit tests are used and wrapped in Go unit test functions. Tests that perform `os.exit()` are modified to perform `error()` instead so that errors are made detectable.
