package server

import (
	"github.com/agile-app/flexdb/internal/profile"
)

// PublishAPIs publishes apis according to the deployment configuration.
func (s *Server) PublishAPIs() {
	s.exportUserAPIs() // TODO kill it: frontend can tell itself

	if profile.P.ExportMetaAPIs {
		s.exportMetaAPIs()
	}

	if profile.P.ExportSchemaAPIs {
		s.exportSchemaAPIs()
	}

	if profile.P.ExportDataAPIs {
		s.exportDataAPIs()
	}
}
