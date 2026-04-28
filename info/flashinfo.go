package flashinfo

func Default() Flashinfo {
	return Flashinfo{}
}

type Flashinfo struct {
	PartNumber string

	Vendor string

	Type NandType

	Capacity string

	DeviceWidth int

	Process string

	CellLevel string

	CE int

	CH int

	RB int

	Die int

	Voltage string

	Generation int

	Async bool

	Sync bool

	Spi bool

	Toggle bool

	ToggleN string

	Package string

	IDs []string

	EnterPrise bool

	Page_size string

	Block_size string

	Lead_Free bool

	Halogen_Free bool
	//特殊参数，为true时回传重调Decode
	Retry bool

	Unsupported_Reason string

	TempErature string

	SpeedGrade string

	CapGrade string

	EccEnable bool

	// SpecTek specific: when true, indicates the device uses half-page addressing/size behavior.
	// Upstream FlashDetector exposes this as extra info (halfPageAndSize).
	HalfPageAndSize bool

	Wafer bool
}

type NandType string

const (
	NAND        NandType = "Nand"
	INAND       NandType = "iNand"
	E2NAND      NandType = "E2NAND"
	Xpoint_3D   NandType = "3dXPoint"
	UNKWON_NAND NandType = ""
	CON         NandType = "Con"
)
