package connector

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SSO is a single sign on connector.
type SSO interface {
	// CurrentUser extract pin from http context.
	CurrentUser(c *gin.Context) (pin string, err error)

	SearchUser(hint string) (interface{}, error)
}

// RegisterSSO registers a SSO for an org.
func RegisterSSO(orgID OrgID, s SSO) {
	if _, present := ssoConnectors[orgID]; present {
		panic(fmt.Errorf("orgID:%d already used", orgID))
	}

	ssoConnectors[orgID] = s
}

// SSOConnector returns a SSO connector.
func SSOConnector(orgID OrgID) SSO {
	return ssoConnectors[orgID]
}
