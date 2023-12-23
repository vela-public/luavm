package luavm

import "github.com/vela-public/lua"

func Callback(fn func(co *lua.LState)) func(co *lua.LState) {
	return func(co *lua.LState) {
		fn(co)
	}
}
