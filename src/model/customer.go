package model

import (
	"database/sql"
	"litshop/src/lvm/jwt"
	"litshop/src/lvm/literr"
	"litshop/src/lvm/types"
	"litshop/src/pkg/util"
)

//go:generate genrepository -dir . -structs=Customer
type Customer struct {
	Model

	NickName string               `json:"nick_name" gorm:"type:varchar(255)"`
	Avatar   string               `json:"avatar" gorm:"type:varchar(255)"`
	Status   types.CustomerStatus `json:"status" gorm:"type:integer(10)"`
	Phone    string               `json:"phone" gorm:"type:varchar(255)"`
	Email    string               `json:"email" gorm:"type:varchar(255)"`
	Age      int                  `json:"age" gorm:"type:varchar(255)"`
	Birth    sql.NullTime         `json:"birth"`
	Gender   types.CustomerGender `json:"gender" gorm:"type:tinyint"`
	Country  string               `json:"country" gorm:"type:varchar(255)"`
	Province string               `json:"province" gorm:"type:varchar(255)"`
	City     string               `json:"city" gorm:"type:varchar(255)"`
	Password string               `json:"password" gorm:"type:varchar(255)"`
}

func (*Customer) Connection() string {
	return "customer"
}

func (*Customer) TableName() string {
	return "customer"
}

func (m *Customer) Token() (string, error) {
	token, err := jwt.GenJwtToken(m)
	if err != nil {
		return "", literr.NewWithCode(literr.ErrCodeInternalError)
	}

	return token, nil
}

func (m *Customer) GetId() string {
	return util.Interface2String(m.ID)
}
