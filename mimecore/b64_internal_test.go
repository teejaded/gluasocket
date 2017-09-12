package gluasocket_mimecore

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64EncodeMoepsi(t *testing.T) {
	assert := assert.New(t)
	var input, buffer bytes.Buffer

	b64encode('M', &input, &buffer)
	assert.Equal(1, input.Len())
	assert.Equal("M", input.String())
	assert.Equal(0, buffer.Len())

	b64encode('ö', &input, &buffer)
	assert.Equal(0, input.Len())
	//	assert.Equal("TcO2", input.String())
	assert.Equal(4, buffer.Len())

	b64encode('p', &input, &buffer)
	assert.Equal(1, input.Len())
	assert.Equal("p", input.String())
	assert.Equal(4, buffer.Len())

	b64encode('s', &input, &buffer)
	assert.Equal(2, input.Len())
	assert.Equal("ps", input.String())
	assert.Equal(4, buffer.Len())

	b64encode('i', &input, &buffer)
	assert.Equal(0, input.Len())
	assert.Equal("TcO2cHNp", buffer.String())
}
