package entity

type AppStatus int16

// Label returns human readable kind label.
func (as AppStatus) Label() string {
	return appStatusLabels[as]
}

const (
	unknownAppStatus AppStatus = 0

	AppStatusInit    AppStatus = 1
	AppStatusOnline  AppStatus = 2
	AppStatusOffline AppStatus = 3
)
