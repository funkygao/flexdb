package filter

import (
	"github.com/agile-app/flexdb/internal/profile"
	"github.com/agile-app/flexdb/pkg/api"
)

var whiteList = map[string]struct{}{
	// L2 Lead
	/*
		"ext.wangwenkai":  {},
		"ouyangxudong":    {},
		"liyaman":         {},
		"gaojinge":        {},
		"bjxzh":           {},
		"bjxjf":           {},
		"xiaoguang.wang":  {},
		"bjzhangmingming": {},
		"songning":        {},
		"bjzhoutong":      {},
		"bjmwn":           {},
		"chensenbiao":     {},
		"baixin9":         {},
		"bjxierui":        {},
		"zhangjunjun19":   {},
		"shenwenbin9":     {},*/

	// HR
	"gaojinge":    {},
	"zhanghan234": {},
	"qiaoran6":    {},
	"wanghan161":  {},
	"liuyifen":    {},
}

func SatisfyToBeKilledPermRule(pin string) bool {
	if profile.LocalDebug() && pin == api.Debugger {
		return true
	}

	if _, yes := whiteList[pin]; yes {
		return true
	}

	return false
}
