package gluasocket_mimecore_test

import (
	"testing"

	"github.com/BixData/gluasocket/mimecore"
	"github.com/stretchr/testify/assert"
	"github.com/yuin/gopher-lua"
)

func TestB64WithDiegoPassword(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("mime.core", gluasocket_mimecore.Loader)
	assert.NoError(L.DoString(`return require 'mime.core'.b64('diego:password')`))
	//	A := L.Get(-2)
	B := L.Get(-1)
	//	assert.Equal("ZGllZ286cGFzc3dvcmQ=", A.String())
	assert.Equal(lua.LTNil, B.Type())
}
