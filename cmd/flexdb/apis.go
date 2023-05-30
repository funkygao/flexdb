package main

import (
	"fmt"

	"github.com/agile-app/flexdb/internal/router"
	"github.com/funkygao/columnize"
	"github.com/pmylund/sortutil"
)

func showAPIs() {
	type api struct {
		path    string
		method  string
		handler interface{}
	}
	routes := router.Offer().Routes()
	apis := make([]api, 0, len(routes))
	for _, route := range router.Offer().Routes() {
		apis = append(apis, api{method: route.Method, path: route.Path, handler: route.Handler})
	}
	sortutil.AscByField(apis, "path")

	lines := []string{"Resource|Verb"}
	for _, api := range apis {
		lines = append(lines, fmt.Sprintf("%s|%s", api.path, api.method))
	}
	fmt.Println(columnize.SimpleFormat(lines))
}
