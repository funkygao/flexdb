package api

import "fmt"

func RowCRUDPage(appUUID string, appID, modelID int64) string {
	return fmt.Sprintf("/crud?app=%d&id=%d&uuid=%s",
		appID, modelID, appUUID)
}

func ModelSchemaDesignPage(appUUID string, appID, modelID int64) string {
	return fmt.Sprintf("/schema?app=%d&id=%d&uuid=%s",
		appID, modelID, appUUID)
}

func AppDesignPage(appUUID string, appID int64, menu string) string {
	return fmt.Sprintf("/design?id=%d&menu=%s&uuid=%s",
		appID, menu, appUUID)
}
