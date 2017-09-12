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

## Design

### Divergence from LuaSocket source codes

#### 1. Unit test calls to `os.exit()` replaced with `error()`

This was necessary in order to detect and report errors from a test runner. Filed [LuaSocket Issue #227](https://github.com/diegonehab/luasocket/issues/227).

#### 2. Finalized Exceptions moved to pure-Lua

LuaSocket's exception handling is based on Diego's [Finalized Exceptions](http://lua-users.org/wiki/FinalizedExceptions) whitepaper.

After struggling to port the C-based `socket.newtry()` and `socket.protect()` functions to GopherLua, an easier path emerged when I discovered the pure-Lua implementations found in the unmerged [LuaSocket Pull Request #161](https://github.com/diegonehab/luasocket/pull/161), which introduces a new `socket.except` module, and makes tiny modifications to the `socket` module in order to use it. This served the purposes of this project perfectly.
