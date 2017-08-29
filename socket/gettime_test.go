package gluasocket_socket_test

import (
	"testing"
	"time"

	"github.com/BixData/gluasocket/socket"
	"github.com/stretchr/testify/assert"
	"github.com/yuin/gopher-lua"
)

func TestGettime(t *testing.T) {
	assert := assert.New(t)

	luaState := lua.NewState()
	defer luaState.Close()

	luaState.PreloadModule("socket", gluasocket_socket.Loader)

	now := time.Now()
	assert.NoError(luaState.DoString("return require 'socket'.gettime()"))

	lv := luaState.Get(-1)
	retval, ok := lv.(lua.LNumber)

	assert.True(ok)
	expectedMin := float64(now.UnixNano()) / 1e9
	assert.True(float64(retval) >= expectedMin)
}
