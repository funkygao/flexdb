package connector

import "fmt"

// S3Store is a S3 alike blob storage connector.
type S3Store interface {
	PutObjectFromFile(bucketName, objectName, filePath string) (ret string, err error)

	PutObjectFromBytes(bucketName, objectName, md5, contentType string, b []byte) (ret string, err error)
}

// RegisterS3Store registers a S3Store for an org.
func RegisterS3Store(orgID OrgID, s S3Store) {
	if _, present := s3Connectors[orgID]; present {
		panic(fmt.Errorf("orgID:%d already used", orgID))
	}

	s3Connectors[orgID] = s
}

// S3Connector returns the S3Store instance of an org.
func S3Connector(orgID OrgID) S3Store {
	return s3Connectors[orgID]
}
