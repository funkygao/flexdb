package profile

import (
	js "github.com/funkygao/jsconf"
)

// LoadFrom load the specified Pig file.
// Different environment can have different Pig filename.
func LoadFrom(fn string) error {
	cf, err := js.Load(fn)
	if err != nil {
		return err
	}

	P = &Profile{Conf: cf}
	P.Debug = cf.Bool("debug", true)
	P.LocalDebug = cf.Bool("local_debug", false)
	P.BindAddr = cf.String("bind", ":8000")
	P.ExportMetaAPIs = cf.Bool("export_meta_api", true)
	P.ExportDataAPIs = cf.Bool("export_data_api", true)
	P.ExportSchemaAPIs = cf.Bool("export_schema_api", true)
	P.APIBaseEndpoint = cf.String("api_base_url", "http://localhost:8000")
	P.PaginationMaxRowsScan = cf.Int("maxpages", 2000)
	P.MetaDSN = cf.String("mdb", "root:@tcp(127.0.0.1:3306)/flexmeta?charset=utf8mb4&parseTime=True&loc=Local")
	P.DataDSNs = cf.StringList("ddb", nil)
	P.CORSAllowOrigins = cf.StringList("cors_allow_origins", []string{"http://localhost:7050"})
	P.PprofEndpoint = cf.String("pprof", "localhost:10120")
	P.Driver = cf.String("driver", "mysql")
	P.Cache = cf.String("cache", "blackhole")
	P.WarnScannedRowsThreshold = cf.Int("explain.rows.warn", 100)
	P.SLA = cf.String("sla", "dummy")
	P.Telemetry = cf.String("telemetry", "localFile")
	P.DefaultLocale = cf.String("locale", "zh_CN")

	if P.LocalDebug {
		P.Debug = true
	}

	return nil
}

func Debug() bool {
	return P.Debug
}

func LocalDebug() bool {
	return P.LocalDebug
}
