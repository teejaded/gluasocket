package gluasocket_url

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/yuin/gopher-lua"
)

func parseFn(l *lua.LState) int {

	// Read arguments
	urlArg := l.Get(1)
	defaultArg := l.Get(2)

	// Parse
	var parsedUrl *url.URL
	var err error
	if urlArg.Type() == lua.LTString {
		parsedUrl, err = url.Parse(l.ToString(1))
	} else {
		parsedUrl, err = url.Parse("")
	}
	if err != nil {
		l.RaiseError(fmt.Sprintf("socket.url.parse error: %v", err))
		return 0
	}
	pathAndParams := strings.Split(parsedUrl.Path, ";")
	var path, params string
	if len(pathAndParams) > 0 {
		path = pathAndParams[0]
	}
	if len(pathAndParams) > 1 {
		params = pathAndParams[1]
	}
	var password string
	passwordIsSet := false
	if parsedUrl.User != nil {
		password, passwordIsSet = parsedUrl.User.Password()
	}

	// Prepare result. Begin with populating requested defaults
	result := l.NewTable()
	if defaultArg.Type() == lua.LTTable {
		defaultsTable := l.ToTable(2)
		defaultsTable.ForEach(func(key, value lua.LValue) {
			result.RawSet(key, value)
		})
	}

	// scheme
	if parsedUrl.Scheme != "" {
		result.RawSetString("scheme", lua.LString(parsedUrl.Scheme))
	} else {
		result.RawSetString("scheme", lua.LNil)
	}

	// authority
	if parsedUrl.User != nil {
		if parsedUrl.Host != "" {
			result.RawSetString("authority", lua.LString(parsedUrl.User.String()+"@"+parsedUrl.Host))
		} else {
			result.RawSetString("authority", lua.LNil)
		}
	} else {
		if parsedUrl.Host != "" {
			result.RawSetString("authority", lua.LString(parsedUrl.Host))
		} else {
			result.RawSetString("authority", lua.LNil)
		}
	}

	// path
	if path != "" {
		result.RawSetString("path", lua.LString(path))
	} else {
		result.RawSetString("path", lua.LNil)
	}

	// params
	if params != "" || (urlArg.Type() != lua.LTNil && strings.Contains(urlArg.String(), ";")) {
		result.RawSetString("params", lua.LString(params))
	} else {
		result.RawSetString("params", lua.LNil)
	}

	// query
	if parsedUrl.RawQuery != "" || (urlArg.Type() != lua.LTNil && strings.Contains(urlArg.String(), "?")) {
		result.RawSetString("query", lua.LString(parsedUrl.RawQuery))
	} else {
		result.RawSetString("query", lua.LNil)
	}

	// fragment
	if parsedUrl.Fragment != "" || (urlArg.Type() != lua.LTNil && strings.HasSuffix(urlArg.String(), "#")) {
		result.RawSetString("fragment", lua.LString(parsedUrl.Fragment))
	} else {
		result.RawSetString("fragment", lua.LNil)
	}

	// userinfo
	if parsedUrl.User != nil {
		result.RawSetString("userinfo", lua.LString(parsedUrl.User.String()))
	} else {
		result.RawSetString("userinfo", lua.LNil)
	}

	// host
	hostname := parsedUrl.Hostname()
	if hostname != "" {
		result.RawSetString("host", lua.LString(parsedUrl.Hostname()))
	} else {
		result.RawSetString("host", lua.LNil)
	}

	// port
	port := parsedUrl.Port()
	if port != "" {
		result.RawSetString("port", lua.LString(port))
	} else {
		result.RawSetString("port", lua.LNil)
	}

	// user
	if parsedUrl.User != nil {
		result.RawSetString("user", lua.LString(parsedUrl.User.Username()))
	} else {
		result.RawSetString("user", lua.LNil)
	}

	// password
	if passwordIsSet {
		result.RawSetString("password", lua.LString(password))
	} else {
		result.RawSetString("password", lua.LNil)
	}

	// Return result
	l.Push(result)
	return 1
}
