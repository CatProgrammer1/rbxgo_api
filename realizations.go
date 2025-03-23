package rbxgo_api

type Instance struct {
	Archivable                                                                                                                            bool
	Name                                                                                                                                  string
	Parent                                                                                                                                Instance_[any]
	Sandboxed                                                                                                                             bool
	AncestryChanged, AttributeChnaged, ChildAdded, ChildRemoved, DescendantAdded, DescendantRemoving, Destroying, StyledPropertiesChanged RBXScriptSignal
}

func (Instance) AddTag(tag string)                                         {}
func (Instance) ClearAllChildren()                                         {}
func (Instance) Clone() Instance_[any]                                     { return nil }
func (Instance) Destroy()                                                  {}
func (Instance) FindFirstAncestor(name string) Instance_[any]              { return nil }
func (Instance) FindFirstAncestorOfClass(className string) Instance_[any]  { return nil }
func (Instance) FindFirstAncestorWhichIsA(className string) Instance_[any] { return nil }
func (Instance) FindFirstChild(name string, recursive bool) Instance_[any] { return nil }
func (Instance) FindFirstChildOfClass(className string) Instance_[any]     { return nil }
func (Instance) FindFirstChildWhichIsA(className string, recursive bool) Instance_[any] {
	return nil
}
func (Instance) FindFirstDescendant(name string) Instance_[any]                { return nil }
func (Instance) GetAttribute(attribute string) any                             { return nil }
func (Instance) GetAttributes() map[string]any                                 { return nil }
func (Instance) GetChildren() []Instance_[any]                                 { return nil }
func (Instance) GetDebugId(scopeLenght int) string                             { return "" }
func (Instance) GetDescendants() []Instance_[any]                              { return nil }
func (Instance) GetFullName() string                                           { return "" }
func (Instance) GetStyled(name string) any                                     { return nil }
func (Instance) GetTags() []string                                             { return nil }
func (Instance) HasTag(tag string) bool                                        { return false }
func (Instance) IsAncestorOf(descendant Instance_[any]) bool                   { return false }
func (Instance) IsDescendantOf(ancestor Instance_[any]) bool                   { return false }
func (Instance) RemoveTag(tag string)                                          {}
func (Instance) SetAttribute(attribute string, value any)                      {}
func (Instance) WaitForChild(childName string, timeOut float64) Instance_[any] { return nil }
func (Instance) IsA(className string) bool                                     { return false }

/*type Object struct {
	ClassName string
}

func (Object) IsA(className string) bool {
	return false
}

func (Object) GetPropertyChangedSignal(property string) RBXScriptSignal {
	return RBXScriptSignal{}
}*/
