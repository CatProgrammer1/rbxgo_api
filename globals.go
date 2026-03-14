package rbxgo_api //github.com/CatProgrammer1/rbxgo_api

func Print_(a ...any)                                    {}
func Warn_(a ...any)                                     {}
func Error_(message string, level int)                   {}
func ElapsedTime_() uint                                 { return 0 }
func Tick_() float32                                     { return 0 }
func Time_() uint                                        { return 0 }
func Typeof_(v any) string                               { return "" }
func Version_() string                                   { return "" }
func PluginManager() _PluginManager                      { return _PluginManager{} }
func Settings_() _PluginManager                          { return _PluginManager{} }
func Require_(module any) any                            { return nil }
func Assert_(value any, errorMessage string)             {}
func Gcinfo_() float64                                   { return 0 }
func Pairs_(t any)                                       {}
func Ipairs_(t any)                                      {}
func Next_(t, lastKey any) (any, any)                    { return nil, nil }
func Loadstring_(contents, chunkname string) any         { return nil }
func Newproxy_(contents, chunkname string) any           { return nil }
func Pcall_(function any, args ...any) (bool, any)       { return false, nil }
func Xpcall_(function, err any, args ...any) (bool, any) { return false, nil }
func Rawequal_(v1, v2 any) bool                          { return false }
func Rawget_(t, index any) any                           { return nil }
func Rawlen_(v any) uint                                 { return 0 }
func Rawset_(t, index, value any)                        {}
func Select_(index any, args ...any) any                 { return nil }
func Setmetatable_(t, mt any) any                        { return nil }
func Tonumber_(arg any, base int) float64                { return 0 }
func Tostring(e any) string                              { return "" }
func Type_(v any) string                                 { return "" }
func Unpack_(list any, i, j int) any                     { return nil }

var (
	G       = map[any]any{}
	VERSION = ""
)
