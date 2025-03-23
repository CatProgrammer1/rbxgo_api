package rbxgo_api

type EnumItem struct {
	Name  string
	Value int
}

type Enums struct {
	DeviceFeatureType struct {
		DeviceCapture EnumItem
	}
	BodyPart struct {
		Head     EnumItem
		Torso    EnumItem
		LeftArm  EnumItem
		RightArm EnumItem
		LeftLeg  EnumItem
		RightLeg EnumItem
	}
	ParticleEmitterShapeStyle struct {
		Volume  EnumItem
		Surface EnumItem
	}
	AudioApiRollout struct {
		Disabled  EnumItem
		Automatic EnumItem
		Enabled   EnumItem
	}
	CameraSpeedAdjustBinding struct {
		None      EnumItem
		RmbScroll EnumItem
		AltScroll EnumItem
	}
	PoseEasingStyle struct {
		Linear   EnumItem
		Constant EnumItem
		Elastic  EnumItem
		Cubic    EnumItem
		Bounce   EnumItem
		CubicV2  EnumItem
	}
	MarkupKind struct {
		PlainText EnumItem
		Markdown  EnumItem
	}
	AssetTypeVerification struct {
		Default    EnumItem
		ClientOnly EnumItem
		Always     EnumItem
	}
	DraftStatusCode struct {
		OK             EnumItem
		DraftOutdated  EnumItem
		ScriptRemoved  EnumItem
		DraftCommitted EnumItem
	}
	AudioWindowSize struct {
		Small  EnumItem
		Medium EnumItem
		Large  EnumItem
	}
	ListenerType struct {
		Camera         EnumItem
		CFrame         EnumItem
		ObjectPosition EnumItem
		ObjectCFrame   EnumItem
	}
	RunState struct {
		Stopped EnumItem
		Running EnumItem
		Paused  EnumItem
	}
	ChatVersion struct {
		LegacyChatService EnumItem
		TextChatService   EnumItem
	}
	Language struct {
		Default EnumItem
	}
	IXPLoadingStatus struct {
		None             EnumItem
		Pending          EnumItem
		Initialized      EnumItem
		ErrorInvalidUser EnumItem
		ErrorConnection  EnumItem
		ErrorJsonParse   EnumItem
		ErrorTimedOut    EnumItem
	}
	AnimationClipFromVideoStatus struct {
		Initializing          EnumItem
		Pending               EnumItem
		Processing            EnumItem
		ErrorGeneric          EnumItem
		Success               EnumItem
		ErrorVideoTooLong     EnumItem
		ErrorNoPersonDetected EnumItem
		ErrorVideoUnstable    EnumItem
		Timeout               EnumItem
		Cancelled             EnumItem
		ErrorMultiplePeople   EnumItem
		ErrorUploadingVideo   EnumItem
	}
	RunContext struct {
		Legacy EnumItem
		Server EnumItem
		Client EnumItem
		Plugin EnumItem
	}
	AvatarThumbnailCustomizationType struct {
		Closeup  EnumItem
		FullBody EnumItem
	}
	MarketplaceItemPurchaseStatus struct {
		Success                  EnumItem
		SystemError              EnumItem
		AlreadyOwned             EnumItem
		InsufficientRobux        EnumItem
		QuantityLimitExceeded    EnumItem
		QuotaExceeded            EnumItem
		NotForSale               EnumItem
		NotAvailableForPurchaser EnumItem
		PriceMismatch            EnumItem
		SoldOut                  EnumItem
		PurchaserIsSeller        EnumItem
		InsufficientMembership   EnumItem
		PlaceInvalid             EnumItem
	}
	PromptCreateAvatarResult struct {
		Success                    EnumItem
		PermissionDenied           EnumItem
		Timeout                    EnumItem
		UploadFailed               EnumItem
		NoUserInput                EnumItem
		InvalidHumanoidDescription EnumItem
		UGCValidationFailed        EnumItem
		ModeratedName              EnumItem
		MaxOutfits                 EnumItem
		PurchaseFailure            EnumItem
		UnknownFailure             EnumItem
		TokenInvalid               EnumItem
	}
	KeyInterpolationMode struct {
		Constant EnumItem
		Linear   EnumItem
		Cubic    EnumItem
	}
	KeywordFilterType struct {
		Include EnumItem
		Exclude EnumItem
	}
	InitialDockState struct {
		Top    EnumItem
		Bottom EnumItem
		Left   EnumItem
		Right  EnumItem
		Float  EnumItem
	}
	DebuggerFrameType struct {
		C   EnumItem
		Lua EnumItem
	}
	DialogTone struct {
		Neutral  EnumItem
		Friendly EnumItem
		Enemy    EnumItem
	}
	BulkMoveMode struct {
		FireAllEvents     EnumItem
		FireCFrameChanged EnumItem
	}
	ProximityPromptInputType struct {
		Keyboard EnumItem
		Gamepad  EnumItem
		Touch    EnumItem
	}
	FriendRequestEvent struct {
		Issue  EnumItem
		Revoke EnumItem
		Accept EnumItem
		Deny   EnumItem
	}
	BinType struct {
		Script   EnumItem
		GameTool EnumItem
		Grab     EnumItem
		Clone    EnumItem
		Hammer   EnumItem
	}
	EasingDirection struct {
		In    EnumItem
		Out   EnumItem
		InOut EnumItem
	}
	FrameStyle struct {
		Custom       EnumItem
		ChatBlue     EnumItem
		RobloxSquare EnumItem
		RobloxRound  EnumItem
		ChatGreen    EnumItem
		ChatRed      EnumItem
		DropShadow   EnumItem
	}
	HttpContentType struct {
		ApplicationJson       EnumItem
		ApplicationXml        EnumItem
		ApplicationUrlEncoded EnumItem
		TextPlain             EnumItem
		TextXml               EnumItem
	}
	AudioFilterType struct {
		Peak         EnumItem
		LowShelf     EnumItem
		HighShelf    EnumItem
		Lowpass12dB  EnumItem
		Lowpass24dB  EnumItem
		Lowpass48dB  EnumItem
		Highpass12dB EnumItem
		Highpass24dB EnumItem
		Highpass48dB EnumItem
		Bandpass     EnumItem
		Notch        EnumItem
	}
	CameraPanMode struct {
		Classic  EnumItem
		EdgeBump EnumItem
	}
	Material struct {
		Plastic       EnumItem
		SmoothPlastic EnumItem
		Neon          EnumItem
		Wood          EnumItem
		WoodPlanks    EnumItem
		Marble        EnumItem
		Slate         EnumItem
		Concrete      EnumItem
		Granite       EnumItem
		Brick         EnumItem
		Pebble        EnumItem
		Cobblestone   EnumItem
		Rock          EnumItem
		Sandstone     EnumItem
		Basalt        EnumItem
		CrackedLava   EnumItem
		Limestone     EnumItem
		Pavement      EnumItem
		CorrodedMetal EnumItem
		DiamondPlate  EnumItem
		Foil          EnumItem
		Metal         EnumItem
		Grass         EnumItem
		LeafyGrass    EnumItem
		Sand          EnumItem
		Fabric        EnumItem
		Snow          EnumItem
		Mud           EnumItem
		Ground        EnumItem
		Asphalt       EnumItem
		Salt          EnumItem
		Ice           EnumItem
		Glacier       EnumItem
		Glass         EnumItem
		ForceField    EnumItem
		Air           EnumItem
		Water         EnumItem
		Cardboard     EnumItem
		Carpet        EnumItem
		CeramicTiles  EnumItem
		ClayRoofTiles EnumItem
		RoofShingles  EnumItem
		Leather       EnumItem
		Plaster       EnumItem
		Rubber        EnumItem
	}
	PrivilegeType struct {
		Owner   EnumItem
		Admin   EnumItem
		Member  EnumItem
		Visitor EnumItem
		Banned  EnumItem
	}
	MessageType struct {
		MessageOutput  EnumItem
		MessageInfo    EnumItem
		MessageWarning EnumItem
		MessageError   EnumItem
	}
	DebuggerEndReason struct {
		ClientRequest          EnumItem
		Timeout                EnumItem
		InvalidHost            EnumItem
		Disconnected           EnumItem
		ServerShutdown         EnumItem
		ServerProtocolMismatch EnumItem
		ConfigurationFailed    EnumItem
		RpcError               EnumItem
	}
	GearGenreSetting struct {
		AllGenres         EnumItem
		MatchingGenreOnly EnumItem
	}
	AnimatorRetargetingMode struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	HumanoidHealthDisplayType struct {
		DisplayWhenDamaged EnumItem
		AlwaysOn           EnumItem
		AlwaysOff          EnumItem
	}
	ReverbType struct {
		NoReverb         EnumItem
		GenericReverb    EnumItem
		PaddedCell       EnumItem
		Room             EnumItem
		Bathroom         EnumItem
		LivingRoom       EnumItem
		StoneRoom        EnumItem
		Auditorium       EnumItem
		ConcertHall      EnumItem
		Cave             EnumItem
		Arena            EnumItem
		Hangar           EnumItem
		CarpettedHallway EnumItem
		Hallway          EnumItem
		StoneCorridor    EnumItem
		Alley            EnumItem
		Forest           EnumItem
		City             EnumItem
		Mountains        EnumItem
		Quarry           EnumItem
		Plain            EnumItem
		ParkingLot       EnumItem
		SewerPipe        EnumItem
		UnderWater       EnumItem
	}
	PhysicsSimulationRate struct {
		Fixed240Hz EnumItem
		Fixed120Hz EnumItem
		Fixed60Hz  EnumItem
	}
	AnalyticsCustomFieldKeys struct {
		CustomField01 EnumItem
		CustomField02 EnumItem
		CustomField03 EnumItem
	}
	ModerationStatus struct {
		ReviewedApproved EnumItem
		ReviewedRejected EnumItem
		NotReviewed      EnumItem
		NotApplicable    EnumItem
		Invalid          EnumItem
	}
	JointCreationMode struct {
		All     EnumItem
		Surface EnumItem
		None    EnumItem
	}
	AlignType struct {
		PrimaryAxisParallel      EnumItem
		PrimaryAxisPerpendicular EnumItem
		PrimaryAxisLookAt        EnumItem
		AllAxes                  EnumItem
		Parallel                 EnumItem
		Perpendicular            EnumItem
	}
	OutputLayoutMode struct {
		Horizontal EnumItem
		Vertical   EnumItem
	}
	ForceLimitMode struct {
		Magnitude EnumItem
		PerAxis   EnumItem
	}
	LineJoinMode struct {
		Round EnumItem
		Bevel EnumItem
		Miter EnumItem
	}
	DeviceType struct {
		Unknown EnumItem
		Desktop EnumItem
		Tablet  EnumItem
		Phone   EnumItem
	}
	RollOffMode struct {
		Inverse        EnumItem
		Linear         EnumItem
		LinearSquare   EnumItem
		InverseTapered EnumItem
	}
	PathStatus struct {
		Success            EnumItem
		NoPath             EnumItem
		ClosestNoPath      EnumItem
		ClosestOutOfRange  EnumItem
		FailStartNotEmpty  EnumItem
		FailFinishNotEmpty EnumItem
	}
	HapticEffectType struct {
		Custom            EnumItem
		UIHover           EnumItem
		UIClick           EnumItem
		UINotification    EnumItem
		GameplayExplosion EnumItem
		GameplayCollision EnumItem
	}
	RaycastFilterType struct {
		Exclude EnumItem
		Include EnumItem
	}
	ScreenInsets struct {
		None             EnumItem
		DeviceSafeInsets EnumItem
		CoreUISafeInsets EnumItem
		TopbarSafeInsets EnumItem
	}
	ContextActionResult struct {
		Sink EnumItem
		Pass EnumItem
	}
	Platform struct {
		Windows    EnumItem
		OSX        EnumItem
		IOS        EnumItem
		Android    EnumItem
		XBoxOne    EnumItem
		PS4        EnumItem
		PS3        EnumItem
		XBox360    EnumItem
		WiiU       EnumItem
		NX         EnumItem
		Ouya       EnumItem
		AndroidTV  EnumItem
		Chromecast EnumItem
		Linux      EnumItem
		SteamOS    EnumItem
		WebOS      EnumItem
		DOS        EnumItem
		BeOS       EnumItem
		UWP        EnumItem
		PS5        EnumItem
		MetaOS     EnumItem
		None       EnumItem
	}
	ScaleType struct {
		Stretch EnumItem
		Slice   EnumItem
		Tile    EnumItem
		Fit     EnumItem
		Crop    EnumItem
	}
	FillDirection struct {
		Horizontal EnumItem
		Vertical   EnumItem
	}
	ConnectionState struct {
		Connected    EnumItem
		Disconnected EnumItem
	}
	PromptPublishAssetResult struct {
		Success          EnumItem
		PermissionDenied EnumItem
		Timeout          EnumItem
		UploadFailed     EnumItem
		NoUserInput      EnumItem
		UnknownFailure   EnumItem
	}
	AccessModifierType struct {
		Allow EnumItem
		Deny  EnumItem
	}
	AnalyticsProgressionType struct {
		Custom   EnumItem
		Start    EnumItem
		Fail     EnumItem
		Complete EnumItem
	}
	FluidFidelity struct {
		Automatic            EnumItem
		UseCollisionGeometry EnumItem
		UsePreciseGeometry   EnumItem
	}
	DebuggerExceptionBreakMode struct {
		Never     EnumItem
		Always    EnumItem
		Unhandled EnumItem
	}
	FormFactor struct {
		Symmetric EnumItem
		Brick     EnumItem
		Plate     EnumItem
		Custom    EnumItem
	}
	AccessoryType struct {
		Unknown    EnumItem
		Hat        EnumItem
		Hair       EnumItem
		Face       EnumItem
		Neck       EnumItem
		Shoulder   EnumItem
		Front      EnumItem
		Back       EnumItem
		Waist      EnumItem
		TShirt     EnumItem
		Shirt      EnumItem
		Pants      EnumItem
		Jacket     EnumItem
		Sweater    EnumItem
		Shorts     EnumItem
		LeftShoe   EnumItem
		RightShoe  EnumItem
		DressSkirt EnumItem
		Eyebrow    EnumItem
		Eyelash    EnumItem
	}
	AdShape struct {
		HorizontalRectangle EnumItem
	}
	CloseReason struct {
		Unknown           EnumItem
		RobloxMaintenance EnumItem
		DeveloperShutdown EnumItem
		DeveloperUpdate   EnumItem
		ServerEmpty       EnumItem
		OutOfMemory       EnumItem
	}
	PermissionLevelShown struct {
		Game         EnumItem
		RobloxGame   EnumItem
		RobloxScript EnumItem
		Studio       EnumItem
		Roblox       EnumItem
	}
	MeshPartHeadsAndAccessories struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	DevTouchCameraMovementMode struct {
		UserChoice EnumItem
		Classic    EnumItem
		Follow     EnumItem
		Orbital    EnumItem
	}
	PlayerCharacterDestroyBehavior struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	AdUIEventType struct {
		AdLabelClicked          EnumItem
		VolumeButtonClicked     EnumItem
		FullscreenButtonClicked EnumItem
		PlayButtonClicked       EnumItem
		PauseButtonClicked      EnumItem
		CloseButtonClicked      EnumItem
		WhyThisAdClicked        EnumItem
		PlayEventTriggered      EnumItem
		PauseEventTriggered     EnumItem
	}
	InOut struct {
		Edge   EnumItem
		Inset  EnumItem
		Center EnumItem
	}
	PrimitiveType struct {
		Null        EnumItem
		Ball        EnumItem
		Cylinder    EnumItem
		Block       EnumItem
		Wedge       EnumItem
		CornerWedge EnumItem
	}
	CameraType struct {
		Fixed      EnumItem
		Attach     EnumItem
		Watch      EnumItem
		Track      EnumItem
		Follow     EnumItem
		Custom     EnumItem
		Scriptable EnumItem
		Orbital    EnumItem
	}
	SaveFilter struct {
		SaveWorld EnumItem
		SaveGame  EnumItem
		SaveAll   EnumItem
	}
	HttpError struct {
		OK                  EnumItem
		InvalidUrl          EnumItem
		DnsResolve          EnumItem
		ConnectFail         EnumItem
		OutOfMemory         EnumItem
		TimedOut            EnumItem
		TooManyRedirects    EnumItem
		InvalidRedirect     EnumItem
		NetFail             EnumItem
		Aborted             EnumItem
		SslConnectFail      EnumItem
		SslVerificationFail EnumItem
		Unknown             EnumItem
	}
	HumanoidDisplayDistanceType struct {
		Viewer  EnumItem
		Subject EnumItem
		None    EnumItem
	}
	AnalyticsLogLevel struct {
		Trace       EnumItem
		Debug       EnumItem
		Information EnumItem
		Warning     EnumItem
		Error       EnumItem
		Fatal       EnumItem
	}
	RenderingCacheOptimizationMode struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	PropertyStatus struct {
		Ok      EnumItem
		Warning EnumItem
		Error   EnumItem
	}
	DevTouchMovementMode struct {
		UserChoice        EnumItem
		Thumbstick        EnumItem
		DPad              EnumItem
		Thumbpad          EnumItem
		ClickToMove       EnumItem
		Scriptable        EnumItem
		DynamicThumbstick EnumItem
	}
	ChatMode struct {
		Menu        EnumItem
		TextAndMenu EnumItem
	}
	HighlightDepthMode struct {
		AlwaysOnTop EnumItem
		Occluded    EnumItem
	}
	GearType struct {
		MeleeWeapons        EnumItem
		RangedWeapons       EnumItem
		Explosives          EnumItem
		PowerUps            EnumItem
		NavigationEnhancers EnumItem
		MusicalInstruments  EnumItem
		SocialItems         EnumItem
		BuildingTools       EnumItem
		Transport           EnumItem
	}
	HttpRequestType struct {
		Default            EnumItem
		MarketplaceService EnumItem
		Players            EnumItem
		Chat               EnumItem
		Avatar             EnumItem
		Analytics          EnumItem
		Localization       EnumItem
	}
	FACSDataLod struct {
		LOD0     EnumItem
		LOD1     EnumItem
		LODCount EnumItem
	}
	AppUpdateStatus struct {
		Unknown               EnumItem
		NotSupported          EnumItem
		Failed                EnumItem
		NotAvailable          EnumItem
		Available             EnumItem
		AvailableBoundChannel EnumItem
	}
	PoseEasingDirection struct {
		In    EnumItem
		Out   EnumItem
		InOut EnumItem
	}
	AlphaMode struct {
		Overlay      EnumItem
		Transparency EnumItem
	}
	FontSize struct {
		Size8  EnumItem
		Size9  EnumItem
		Size10 EnumItem
		Size11 EnumItem
		Size12 EnumItem
		Size14 EnumItem
		Size18 EnumItem
		Size24 EnumItem
		Size36 EnumItem
		Size48 EnumItem
		Size28 EnumItem
		Size32 EnumItem
		Size42 EnumItem
		Size60 EnumItem
		Size96 EnumItem
	}
	AvatarAssetType struct {
		TShirt              EnumItem
		Hat                 EnumItem
		Shirt               EnumItem
		Pants               EnumItem
		Head                EnumItem
		Face                EnumItem
		Gear                EnumItem
		Torso               EnumItem
		RightArm            EnumItem
		LeftArm             EnumItem
		LeftLeg             EnumItem
		RightLeg            EnumItem
		HairAccessory       EnumItem
		FaceAccessory       EnumItem
		NeckAccessory       EnumItem
		ShoulderAccessory   EnumItem
		FrontAccessory      EnumItem
		BackAccessory       EnumItem
		WaistAccessory      EnumItem
		ClimbAnimation      EnumItem
		FallAnimation       EnumItem
		IdleAnimation       EnumItem
		JumpAnimation       EnumItem
		RunAnimation        EnumItem
		SwimAnimation       EnumItem
		WalkAnimation       EnumItem
		MoodAnimation       EnumItem
		EmoteAnimation      EnumItem
		TShirtAccessory     EnumItem
		ShirtAccessory      EnumItem
		PantsAccessory      EnumItem
		JacketAccessory     EnumItem
		SweaterAccessory    EnumItem
		ShortsAccessory     EnumItem
		LeftShoeAccessory   EnumItem
		RightShoeAccessory  EnumItem
		DressSkirtAccessory EnumItem
		EyebrowAccessory    EnumItem
		EyelashAccessory    EnumItem
		DynamicHead         EnumItem
	}
	PathWaypointAction struct {
		Walk   EnumItem
		Jump   EnumItem
		Custom EnumItem
	}
	ProductLocationRestriction struct {
		AvatarShop   EnumItem
		AllowedGames EnumItem
		AllGames     EnumItem
	}
	PhysicsSteppingMethod struct {
		Default  EnumItem
		Fixed    EnumItem
		Adaptive EnumItem
	}
	ElasticBehavior struct {
		WhenScrollable EnumItem
		Always         EnumItem
		Never          EnumItem
	}
	ImageCombineType struct {
		BlendSourceOver EnumItem
		Overwrite       EnumItem
		Add             EnumItem
		Multiply        EnumItem
		AlphaBlend      EnumItem
	}
	RotationOrder struct {
		XYZ EnumItem
		XZY EnumItem
		YZX EnumItem
		YXZ EnumItem
		ZXY EnumItem
		ZYX EnumItem
	}
	RestPose struct {
		Default        EnumItem
		RotationsReset EnumItem
		Custom         EnumItem
	}
	InviteState struct {
		Placed   EnumItem
		Accepted EnumItem
		Declined EnumItem
		Missed   EnumItem
	}
	ContextActionPriority struct {
		Low    EnumItem
		Medium EnumItem
		High   EnumItem
	}
	AdEventType struct {
		RewardedAdLoaded   EnumItem
		RewardedAdGrant    EnumItem
		RewardedAdUnloaded EnumItem
		VideoLoaded        EnumItem
		VideoRemoved       EnumItem
		UserCompletedVideo EnumItem
	}
	OutfitType struct {
		All         EnumItem
		Avatar      EnumItem
		DynamicHead EnumItem
	}
	MarketplaceProductType struct {
		AvatarAsset  EnumItem
		AvatarBundle EnumItem
	}
	CellMaterial struct {
		Empty       EnumItem
		Grass       EnumItem
		Sand        EnumItem
		Brick       EnumItem
		Granite     EnumItem
		Asphalt     EnumItem
		Iron        EnumItem
		Aluminum    EnumItem
		Gold        EnumItem
		WoodPlank   EnumItem
		WoodLog     EnumItem
		Gravel      EnumItem
		CinderBlock EnumItem
		MossyStone  EnumItem
		Cement      EnumItem
		RedPlastic  EnumItem
		BluePlastic EnumItem
		Water       EnumItem
	}
	GamepadType struct {
		Unknown EnumItem
		PS4     EnumItem
		PS5     EnumItem
		XboxOne EnumItem
	}
	ScrollBarInset struct {
		None      EnumItem
		ScrollBar EnumItem
		Always    EnumItem
	}
	CreatorType struct {
		User  EnumItem
		Group EnumItem
	}
	ProximityPromptStyle struct {
		Default EnumItem
		Custom  EnumItem
	}
	AspectType struct {
		FitWithinMaxSize    EnumItem
		ScaleWithParentSize EnumItem
	}
	CenterDialogType struct {
		UnsolicitedDialog     EnumItem
		PlayerInitiatedDialog EnumItem
		ModalDialog           EnumItem
		QuitDialog            EnumItem
	}
	DragDetectorResponseStyle struct {
		Geometric EnumItem
		Physical  EnumItem
		Custom    EnumItem
	}
	FacialAnimationStreamingState struct {
		None   EnumItem
		Audio  EnumItem
		Video  EnumItem
		Place  EnumItem
		Server EnumItem
	}
	ControlMode struct {
		Classic         EnumItem
		MouseLockSwitch EnumItem
	}
	CompletionItemKind struct {
		Text          EnumItem
		Method        EnumItem
		Function      EnumItem
		Constructor   EnumItem
		Field         EnumItem
		Variable      EnumItem
		Class         EnumItem
		Interface     EnumItem
		Module        EnumItem
		Property      EnumItem
		Unit          EnumItem
		Value         EnumItem
		Enum          EnumItem
		Keyword       EnumItem
		Snippet       EnumItem
		Color         EnumItem
		File          EnumItem
		Reference     EnumItem
		Folder        EnumItem
		EnumMember    EnumItem
		Constant      EnumItem
		Struct        EnumItem
		Event         EnumItem
		Operator      EnumItem
		TypeParameter EnumItem
	}
	PrimalPhysicsSolver struct {
		Default      EnumItem
		Experimental EnumItem
		Disabled     EnumItem
	}
	LiveEditingBroadcastMessageType struct {
		Normal  EnumItem
		Warning EnumItem
		Error   EnumItem
	}
	FunctionalTestResult struct {
		Passed  EnumItem
		Warning EnumItem
		Error   EnumItem
	}
	PlayerChatType struct {
		All     EnumItem
		Team    EnumItem
		Whisper EnumItem
	}
	DebuggerPauseReason struct {
		Unknown    EnumItem
		Requested  EnumItem
		Breakpoint EnumItem
		Exception  EnumItem
		SingleStep EnumItem
		Entrypoint EnumItem
	}
	ItemLineAlignment struct {
		Automatic EnumItem
		Start     EnumItem
		Center    EnumItem
		End       EnumItem
		Stretch   EnumItem
	}
	SavedQualitySetting struct {
		Automatic      EnumItem
		QualityLevel1  EnumItem
		QualityLevel2  EnumItem
		QualityLevel3  EnumItem
		QualityLevel4  EnumItem
		QualityLevel5  EnumItem
		QualityLevel6  EnumItem
		QualityLevel7  EnumItem
		QualityLevel8  EnumItem
		QualityLevel9  EnumItem
		QualityLevel10 EnumItem
	}
	MoveState struct {
		Stopped  EnumItem
		Coasting EnumItem
		Pushing  EnumItem
		Stopping EnumItem
		AirFree  EnumItem
	}
	AdUIType struct {
		None  EnumItem
		Image EnumItem
		Video EnumItem
	}
	EnviromentalPhysicsThrottle struct {
		DefaultAuto EnumItem
		Disabled    EnumItem
		Always      EnumItem
		Skip2       EnumItem
		Skip4       EnumItem
		Skip8       EnumItem
		Skip16      EnumItem
	}
	EasingStyle struct {
		Linear      EnumItem
		Sine        EnumItem
		Back        EnumItem
		Quad        EnumItem
		Quart       EnumItem
		Quint       EnumItem
		Bounce      EnumItem
		Elastic     EnumItem
		Exponential EnumItem
		Circular    EnumItem
		Cubic       EnumItem
	}
	DeveloperMemoryTag struct {
		Internal                 EnumItem
		HttpCache                EnumItem
		Instances                EnumItem
		Signals                  EnumItem
		LuaHeap                  EnumItem
		Script                   EnumItem
		PhysicsCollision         EnumItem
		PhysicsParts             EnumItem
		GraphicsSolidModels      EnumItem
		GraphicsMeshParts        EnumItem
		GraphicsParticles        EnumItem
		GraphicsParts            EnumItem
		GraphicsSpatialHash      EnumItem
		GraphicsTerrain          EnumItem
		GraphicsTexture          EnumItem
		GraphicsTextureCharacter EnumItem
		Sounds                   EnumItem
		StreamingSounds          EnumItem
		TerrainVoxels            EnumItem
		Gui                      EnumItem
		Animation                EnumItem
		Navigation               EnumItem
		GeometryCSG              EnumItem
	}
	NetworkOwnership struct {
		Automatic EnumItem
		Manual    EnumItem
		OnContact EnumItem
	}
	Button struct {
		Jump     EnumItem
		Dismount EnumItem
	}
	AssetFetchStatus struct {
		Success  EnumItem
		Failure  EnumItem
		None     EnumItem
		Loading  EnumItem
		TimedOut EnumItem
	}
	InterpolationThrottlingMode struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	ActuatorRelativeTo struct {
		Attachment0 EnumItem
		Attachment1 EnumItem
		World       EnumItem
	}
	ActionType struct {
		Nothing EnumItem
		Pause   EnumItem
		Lose    EnumItem
		Draw    EnumItem
		Win     EnumItem
	}
	CatalogSortAggregation struct {
		Past12Hours EnumItem
		PastDay     EnumItem
		Past3Days   EnumItem
		PastWeek    EnumItem
		PastMonth   EnumItem
		AllTime     EnumItem
	}
	CompileTarget struct {
		Client        EnumItem
		CoreScript    EnumItem
		Studio        EnumItem
		CoreScriptRaw EnumItem
	}
	PromptCreateAssetResult struct {
		Success          EnumItem
		PermissionDenied EnumItem
		Timeout          EnumItem
		UploadFailed     EnumItem
		NoUserInput      EnumItem
		UnknownFailure   EnumItem
	}
	LeftRight struct {
		Left   EnumItem
		Center EnumItem
		Right  EnumItem
	}
	KeyCode struct {
		Unknown           EnumItem
		Backspace         EnumItem
		Tab               EnumItem
		Clear             EnumItem
		Return            EnumItem
		Pause             EnumItem
		Escape            EnumItem
		Space             EnumItem
		QuotedDouble      EnumItem
		Hash              EnumItem
		Dollar            EnumItem
		Percent           EnumItem
		Ampersand         EnumItem
		Quote             EnumItem
		LeftParenthesis   EnumItem
		RightParenthesis  EnumItem
		Asterisk          EnumItem
		Plus              EnumItem
		Comma             EnumItem
		Minus             EnumItem
		Period            EnumItem
		Slash             EnumItem
		Zero              EnumItem
		One               EnumItem
		Two               EnumItem
		Three             EnumItem
		Four              EnumItem
		Five              EnumItem
		Six               EnumItem
		Seven             EnumItem
		Eight             EnumItem
		Nine              EnumItem
		Colon             EnumItem
		Semicolon         EnumItem
		LessThan          EnumItem
		Equals            EnumItem
		GreaterThan       EnumItem
		Question          EnumItem
		At                EnumItem
		LeftBracket       EnumItem
		BackSlash         EnumItem
		RightBracket      EnumItem
		Caret             EnumItem
		Underscore        EnumItem
		Backquote         EnumItem
		A                 EnumItem
		B                 EnumItem
		C                 EnumItem
		D                 EnumItem
		E                 EnumItem
		F                 EnumItem
		G                 EnumItem
		H                 EnumItem
		I                 EnumItem
		J                 EnumItem
		K                 EnumItem
		L                 EnumItem
		M                 EnumItem
		N                 EnumItem
		O                 EnumItem
		P                 EnumItem
		Q                 EnumItem
		R                 EnumItem
		S                 EnumItem
		T                 EnumItem
		U                 EnumItem
		V                 EnumItem
		W                 EnumItem
		X                 EnumItem
		Y                 EnumItem
		Z                 EnumItem
		LeftCurly         EnumItem
		Pipe              EnumItem
		RightCurly        EnumItem
		Tilde             EnumItem
		Delete            EnumItem
		World0            EnumItem
		World1            EnumItem
		World2            EnumItem
		World3            EnumItem
		World4            EnumItem
		World5            EnumItem
		World6            EnumItem
		World7            EnumItem
		World8            EnumItem
		World9            EnumItem
		World10           EnumItem
		World11           EnumItem
		World12           EnumItem
		World13           EnumItem
		World14           EnumItem
		World15           EnumItem
		World16           EnumItem
		World17           EnumItem
		World18           EnumItem
		World19           EnumItem
		World20           EnumItem
		World21           EnumItem
		World22           EnumItem
		World23           EnumItem
		World24           EnumItem
		World25           EnumItem
		World26           EnumItem
		World27           EnumItem
		World28           EnumItem
		World29           EnumItem
		World30           EnumItem
		World31           EnumItem
		World32           EnumItem
		World33           EnumItem
		World34           EnumItem
		World35           EnumItem
		World36           EnumItem
		World37           EnumItem
		World38           EnumItem
		World39           EnumItem
		World40           EnumItem
		World41           EnumItem
		World42           EnumItem
		World43           EnumItem
		World44           EnumItem
		World45           EnumItem
		World46           EnumItem
		World47           EnumItem
		World48           EnumItem
		World49           EnumItem
		World50           EnumItem
		World51           EnumItem
		World52           EnumItem
		World53           EnumItem
		World54           EnumItem
		World55           EnumItem
		World56           EnumItem
		World57           EnumItem
		World58           EnumItem
		World59           EnumItem
		World60           EnumItem
		World61           EnumItem
		World62           EnumItem
		World63           EnumItem
		World64           EnumItem
		World65           EnumItem
		World66           EnumItem
		World67           EnumItem
		World68           EnumItem
		World69           EnumItem
		World70           EnumItem
		World71           EnumItem
		World72           EnumItem
		World73           EnumItem
		World74           EnumItem
		World75           EnumItem
		World76           EnumItem
		World77           EnumItem
		World78           EnumItem
		World79           EnumItem
		World80           EnumItem
		World81           EnumItem
		World82           EnumItem
		World83           EnumItem
		World84           EnumItem
		World85           EnumItem
		World86           EnumItem
		World87           EnumItem
		World88           EnumItem
		World89           EnumItem
		World90           EnumItem
		World91           EnumItem
		World92           EnumItem
		World93           EnumItem
		World94           EnumItem
		World95           EnumItem
		KeypadZero        EnumItem
		KeypadOne         EnumItem
		KeypadTwo         EnumItem
		KeypadThree       EnumItem
		KeypadFour        EnumItem
		KeypadFive        EnumItem
		KeypadSix         EnumItem
		KeypadSeven       EnumItem
		KeypadEight       EnumItem
		KeypadNine        EnumItem
		KeypadPeriod      EnumItem
		KeypadDivide      EnumItem
		KeypadMultiply    EnumItem
		KeypadMinus       EnumItem
		KeypadPlus        EnumItem
		KeypadEnter       EnumItem
		KeypadEquals      EnumItem
		Up                EnumItem
		Down              EnumItem
		Right             EnumItem
		Left              EnumItem
		Insert            EnumItem
		Home              EnumItem
		End               EnumItem
		PageUp            EnumItem
		PageDown          EnumItem
		F1                EnumItem
		F2                EnumItem
		F3                EnumItem
		F4                EnumItem
		F5                EnumItem
		F6                EnumItem
		F7                EnumItem
		F8                EnumItem
		F9                EnumItem
		F10               EnumItem
		F11               EnumItem
		F12               EnumItem
		F13               EnumItem
		F14               EnumItem
		F15               EnumItem
		NumLock           EnumItem
		CapsLock          EnumItem
		ScrollLock        EnumItem
		RightShift        EnumItem
		LeftShift         EnumItem
		RightControl      EnumItem
		LeftControl       EnumItem
		RightAlt          EnumItem
		LeftAlt           EnumItem
		RightMeta         EnumItem
		LeftMeta          EnumItem
		LeftSuper         EnumItem
		RightSuper        EnumItem
		Mode              EnumItem
		Compose           EnumItem
		Help              EnumItem
		Print             EnumItem
		SysReq            EnumItem
		Break             EnumItem
		Menu              EnumItem
		Power             EnumItem
		Euro              EnumItem
		Undo              EnumItem
		ButtonX           EnumItem
		ButtonY           EnumItem
		ButtonA           EnumItem
		ButtonB           EnumItem
		ButtonR1          EnumItem
		ButtonL1          EnumItem
		ButtonR2          EnumItem
		ButtonL2          EnumItem
		ButtonR3          EnumItem
		ButtonL3          EnumItem
		ButtonStart       EnumItem
		ButtonSelect      EnumItem
		DPadLeft          EnumItem
		DPadRight         EnumItem
		DPadUp            EnumItem
		DPadDown          EnumItem
		Thumbstick1       EnumItem
		Thumbstick2       EnumItem
		MouseLeftButton   EnumItem
		MouseRightButton  EnumItem
		MouseMiddleButton EnumItem
		MouseBackButton   EnumItem
		MouseNoButton     EnumItem
		MouseX            EnumItem
		MouseY            EnumItem
	}
	LoadCharacterLayeredClothing struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	QualityLevel struct {
		Automatic EnumItem
		Level01   EnumItem
		Level02   EnumItem
		Level03   EnumItem
		Level04   EnumItem
		Level05   EnumItem
		Level06   EnumItem
		Level07   EnumItem
		Level08   EnumItem
		Level09   EnumItem
		Level10   EnumItem
		Level11   EnumItem
		Level12   EnumItem
		Level13   EnumItem
		Level14   EnumItem
		Level15   EnumItem
		Level16   EnumItem
		Level17   EnumItem
		Level18   EnumItem
		Level19   EnumItem
		Level20   EnumItem
		Level21   EnumItem
	}
	FontWeight struct {
		Thin       EnumItem
		ExtraLight EnumItem
		Light      EnumItem
		Regular    EnumItem
		Medium     EnumItem
		SemiBold   EnumItem
		Bold       EnumItem
		ExtraBold  EnumItem
		Heavy      EnumItem
	}
	ScrollingDirection struct {
		X  EnumItem
		Y  EnumItem
		XY EnumItem
	}
	ParticleFlipbookMode struct {
		Loop     EnumItem
		OneShot  EnumItem
		PingPong EnumItem
		Random   EnumItem
	}
	CurrencyType struct {
		Default EnumItem
		Robux   EnumItem
		Tix     EnumItem
	}
	ReplicateInstanceDestroySetting struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	FramerateManagerMode struct {
		Automatic EnumItem
		On        EnumItem
		Off       EnumItem
	}
	DragDetectorPermissionPolicy struct {
		Nobody     EnumItem
		Everybody  EnumItem
		Scriptable EnumItem
	}
	AvatarUnificationMode struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	HorizontalAlignment struct {
		Center EnumItem
		Left   EnumItem
		Right  EnumItem
	}
	IKCollisionsMode struct {
		NoCollisions               EnumItem
		OtherMechanismsAnchored    EnumItem
		IncludeContactedMechanisms EnumItem
	}
	CollaboratorStatus struct {
		None             EnumItem
		Editing3D        EnumItem
		Scripting        EnumItem
		PrivateScripting EnumItem
	}
	CatalogCategoryFilter struct {
		None               EnumItem
		Featured           EnumItem
		Collectibles       EnumItem
		CommunityCreations EnumItem
		Premium            EnumItem
		Recommended        EnumItem
	}
	ResamplerMode struct {
		Default   EnumItem
		Pixelated EnumItem
	}
	AssetCreatorType struct {
		User  EnumItem
		Group EnumItem
	}
	CageType struct {
		Inner EnumItem
		Outer EnumItem
	}
	FontStyle struct {
		Normal EnumItem
		Italic EnumItem
	}
	CameraMode struct {
		Classic         EnumItem
		LockFirstPerson EnumItem
	}
	AudioSubType struct {
		Music       EnumItem
		SoundEffect EnumItem
	}
	BreakReason struct {
		Other             EnumItem
		Error             EnumItem
		SpecialBreakpoint EnumItem
		UserBreakpoint    EnumItem
	}
	MeshScaleUnit struct {
		Stud  EnumItem
		Meter EnumItem
		CM    EnumItem
		MM    EnumItem
		Foot  EnumItem
		Inch  EnumItem
	}
	RigScale struct {
		Default     EnumItem
		Rthro       EnumItem
		RthroNarrow EnumItem
	}
	MarketplaceBulkPurchasePromptStatus struct {
		Completed EnumItem
		Aborted   EnumItem
		Error     EnumItem
	}
	FilterResult struct {
		Accepted EnumItem
		Rejected EnumItem
	}
	AppLifecycleManagerState struct {
		Detached EnumItem
		Active   EnumItem
		Inactive EnumItem
		Hidden   EnumItem
	}
	ParticleEmitterShapeInOut struct {
		Outward  EnumItem
		Inward   EnumItem
		InAndOut EnumItem
	}
	CreateOutfitFailure struct {
		InvalidName        EnumItem
		OutfitLimitReached EnumItem
		Other              EnumItem
	}
	NetworkStatus struct {
		Unknown      EnumItem
		Connected    EnumItem
		Disconnected EnumItem
	}
	CellOrientation struct {
		NegZ EnumItem
		X    EnumItem
		Z    EnumItem
		NegX EnumItem
	}
	AnalyticsEconomyFlowType struct {
		Sink   EnumItem
		Source EnumItem
	}
	FieldOfViewMode struct {
		Vertical EnumItem
		Diagonal EnumItem
		MaxAxis  EnumItem
	}
	ApplyStrokeMode struct {
		Contextual EnumItem
		Border     EnumItem
	}
	NameOcclusion struct {
		NoOcclusion    EnumItem
		EnemyOcclusion EnumItem
		OccludeAll     EnumItem
	}
	PlayerActions struct {
		CharacterForward  EnumItem
		CharacterBackward EnumItem
		CharacterLeft     EnumItem
		CharacterRight    EnumItem
		CharacterJump     EnumItem
	}
	Font struct {
		Legacy               EnumItem
		Arial                EnumItem
		ArialBold            EnumItem
		SourceSans           EnumItem
		SourceSansBold       EnumItem
		SourceSansLight      EnumItem
		SourceSansItalic     EnumItem
		Bodoni               EnumItem
		Garamond             EnumItem
		Cartoon              EnumItem
		Code                 EnumItem
		Highway              EnumItem
		SciFi                EnumItem
		Arcade               EnumItem
		Fantasy              EnumItem
		Antique              EnumItem
		SourceSansSemibold   EnumItem
		Gotham               EnumItem
		GothamMedium         EnumItem
		GothamBold           EnumItem
		GothamBlack          EnumItem
		AmaticSC             EnumItem
		Bangers              EnumItem
		Creepster            EnumItem
		DenkOne              EnumItem
		Fondamento           EnumItem
		FredokaOne           EnumItem
		GrenzeGotisch        EnumItem
		IndieFlower          EnumItem
		JosefinSans          EnumItem
		Jura                 EnumItem
		Kalam                EnumItem
		LuckiestGuy          EnumItem
		Merriweather         EnumItem
		Michroma             EnumItem
		Nunito               EnumItem
		Oswald               EnumItem
		PatrickHand          EnumItem
		PermanentMarker      EnumItem
		Roboto               EnumItem
		RobotoCondensed      EnumItem
		RobotoMono           EnumItem
		Sarpanch             EnumItem
		SpecialElite         EnumItem
		TitilliumWeb         EnumItem
		Ubuntu               EnumItem
		BuilderSans          EnumItem
		BuilderSansMedium    EnumItem
		BuilderSansBold      EnumItem
		BuilderSansExtraBold EnumItem
		Arimo                EnumItem
		ArimoBold            EnumItem
		Unknown              EnumItem
	}
	RenderingTestComparisonMethod struct {
		psnr EnumItem
		diff EnumItem
	}
	HoverAnimateSpeed struct {
		VerySlow EnumItem
		Slow     EnumItem
		Medium   EnumItem
		Fast     EnumItem
		VeryFast EnumItem
	}
	AnalyticsEconomyAction struct {
		Default EnumItem
		Acquire EnumItem
		Spend   EnumItem
	}
	ParticleEmitterShape struct {
		Box      EnumItem
		Sphere   EnumItem
		Cylinder EnumItem
		Disc     EnumItem
	}
	ChatCallbackType struct {
		OnCreatingChatWindow      EnumItem
		OnClientSendingMessage    EnumItem
		OnClientFormattingMessage EnumItem
		OnServerReceivingMessage  EnumItem
	}
	CharacterControlMode struct {
		Default                EnumItem
		Legacy                 EnumItem
		NoCharacterController  EnumItem
		LuaCharacterController EnumItem
	}
	MuteState struct {
		Unmuted EnumItem
		Muted   EnumItem
	}
	AnalyticsProgressionStatus struct {
		Default  EnumItem
		Begin    EnumItem
		Complete EnumItem
		Abandon  EnumItem
		Fail     EnumItem
	}
	R15CollisionType struct {
		OuterBox EnumItem
		InnerBox EnumItem
	}
	AvatarItemType struct {
		Asset  EnumItem
		Bundle EnumItem
	}
	MeshType struct {
		Head           EnumItem
		Torso          EnumItem
		Wedge          EnumItem
		Sphere         EnumItem
		Cylinder       EnumItem
		FileMesh       EnumItem
		Brick          EnumItem
		Prism          EnumItem
		Pyramid        EnumItem
		ParallelRamp   EnumItem
		RightAngleRamp EnumItem
		CornerWedge    EnumItem
	}
	AutomaticSize struct {
		None EnumItem
		X    EnumItem
		Y    EnumItem
		XY   EnumItem
	}
	FluidForces struct {
		Default      EnumItem
		Experimental EnumItem
	}
	SafeAreaCompatibility struct {
		None                EnumItem
		FullscreenExtension EnumItem
	}
	ActuatorType struct {
		None  EnumItem
		Motor EnumItem
		Servo EnumItem
	}
	PlaybackState struct {
		Begin     EnumItem
		Delayed   EnumItem
		Playing   EnumItem
		Paused    EnumItem
		Completed EnumItem
		Cancelled EnumItem
	}
	CompletionItemTag struct {
		Deprecated                    EnumItem
		IncorrectIndexType            EnumItem
		PluginPermissions             EnumItem
		CommandLinePermissions        EnumItem
		RobloxPermissions             EnumItem
		AddParens                     EnumItem
		PutCursorInParens             EnumItem
		TypeCorrect                   EnumItem
		ClientServerBoundaryViolation EnumItem
		Invalidated                   EnumItem
		PutCursorBeforeEnd            EnumItem
	}
	MaterialPattern struct {
		Regular EnumItem
		Organic EnumItem
	}
	LoadDynamicHeads struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	DebuggerStatus struct {
		Success          EnumItem
		Timeout          EnumItem
		ConnectionLost   EnumItem
		InvalidResponse  EnumItem
		InternalError    EnumItem
		InvalidState     EnumItem
		RpcError         EnumItem
		InvalidArgument  EnumItem
		ConnectionClosed EnumItem
	}
	HumanoidRigType struct {
		R6  EnumItem
		R15 EnumItem
	}
	HumanoidStateType struct {
		FallingDown       EnumItem
		Ragdoll           EnumItem
		GettingUp         EnumItem
		Jumping           EnumItem
		Swimming          EnumItem
		Freefall          EnumItem
		Flying            EnumItem
		Landed            EnumItem
		Running           EnumItem
		RunningNoPhysics  EnumItem
		StrafingNoPhysics EnumItem
		Climbing          EnumItem
		Seated            EnumItem
		PlatformStanding  EnumItem
		Dead              EnumItem
		Physics           EnumItem
		None              EnumItem
	}
	MeshPartDetailLevel struct {
		DistanceBased EnumItem
		Level00       EnumItem
		Level01       EnumItem
		Level02       EnumItem
		Level03       EnumItem
		Level04       EnumItem
	}
	PreferredTextSize struct {
		Medium  EnumItem
		Large   EnumItem
		Larger  EnumItem
		Largest EnumItem
	}
	GuiType struct {
		Core             EnumItem
		Custom           EnumItem
		PlayerNameplates EnumItem
		CustomBillboards EnumItem
		CoreBillboards   EnumItem
	}
	DraggerMovementMode struct {
		Geometric EnumItem
		Physical  EnumItem
	}
	SaveAvatarThumbnailCustomizationFailure struct {
		BadThumbnailType  EnumItem
		BadYRotDeg        EnumItem
		BadFieldOfViewDeg EnumItem
		BadDistanceScale  EnumItem
		Other             EnumItem
		Throttled         EnumItem
	}
	HttpCachePolicy struct {
		None                    EnumItem
		Full                    EnumItem
		DataOnly                EnumItem
		Default                 EnumItem
		InternalRedirectRefresh EnumItem
	}
	ReservedHighlightId struct {
		Standard  EnumItem
		Selection EnumItem
		Hover     EnumItem
		Active    EnumItem
	}
	DraggerCoordinateSpace struct {
		Object EnumItem
		World  EnumItem
	}
	ComputerMovementMode struct {
		Default       EnumItem
		KeyboardMouse EnumItem
		ClickToMove   EnumItem
	}
	ReturnKeyType struct {
		Default EnumItem
		Done    EnumItem
		Go      EnumItem
		Next    EnumItem
		Search  EnumItem
		Send    EnumItem
	}
	SalesTypeFilter struct {
		All          EnumItem
		Collectibles EnumItem
		Premium      EnumItem
	}
	OrientationAlignmentMode struct {
		OneAttachment EnumItem
		TwoAttachment EnumItem
	}
	ImageAlphaType struct {
		Default         EnumItem
		LockCanvasAlpha EnumItem
		LockCanvasColor EnumItem
	}
	DevComputerMovementMode struct {
		UserChoice    EnumItem
		KeyboardMouse EnumItem
		ClickToMove   EnumItem
		Scriptable    EnumItem
	}
	ScopeCheckResult struct {
		ConsentAccepted EnumItem
		InvalidScopes   EnumItem
		Timeout         EnumItem
		NoUserInput     EnumItem
		BackendError    EnumItem
		UnexpectedError EnumItem
		InvalidArgument EnumItem
		ConsentDenied   EnumItem
	}
	RibbonTool struct {
		Select         EnumItem
		Scale          EnumItem
		Rotate         EnumItem
		Move           EnumItem
		Transform      EnumItem
		ColorPicker    EnumItem
		MaterialPicker EnumItem
		Group          EnumItem
		Ungroup        EnumItem
		None           EnumItem
		PivotEditor    EnumItem
	}
	ChatPrivacyMode struct {
		AllUsers EnumItem
		NoOne    EnumItem
		Friends  EnumItem
	}
	PackagePermission struct {
		None     EnumItem
		NoAccess EnumItem
		Revoked  EnumItem
		UseView  EnumItem
		Edit     EnumItem
		Own      EnumItem
	}
	BodyPartR15 struct {
		Head          EnumItem
		UpperTorso    EnumItem
		LowerTorso    EnumItem
		LeftFoot      EnumItem
		LeftLowerLeg  EnumItem
		LeftUpperLeg  EnumItem
		RightFoot     EnumItem
		RightLowerLeg EnumItem
		RightUpperLeg EnumItem
		LeftHand      EnumItem
		LeftLowerArm  EnumItem
		LeftUpperArm  EnumItem
		RightHand     EnumItem
		RightLowerArm EnumItem
		RightUpperArm EnumItem
		RootPart      EnumItem
		Unknown       EnumItem
	}
	CollisionFidelity struct {
		Default                    EnumItem
		Hull                       EnumItem
		Box                        EnumItem
		PreciseConvexDecomposition EnumItem
	}
	ScreenOrientation struct {
		LandscapeLeft   EnumItem
		LandscapeRight  EnumItem
		LandscapeSensor EnumItem
		Portrait        EnumItem
		Sensor          EnumItem
	}
	RenderPriority struct {
		First     EnumItem
		Input     EnumItem
		Camera    EnumItem
		Character EnumItem
		Last      EnumItem
	}
	BorderMode struct {
		Outline EnumItem
		Middle  EnumItem
		Inset   EnumItem
	}
	Axis struct {
		X EnumItem
		Y EnumItem
		Z EnumItem
	}
	CellBlock struct {
		Solid              EnumItem
		VerticalWedge      EnumItem
		CornerWedge        EnumItem
		InverseCornerWedge EnumItem
		HorizontalWedge    EnumItem
	}
	ExplosionType struct {
		NoCraters EnumItem
		Craters   EnumItem
	}
	JoinSource struct {
		CreatedItemAttribution EnumItem
	}
	DialogBehaviorType struct {
		SinglePlayer    EnumItem
		MultiplePlayers EnumItem
	}
	RenderFidelity struct {
		Automatic   EnumItem
		Precise     EnumItem
		Performance EnumItem
	}
	HttpCompression struct {
		None EnumItem
		Gzip EnumItem
	}
	CreatorTypeFilter struct {
		User  EnumItem
		Group EnumItem
		All   EnumItem
	}
	DataStoreRequestType struct {
		GetAsync                EnumItem
		SetIncrementAsync       EnumItem
		UpdateAsync             EnumItem
		GetSortedAsync          EnumItem
		SetIncrementSortedAsync EnumItem
		OnUpdate                EnumItem
		ListAsync               EnumItem
		GetVersionAsync         EnumItem
		RemoveVersionAsync      EnumItem
	}
	MoverConstraintRootBehaviorMode struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	PositionAlignmentMode struct {
		OneAttachment EnumItem
		TwoAttachment EnumItem
	}
	OverrideMouseIconBehavior struct {
		None      EnumItem
		ForceShow EnumItem
		ForceHide EnumItem
	}
	MouseBehavior struct {
		Default             EnumItem
		LockCenter          EnumItem
		LockCurrentPosition EnumItem
	}
	InputType struct {
		NoInput  EnumItem
		Constant EnumItem
		Sin      EnumItem
	}
	Limb struct {
		Head     EnumItem
		Torso    EnumItem
		LeftArm  EnumItem
		RightArm EnumItem
		LeftLeg  EnumItem
		RightLeg EnumItem
		Unknown  EnumItem
	}
	SecurityCapability struct {
		RunClientScript    EnumItem
		RunServerScript    EnumItem
		AccessOutsideWrite EnumItem
		AssetRequire       EnumItem
		LoadString         EnumItem
		ScriptGlobals      EnumItem
		CreateInstances    EnumItem
		Basic              EnumItem
		Audio              EnumItem
		DataStore          EnumItem
		Network            EnumItem
		Physics            EnumItem
		UI                 EnumItem
		CSG                EnumItem
		Chat               EnumItem
		Animation          EnumItem
		Avatar             EnumItem
		Input              EnumItem
		Environment        EnumItem
		RemoteEvent        EnumItem
		LegacySound        EnumItem
	}
	RejectCharacterDeletions struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	LocationType struct {
		Character      EnumItem
		Camera         EnumItem
		ObjectPosition EnumItem
	}
	FriendStatus struct {
		Unknown               EnumItem
		NotFriend             EnumItem
		Friend                EnumItem
		FriendRequestSent     EnumItem
		FriendRequestReceived EnumItem
	}
	DeviceLevel struct {
		Low    EnumItem
		Medium EnumItem
		High   EnumItem
	}
	ProximityPromptExclusivity struct {
		OnePerButton EnumItem
		OneGlobally  EnumItem
		AlwaysShow   EnumItem
	}
	PartType struct {
		Ball        EnumItem
		Block       EnumItem
		Cylinder    EnumItem
		Wedge       EnumItem
		CornerWedge EnumItem
	}
	CustomCameraMode struct {
		Default EnumItem
		Classic EnumItem
		Follow  EnumItem
	}
	GuiState struct {
		Idle            EnumItem
		Hover           EnumItem
		Press           EnumItem
		NonInteractable EnumItem
	}
	NormalId struct {
		Right  EnumItem
		Top    EnumItem
		Back   EnumItem
		Left   EnumItem
		Bottom EnumItem
		Front  EnumItem
	}
	AutoIndentRule struct {
		Off      EnumItem
		Absolute EnumItem
		Relative EnumItem
	}
	ProductPurchaseDecision struct {
		NotProcessedYet EnumItem
		PurchaseGranted EnumItem
	}
	AnimationPriority struct {
		Core     EnumItem
		Idle     EnumItem
		Movement EnumItem
		Action   EnumItem
		Action2  EnumItem
		Action3  EnumItem
		Action4  EnumItem
	}
	GraphicsMode struct {
		Automatic  EnumItem
		Direct3D11 EnumItem
		OpenGL     EnumItem
		Metal      EnumItem
		Vulkan     EnumItem
		NoGraphics EnumItem
	}
	RuntimeUndoBehavior struct {
		Aggregate EnumItem
		Snapshot  EnumItem
		Hybrid    EnumItem
	}
	HumanoidCollisionType struct {
		OuterBox EnumItem
		InnerBox EnumItem
	}
	IKControlConstraintSupport struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	AdornCullingMode struct {
		Automatic EnumItem
		Never     EnumItem
	}
	DevCameraOcclusionMode struct {
		Zoom      EnumItem
		Invisicam EnumItem
	}
	AdUnitStatus struct {
		Inactive EnumItem
		Active   EnumItem
	}
	ParticleOrientation struct {
		FacingCamera          EnumItem
		FacingCameraWorldUp   EnumItem
		VelocityParallel      EnumItem
		VelocityPerpendicular EnumItem
	}
	AssetType struct {
		Image               EnumItem
		TShirt              EnumItem
		Audio               EnumItem
		Mesh                EnumItem
		Lua                 EnumItem
		Hat                 EnumItem
		Place               EnumItem
		Model               EnumItem
		Shirt               EnumItem
		Pants               EnumItem
		Decal               EnumItem
		Head                EnumItem
		Face                EnumItem
		Gear                EnumItem
		Badge               EnumItem
		Animation           EnumItem
		Torso               EnumItem
		RightArm            EnumItem
		LeftArm             EnumItem
		LeftLeg             EnumItem
		RightLeg            EnumItem
		Package             EnumItem
		GamePass            EnumItem
		Plugin              EnumItem
		MeshPart            EnumItem
		HairAccessory       EnumItem
		FaceAccessory       EnumItem
		NeckAccessory       EnumItem
		ShoulderAccessory   EnumItem
		FrontAccessory      EnumItem
		BackAccessory       EnumItem
		WaistAccessory      EnumItem
		ClimbAnimation      EnumItem
		DeathAnimation      EnumItem
		FallAnimation       EnumItem
		IdleAnimation       EnumItem
		JumpAnimation       EnumItem
		RunAnimation        EnumItem
		SwimAnimation       EnumItem
		WalkAnimation       EnumItem
		PoseAnimation       EnumItem
		EarAccessory        EnumItem
		EyeAccessory        EnumItem
		EmoteAnimation      EnumItem
		Video               EnumItem
		TShirtAccessory     EnumItem
		ShirtAccessory      EnumItem
		PantsAccessory      EnumItem
		JacketAccessory     EnumItem
		SweaterAccessory    EnumItem
		ShortsAccessory     EnumItem
		LeftShoeAccessory   EnumItem
		RightShoeAccessory  EnumItem
		DressSkirtAccessory EnumItem
		FontFamily          EnumItem
		EyebrowAccessory    EnumItem
		EyelashAccessory    EnumItem
		MoodAnimation       EnumItem
		DynamicHead         EnumItem
	}
	CompletionTriggerKind struct {
		Invoked                         EnumItem
		TriggerCharacter                EnumItem
		TriggerForIncompleteCompletions EnumItem
	}
	ModelLevelOfDetail struct {
		Automatic     EnumItem
		StreamingMesh EnumItem
		Disabled      EnumItem
	}
	ModelStreamingBehavior struct {
		Default  EnumItem
		Legacy   EnumItem
		Improved EnumItem
	}
	AppShellActionType struct {
		None                   EnumItem
		OpenApp                EnumItem
		TapChatTab             EnumItem
		TapConversationEntry   EnumItem
		TapAvatarTab           EnumItem
		ReadConversation       EnumItem
		TapGamePageTab         EnumItem
		TapHomePageTab         EnumItem
		GamePageLoaded         EnumItem
		HomePageLoaded         EnumItem
		AvatarEditorPageLoaded EnumItem
	}
	ClientAnimatorThrottlingMode struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	SelectionBehavior struct {
		Escape EnumItem
		Stop   EnumItem
	}
	ModifierKey struct {
		Shift EnumItem
		Ctrl  EnumItem
		Alt   EnumItem
		Meta  EnumItem
	}
	IKControlType struct {
		Transform EnumItem
		Position  EnumItem
		Rotation  EnumItem
		LookAt    EnumItem
	}
	CatalogSortType struct {
		Relevance       EnumItem
		PriceHighToLow  EnumItem
		PriceLowToHigh  EnumItem
		MostFavorited   EnumItem
		RecentlyCreated EnumItem
		Bestselling     EnumItem
	}
	AvatarContextMenuOption struct {
		Friend      EnumItem
		Chat        EnumItem
		Emote       EnumItem
		InspectMenu EnumItem
	}
	DevComputerCameraMovementMode struct {
		UserChoice   EnumItem
		Classic      EnumItem
		Follow       EnumItem
		Orbital      EnumItem
		CameraToggle EnumItem
	}
	ComputerCameraMovementMode struct {
		Default      EnumItem
		Classic      EnumItem
		Follow       EnumItem
		Orbital      EnumItem
		CameraToggle EnumItem
	}
	ChatStyle struct {
		Classic          EnumItem
		Bubble           EnumItem
		ClassicAndBubble EnumItem
	}
	RotationType struct {
		MovementRelative EnumItem
		CameraRelative   EnumItem
	}
	FinishRecordingOperation struct {
		Cancel EnumItem
		Commit EnumItem
		Append EnumItem
	}
	ParticleFlipbookTextureCompatible struct {
		NotCompatible EnumItem
		Compatible    EnumItem
		Unknown       EnumItem
	}
	MembershipType struct {
		None                   EnumItem
		BuildersClub           EnumItem
		TurboBuildersClub      EnumItem
		OutrageousBuildersClub EnumItem
		Premium                EnumItem
	}
	ListDisplayMode struct {
		Horizontal EnumItem
		Vertical   EnumItem
	}
	OperationType struct {
		Null         EnumItem
		Union        EnumItem
		Subtraction  EnumItem
		Intersection EnumItem
		Primitive    EnumItem
	}
	AdTeleportMethod struct {
		Undefined            EnumItem
		PortalForward        EnumItem
		InGameMenuBackButton EnumItem
		UIBackButton         EnumItem
	}
	GameAvatarType struct {
		R6           EnumItem
		R15          EnumItem
		PlayerChoice EnumItem
	}
	SelectionRenderMode struct {
		Outlines      EnumItem
		BoundingBoxes EnumItem
		Both          EnumItem
	}
	DominantAxis struct {
		Width  EnumItem
		Height EnumItem
	}
	ConnectionError struct {
		OK                                      EnumItem
		Unknown                                 EnumItem
		DisconnectErrors                        EnumItem
		DisconnectBadhash                       EnumItem
		DisconnectSecurityKeyMismatch           EnumItem
		DisconnectProtocolMismatch              EnumItem
		DisconnectReceivePacketError            EnumItem
		DisconnectReceivePacketStreamError      EnumItem
		DisconnectSendPacketError               EnumItem
		DisconnectIllegalTeleport               EnumItem
		DisconnectDuplicatePlayer               EnumItem
		DisconnectDuplicateTicket               EnumItem
		DisconnectTimeout                       EnumItem
		DisconnectLuaKick                       EnumItem
		DisconnectOnRemoteSysStats              EnumItem
		DisconnectHashTimeout                   EnumItem
		DisconnectCloudEditKick                 EnumItem
		DisconnectPlayerless                    EnumItem
		DisconnectNewSecurityKeyMismatch        EnumItem
		DisconnectEvicted                       EnumItem
		DisconnectDevMaintenance                EnumItem
		DisconnectRobloxMaintenance             EnumItem
		DisconnectRejoin                        EnumItem
		DisconnectConnectionLost                EnumItem
		DisconnectIdle                          EnumItem
		DisconnectRaknetErrors                  EnumItem
		DisconnectWrongVersion                  EnumItem
		DisconnectBySecurityPolicy              EnumItem
		DisconnectBlockedIP                     EnumItem
		DisconnectClientFailure                 EnumItem
		DisconnectClientRequest                 EnumItem
		DisconnectPrivateServerKickout          EnumItem
		DisconnectModeratedGame                 EnumItem
		ServerShutdown                          EnumItem
		ReplicatorTimeout                       EnumItem
		PlayerRemoved                           EnumItem
		DisconnectOutOfMemoryKeepPlayingLeave   EnumItem
		DisconnectRomarkEndOfTest               EnumItem
		DisconnectCollaboratorPermissionRevoked EnumItem
		DisconnectCollaboratorUnderage          EnumItem
		NetworkInternal                         EnumItem
		NetworkSend                             EnumItem
		NetworkTimeout                          EnumItem
		NetworkMisbehavior                      EnumItem
		NetworkSecurity                         EnumItem
		ReplacementReady                        EnumItem
		ServerEmpty                             EnumItem
		PlacelaunchErrors                       EnumItem
		PlacelaunchDisabled                     EnumItem
		PlacelaunchError                        EnumItem
		PlacelaunchGameEnded                    EnumItem
		PlacelaunchGameFull                     EnumItem
		PlacelaunchUserLeft                     EnumItem
		PlacelaunchRestricted                   EnumItem
		PlacelaunchUnauthorized                 EnumItem
		PlacelaunchFlooded                      EnumItem
		PlacelaunchHashExpired                  EnumItem
		PlacelaunchHashException                EnumItem
		PlacelaunchPartyCannotFit               EnumItem
		PlacelaunchHttpError                    EnumItem
		PlacelaunchUserPrivacyUnauthorized      EnumItem
		PlacelaunchCreatorBan                   EnumItem
		PlacelaunchCustomMessage                EnumItem
		PlacelaunchOtherError                   EnumItem
		TeleportErrors                          EnumItem
		TeleportFailure                         EnumItem
		TeleportGameNotFound                    EnumItem
		TeleportGameEnded                       EnumItem
		TeleportGameFull                        EnumItem
		TeleportUnauthorized                    EnumItem
		TeleportFlooded                         EnumItem
		TeleportIsTeleporting                   EnumItem
	}
	SelfViewPosition struct {
		LastPosition EnumItem
		TopLeft      EnumItem
		TopRight     EnumItem
		BottomLeft   EnumItem
		BottomRight  EnumItem
	}
	ExperienceAuthScope struct {
		DefaultScope        EnumItem
		CreatorAssetsCreate EnumItem
	}
	ChatColor struct {
		Blue  EnumItem
		Green EnumItem
		Red   EnumItem
		White EnumItem
	}
	BundleType struct {
		BodyParts         EnumItem
		Animations        EnumItem
		Shoes             EnumItem
		DynamicHead       EnumItem
		DynamicHeadAvatar EnumItem
	}
	NoiseType struct {
		SimplexGabor EnumItem
	}
	AvatarChatServiceFeature struct {
		None                 EnumItem
		UniverseAudio        EnumItem
		UniverseVideo        EnumItem
		PlaceAudio           EnumItem
		PlaceVideo           EnumItem
		UserAudioEligible    EnumItem
		UserAudio            EnumItem
		UserVideoEligible    EnumItem
		UserVideo            EnumItem
		UserBanned           EnumItem
		UserVerifiedForVoice EnumItem
	}
	ParticleFlipbookLayout struct {
		None    EnumItem
		Grid2x2 EnumItem
		Grid4x4 EnumItem
		Grid8x8 EnumItem
	}
	InfoType struct {
		Asset        EnumItem
		Product      EnumItem
		GamePass     EnumItem
		Subscription EnumItem
		Bundle       EnumItem
	}
	AvatarPromptResult struct {
		Success          EnumItem
		PermissionDenied EnumItem
		Failed           EnumItem
	}
	HandlesStyle struct {
		Resize   EnumItem
		Movement EnumItem
	}
	ButtonStyle struct {
		Custom                    EnumItem
		RobloxButtonDefault       EnumItem
		RobloxButton              EnumItem
		RobloxRoundButton         EnumItem
		RobloxRoundDefaultButton  EnumItem
		RobloxRoundDropdownButton EnumItem
	}
	DragDetectorDragStyle struct {
		TranslateLine        EnumItem
		TranslatePlane       EnumItem
		TranslatePlaneOrLine EnumItem
		TranslateLineOrPlane EnumItem
		TranslateViewPlane   EnumItem
		RotateAxis           EnumItem
		RotateTrackball      EnumItem
		Scriptable           EnumItem
		BestForDevice        EnumItem
	}
	RigType struct {
		R15    EnumItem
		Custom EnumItem
		None   EnumItem
	}
	BreakpointRemoveReason struct {
		Requested     EnumItem
		ScriptChanged EnumItem
		ScriptRemoved EnumItem
	}
	RtlTextSupport struct {
		Default  EnumItem
		Disabled EnumItem
		Enabled  EnumItem
	}
	AnalyticsEconomyTransactionType struct {
		IAP                EnumItem
		Shop               EnumItem
		Gameplay           EnumItem
		ContextualPurchase EnumItem
		TimedReward        EnumItem
		Onboarding         EnumItem
	}
	OutfitSource struct {
		All       EnumItem
		Created   EnumItem
		Purchased EnumItem
	}
	LiveEditingAtomicUpdateResponse struct {
		Success                 EnumItem
		FailureGuidNotFound     EnumItem
		FailureHashMismatch     EnumItem
		FailureOperationIllegal EnumItem
	}
	CoreGuiType struct {
		PlayerList EnumItem
		Health     EnumItem
		Backpack   EnumItem
		Chat       EnumItem
		All        EnumItem
		EmotesMenu EnumItem
		SelfView   EnumItem
		Captures   EnumItem
	}
	CommandPermission struct {
		Plugin    EnumItem
		LocalUser EnumItem
	}
	DialogPurpose struct {
		Quest EnumItem
		Help  EnumItem
		Shop  EnumItem
	}
	Genre struct {
		All         EnumItem
		TownAndCity EnumItem
		Fantasy     EnumItem
		SciFi       EnumItem
		Ninja       EnumItem
		Scary       EnumItem
		Pirate      EnumItem
		Adventure   EnumItem
		Sports      EnumItem
		Funny       EnumItem
		WildWest    EnumItem
		War         EnumItem
		SkatePark   EnumItem
		Tutorial    EnumItem
	}
	AppShellFeature struct {
		None         EnumItem
		Chat         EnumItem
		AvatarEditor EnumItem
		GamePage     EnumItem
		HomePage     EnumItem
		More         EnumItem
		Landing      EnumItem
	}
	ModelStreamingMode struct {
		Default             EnumItem
		Atomic              EnumItem
		Persistent          EnumItem
		PersistentPerPlayer EnumItem
		Nonatomic           EnumItem
	}
}

var (
	Enum = Enums{}
)
