package local

import (
	"litshop/src/config"
	"litshop/src/lvm/local/en"
	"litshop/src/lvm/local/zh-CN"
)

var langList map[string]map[string]string

func init() {
	langList = make(map[string]map[string]string, 2)

	langList["zh-CN"] = zh_CN.Local()
	langList["en"] = en.Local()
}

func Trans(key string) string {
	l, ok := langList[config.GetString("lang")]
	if !ok {
		return key
	}

	s, ok := l[key]
	if !ok {
		return key
	}

	return s
}
