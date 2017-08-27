package gluasocket_url

import (
	"github.com/yuin/gopher-lua"
)

const urlDotLuaBuildFn = `function luasocketBuild(parsed)
local base = _G
local url = parsed.path or ""
if parsed.params then url = url .. ";" .. parsed.params end
if parsed.query then url = url .. "?" .. parsed.query end
local authority = parsed.authority
if parsed.host then
    authority = parsed.host
    if string.find(authority, ":") then
        authority = "[" .. authority .. "]"
    end
    if parsed.port then authority = authority .. ":" .. base.tostring(parsed.port) end
    local userinfo = parsed.userinfo
    if parsed.user then
        userinfo = parsed.user
        if parsed.password then
            userinfo = userinfo .. ":" .. parsed.password
        end
    end
    if userinfo then authority = userinfo .. "@" .. authority end
end
if authority then url = "//" .. authority .. url end
if parsed.scheme then url = parsed.scheme .. ":" .. url end
if parsed.fragment then url = url .. "#" .. parsed.fragment end
return url
end
`

func buildFn(l *lua.LState) int {
	// Read arguments
	parsedUrlArg := l.Get(1)

	// Perform the operation in Lua
	if err := l.DoString(urlDotLuaBuildFn); err != nil {
		l.RaiseError("socket.url.build() error loading Lua: %v", err)
		return 0
	}
	if err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("luasocketBuild"),
		NRet:    1,
		Protect: true,
	}, parsedUrlArg); err != nil {
		l.RaiseError("socket.url.build() error: %v", err)
		return 0
	}

	// Return result
	return 1
}
