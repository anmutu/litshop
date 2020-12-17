package resource

import (
	"litshop/src/model"
	"litshop/src/pkg/d"
)

func CustomerResItem(resource *model.Customer) d.H {
	if resource == nil {
		return d.H{}
	}

	res := resource.Ms()
	delete(res, "password")

	return res
}

func CustomerResCollection(resource []*model.Customer) []d.H {
	var res []d.H
	res = make([]d.H, 0)

	for _, m := range resource {
		res = append(res, CustomerResItem(m))
	}

	return res
}
