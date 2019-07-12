## [Unreleased]
### Fixed
- [PR-14](https://github.com/nubix-io/gluasocket/pull/14): `socket.core` `skip()` results were off by one; provide `settimeout()` missing return value for compatiblility with `try()`
- [PR-15](https://github.com/nubix-io/gluasocket/pull/15): smtp sending is fixed through proper swapping of metatables on connect
- Integrate [luasocket#3ee8951](https://github.com/diegonehab/luasocket/commit/3ee89515a0ef4852f64b13133c22aa7d3a322cfd#diff-4d7e24364dca5902525b2638230cb132) (19 Nov 2017): fixed URL parsing in url.lua: parse fragment after parsing username and password
- Integrate [luasocket#2d6a0f7](https://github.com/diegonehab/luasocket/commit/2d6a0f7bda9241f827a3edbfa738603c024a423b#diff-4d7e24364dca5902525b2638230cb132) (22 Nov 2017): fixed url parsing; postpone fragment parsing after authority parsing; added test cases to test/urltest.lua; fixed reference patterns in check_protect() to upper case hex letters

## [0.1.1] - 2018-01-12
### Added
- Implemented `socket.core.dns` module `getaddrinfo()`
- Implemented `socket.core` module `tcp4()` and `tcp6()`
- Implemented `socket{master}` userData object `bind()` and `listen()`

## [0.1.0] - 2017-10-14
### Added
- Fully implemented `mime.core` module in Go, which includes base64 and quoted-printable decoders & encoders
- Fully support `ltn12`, `mime`, `socket`, `socket.ftp`, `socket.headers`, `socket.smtp`, `socket.tp`, and `socket.url` modules by registering appropriate [LuaSocket](https://github.com/diegonehab/luasocket) sources
- Partially support `http` module `request()`, supporting "simple form" GET and POST, complete with SSL support
- Added experimental support of `socket` module `newtry()` and `protect()` using community [LuaSocket](https://github.com/diegonehab/luasocket) Lua sources
- Implemented `socket.core` module `connect()`, `gettime()`, `skip()`,  `sleep()`, and `tcp()` in Go
- Implemented `socket.core.dns` module `gethostname()` and `toip()` in Go
- Implemented `socket{client}` userData object `close()`, `getfd()`, `receive('*a')`, `receive('*l')`, `receive(<bytes>)`, `send()`, and `settimeout()` in Go
- Implemented `socket{master}` userData object `close()` (a no-op), `connect()`, and `settimeout()` in Go

<small>(formatted per [keepachangelog-1.1.0](http://keepachangelog.com/en/1.0.0/))</small>
