package rbxgo_api

type Vector3 struct {
	X, Y, Z float64
}

type Vector2 struct {
	X, Y float64
}

type Object_ interface {
	IsA(className string) bool
	GetPropertyChangedSignal(property string) RBXScriptSignal
}

type Instance_[INSTANCE any] interface {
	Object_
	AddTag(tag string)
	ClearAllChildren()
	Clone() INSTANCE
	Destroy()
	FindFirstAncestor(name string) INSTANCE
	FindFirstAncestorOfClass(className string) INSTANCE
	FindFirstAncestorWhichIsA(className string) INSTANCE
	FindFirstChild(name string, recursive bool) INSTANCE
	FindFirstChildOfClass(className string) INSTANCE
	FindFirstChildWhichIsA(className string, recursive bool) INSTANCE
	FindFirstDescendant(name string) INSTANCE
	GetAttribute(attribute string) any
	GetAttributes() map[string]any
	GetChildren() []INSTANCE
	GetDebugId(scopeLenght int) string
	GetDescendants() []INSTANCE
	GetFullName() string
	GetStyled(name string) any
	GetTags() []string
	HasTag(tag string) bool
	IsAncestorOf(descendant INSTANCE) bool
	IsDescendantOf(ancestor INSTANCE) bool
	RemoveTag(tag string)
	SetAttribute(attribute string, value any)
	WaitForChild(childName string, timeOut float64) INSTANCE
}

type BasePart interface {
	Instance_[any]
	AngularAccelerationToTorque(angAcceleration Vector3, angVelocity Vector3) Vector3
	ApplyAngularImpulse(impulse Vector3)
	ApplyImpulse(impulse Vector3)
	ApplyImpulseAtPosition(impulse Vector3, position Vector3)
	CanCollideWith(part BasePart) bool
	CanSetNetworkOwnership() (bool, string)
	GetClosestPointOnSurface(position Vector3) Vector3
	GetConnectedParts(recursive bool) []Instance_[any]
	GetJoints() []Instance_[any]
	GetMass() float64
	GetNetworkOwner() Instance_[any]
	GetNetworkOwnershipAuto() bool
	GetNoCollisionConstraints() []Instance_[any]
	GetRootPart() Instance_[any]
	GetTouchingParts() []Instance_[any]
	GetVelocityAtPosition(position Vector3) Vector3
	IsGrounded() bool
	SetNetworkOwnershipAuto()
	TorqueToAngularAcceleration(torque Vector3, angVelocity Vector3) Vector3
}

type RBXScriptSignal struct{}

func (RBXScriptSignal) Connect(function interface{}) RBXScriptConnection {
	return RBXScriptConnection{}
}

func (RBXScriptSignal) ConnectParallel(function interface{}) RBXScriptConnection {
	return RBXScriptConnection{}
}

func (RBXScriptSignal) Once(function interface{}) RBXScriptConnection {
	return RBXScriptConnection{}
}

func (RBXScriptSignal) Wait() any {
	return nil
}

type RBXScriptConnection struct {
	Connected bool
}

func (RBXScriptConnection) Disconnect() {}

type _PluginManager struct {
	Instance
}

func (_PluginManager) ExportPlace(filePath string)     {}
func (_PluginManager) ExportSelection(filePath string) {}

type GlobalSettings struct {
	Instance
}

func (GlobalSettings) GetFFlag(name string) bool       { return false }
func (GlobalSettings) GetFVariable(name string) string { return "" }

type UserSettings struct {
	Instance
}

func (UserSettings) IsUserFeatureEnabled(name string) bool { return false }
func (UserSettings) Reset()                                {}

type UserData struct{}
