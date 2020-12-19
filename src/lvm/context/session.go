package context

import "litshop/src/model"

type Session struct {
	Customer    *model.Customer
	AccessToken string
	RequestId   string
}
