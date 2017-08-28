package gluasocket_mimecore

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuin/gopher-lua"
)

func TestQpQuote0(t *testing.T) {
	assert := assert.New(t)
	var buffer bytes.Buffer
	qpquote('0', &buffer)
	assert.Equal("=30", buffer.String())
}

func TestQpQuote255(t *testing.T) {
	assert := assert.New(t)
	var buffer bytes.Buffer
	qpquote(0xff, &buffer)
	assert.Equal("=FF", buffer.String())
}

func TestQpEncodeMoepsi(t *testing.T) {
	assert := assert.New(t)
	qpsetup()
	marker := "\r\n"
	var input, buffer bytes.Buffer

	qpencode('M', &input, marker, &buffer)
	assert.Equal(0, input.Len())
	assert.Equal("M", buffer.String())

	qpencode('ö', &input, marker, &buffer)
	assert.Equal("M=C3=B6", buffer.String())

	qpencode('p', &input, marker, &buffer)
	assert.Equal("M=C3=B6p", buffer.String())

	qpencode('s', &input, marker, &buffer)
	assert.Equal("M=C3=B6ps", buffer.String())

	qpencode('i', &input, marker, &buffer)
	assert.Equal("M=C3=B6psi", buffer.String())
}

func TestQpFnWithMoepsi(t *testing.T) {
	assert := assert.New(t)
	qpsetup()

	luaState := lua.NewState()
	defer luaState.Close()

	luaState.Push(lua.LString("Möpsi"))
	luaState.Push(lua.LNil)
	luaState.Push(lua.LNil)
	retargs := qpFn(luaState)

	assert.Equal(2, retargs)
	encoded := luaState.ToString(-2)
	remaining := luaState.ToString(-1)
	assert.Equal("M=C3=B6psi", encoded)
	assert.Equal("", remaining)
}
