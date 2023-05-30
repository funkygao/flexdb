package profile

import js "github.com/funkygao/jsconf"

// Profile is the runtime configuration of the current FlexDB instance.
type Profile struct {
	*js.Conf

	BindAddr        string
	APIBaseEndpoint string
	PprofEndpoint   string

	ExportSchemaAPIs bool // amis schema
	ExportMetaAPIs   bool // FlexDB metadata
	ExportDataAPIs   bool // FlexDB data and index
	ExportAdminAPIs  bool // FlexDB ops

	Debug      bool
	LocalDebug bool

	SSOBaseURL string //= "http://localhost:8000"
	SSOApp     string // = "test"
	SSOToken   string // = "test"

	// pagination max rows scan allowed
	PaginationMaxRowsScan int

	Driver           string
	MetaDSN          string
	DataDSNs         []string
	CORSAllowOrigins []string

	SLA           string
	Cache         string
	Telemetry     string
	DefaultLocale string

	WarnScannedRowsThreshold int
}

var (
	// P is the profile.
	P *Profile
)
