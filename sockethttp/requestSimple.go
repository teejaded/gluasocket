package gluasocket_sockethttp

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/yuin/gopher-lua"
)

func requestSimpleFn(L *lua.LState) int {
	url := L.ToString(1)

	httpClient := http.Client{Timeout: time.Second * 15}
	res, err := httpClient.Get(url)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	L.Push(lua.LString(string(body)))
	headers := createHeadersTable(L, res.Header)
	L.Push(headers)
	L.Push(lua.LNumber(res.StatusCode))
	return 3
}

func createHeadersTable(L *lua.LState, header http.Header) *lua.LTable {
	table := L.NewTable()
	for name, value := range header {
		table.RawSetString(strings.ToLower(name), lua.LString(strings.Join(value, "\n")))
	}
	return table
}
