package connector

import "fmt"

// MQ is a MQ connector.
type MQ interface {
	Consume(topic string) (msgs []interface{}, err error)

	Produce(topic string, msg interface{}) error
}

// RegisterMQ registers a MQ connector for an org.
func RegisterMQ(orgID OrgID, s MQ) {
	if _, present := mqConnectors[orgID]; present {
		panic(fmt.Errorf("orgID:%d already used", orgID))
	}

	mqConnectors[orgID] = s
}

// MQConnector returns the MQ instance of an org.
func MQConnector(orgID OrgID) MQ {
	return mqConnectors[orgID]
}
