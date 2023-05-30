package api

import (
	"fmt"

	"github.com/agile-app/flexdb/internal/profile"
)

type H map[string]string

const (
	method       = "method"
	url          = "url"
	responseType = "responseType" // e,g. blob
)

func (h H) Method() string {
	return h[method]
}

func (h H) URL() string {
	return h[url]
}

func UpdateColumnAPI(appID, modelID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/meta/App/%d/Model/%d/Column/$id",
			profile.P.APIBaseEndpoint, appID, modelID),
		method: "put",
	}
}

func SaveColumnsOrderAPI(appID, modelID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/meta/App/%d/Model/%d/Column/",
			profile.P.APIBaseEndpoint, appID, modelID),
		method: "put",
	}
}

func AddColumnAPI(appID, modelID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/meta/App/%d/Model/%d/Column/",
			profile.P.APIBaseEndpoint, appID, modelID),
		method: "post",
	}
}

func DeprecateColumnAPI(appID, modelID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/meta/App/%d/Model/%d/Column/?id=$id",
			profile.P.APIBaseEndpoint, appID, modelID),
		method: "delete",
	}
}

func CreateRowAPI(modelID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/data/Model/%d/",
			profile.P.APIBaseEndpoint, modelID),
		method: "post",
	}
}

func FindRowsAPI(modelID int64) string {
	return fmt.Sprintf("%s/api/v0.1/data/Model/%d/Row",
		profile.P.APIBaseEndpoint, modelID)
}

func EditRowInitAPI(modelID int64) string {
	return fmt.Sprintf("%s/api/v0.1/data/Model/%d/Row/$id",
		profile.P.APIBaseEndpoint, modelID)
}

func UpdateRowAPI(modelID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/data/Model/%d/Row/$id",
			profile.P.APIBaseEndpoint, modelID),
		method: "put",
	}
}

func DeleteRowAPI(modelID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/data/Model/%d/Row/$id",
			profile.P.APIBaseEndpoint, modelID),
		method: "delete",
	}
}

func RowQuickSaveAPI(modelID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/data/Model/%d/",
			profile.P.APIBaseEndpoint, modelID),
		method: "put",
	}
}

func FindAppsAPI(q string) string {
	return fmt.Sprintf("%s/api/v0.1/meta/App/?q=%s", profile.P.APIBaseEndpoint, q)
}

func UploadAppAPI() H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/meta/App/?waitSeconds=2",
			profile.P.APIBaseEndpoint),
		method: "post",
	}
}

func UpdateAppAPI(appID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/meta/App/%d",
			profile.P.APIBaseEndpoint, appID),
		method: "put",
	}
}

func UpdateModelAPI(appID, modelID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/meta/App/%d/Model/%d?waitSeconds=1",
			profile.P.APIBaseEndpoint, appID, modelID),
		method: "put",
	}
}

func ImportRowsAPI(appID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/data/App/%d",
			profile.P.APIBaseEndpoint, appID),
		method: "post",
	}
}

func AutocompleteShareRecommend(appID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/user/recommend/%d?term=$term",
			profile.P.APIBaseEndpoint, appID),
	}
}

func ShareAppToUser(appID int64) H {
	return H{
		url: fmt.Sprintf("%s/api/v0.1/user/share/%d",
			profile.P.APIBaseEndpoint, appID),
		method: "post",
	}
}
