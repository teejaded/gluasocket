package gluasocket_url

import (
	"github.com/yuin/gopher-lua"
)

const urlDotLuaSnippet = `function luasocketParse(url, default)
local base = _G
local parsed = {}
for i,v in base.pairs(default or parsed) do parsed[i] = v end
if not url or url == "" then return nil, "invalid url" end
url = string.gsub(url, "#(.*)$", function(f)
    parsed.fragment = f
    return ""
end)
url = string.gsub(url, "^([%w][%w%+%-%.]*)%:",
    function(s) parsed.scheme = s; return "" end)
url = string.gsub(url, "^//([^/]*)", function(n)
    parsed.authority = n
    return ""
end)
url = string.gsub(url, "%?(.*)", function(q)
    parsed.query = q
    return ""
end)
url = string.gsub(url, "%;(.*)", function(p)
    parsed.params = p
    return ""
end)
if url ~= "" then parsed.path = url end
local authority = parsed.authority
if not authority then return parsed end
authority = string.gsub(authority,"^([^@]*)@",
    function(u) parsed.userinfo = u; return "" end)
authority = string.gsub(authority, ":([^:%]]*)$",
    function(p) parsed.port = p; return "" end)
if authority ~= "" then 
    parsed.host = string.match(authority, "^%[(.+)%]$") or authority 
end
local userinfo = parsed.userinfo
if not userinfo then return parsed end
userinfo = string.gsub(userinfo, ":([^:]*)$",
    function(p) parsed.password = p; return "" end)
parsed.user = userinfo
return parsed
end
`

func parseFn(l *lua.LState) int {

	// Read arguments
	urlArg := l.Get(1)
	defaultArg := l.Get(2)

	// Perform the operation in Lua
	if err := l.DoString(urlDotLuaSnippet); err != nil {
		l.RaiseError("socket.url.parse() error loading Lua: %v", err)
		return 0
	}
	if err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("luasocketParse"),
		NRet:    1,
		Protect: true,
	}, urlArg, defaultArg); err != nil {
		l.RaiseError("socket.url.parse() error: %v", err)
		return 0
	}

	// Return result
	return 1
}
