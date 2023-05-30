package connector

import (
	"fmt"
	"time"
)

// RPC is a general remote procedure call protocol connector.
type RPC interface {
	Invoke(interfaceName, methodName, alias string, args []interface{}, timeout time.Duration) (map[string]interface{}, error)
}

// RegisterRPC registers a RPC for an org.
func RegisterRPC(orgID OrgID, s RPC) {
	if _, present := rpcConnectors[orgID]; present {
		panic(fmt.Errorf("orgID:%d already used", orgID))
	}

	rpcConnectors[orgID] = s
}

// RPCConnector returns the RPC instance of an org.
func RPCConnector(orgID OrgID) RPC {
	return rpcConnectors[orgID]
}
