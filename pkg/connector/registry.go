package connector

// OrgID is org id.
type OrgID int64

var (
	ssoConnectors = make(map[OrgID]SSO, 1)
	rpcConnectors = make(map[OrgID]RPC, 1)
	s3Connectors  = make(map[OrgID]S3Store, 1)
	mqConnectors  = make(map[OrgID]MQ, 1)
)
