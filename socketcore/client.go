package gluasocket_socketcore

import (
	"bufio"
	"net"
	"time"

	"github.com/yuin/gopher-lua"
)

const (
	CLIENT_TYPENAME = "tcp{client}"
)

type Client struct {
	Conn    net.Conn
	Timeout time.Duration
	Reader  *bufio.Reader
}

var clientMethods = map[string]lua.LGFunction{
	"close":       clientCloseMethod,
	"dirty":       clientDirtyMethod,
	"getfd":       clientGetFdMethod,
	"getpeername": clientGetPeerNameMethod,
	"getsockname": clientGetSockNameMethod,
	"getstats":    clientGetStatsMethod,
	"receive":     clientReceiveMethod,
	"settimeout":  clientSetTimeoutMethod,
	"send":        clientSendMethod,
	"setoption":   clientSetOptionMethod,
	"setstats":    clientSetStatsMethod,
	"shutdown":    clientShutdownMethod,
}

// ----------------------------------------------------------------------------

func checkClient(L *lua.LState) *Client {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Master); ok {
		if v.Client != nil {
			return v.Client
		}
	}
	L.ArgError(1, "client expected")
	return nil
}
