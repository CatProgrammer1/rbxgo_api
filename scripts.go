package rbxgo_api

type Script struct {
	Instance
	Enabled    bool
	Disabled   bool
	RunContext EnumItem //Enum.RunContext
	Source     string
}

type LocalScript Script
type ModuleScript Script
