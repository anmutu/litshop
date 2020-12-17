package model

import (
	"database/sql"
	"litshop/src/core/types"
)

type Customer struct {
	Model

	NickName string               `json:"nick_name"`
	Avatar   string               `json:"avatar"`
	Status   types.CustomerStatus `json:"status"`
	Phone    string               `json:"phone"`
	Email    string               `json:"email"`
	Age      int                  `json:"age"`
	Birth    sql.NullTime         `json:"birth"`
	Gender   types.CustomerGender `json:"gender"`
	Country  string               `json:"country"`
	Province string               `json:"province"`
	City     string               `json:"city"`
	Password string               `json:"password"`
}

func (*Customer) Connection() string {
	return "customer"
}

func (*Customer) TableName() string {
	return "customer"
}
