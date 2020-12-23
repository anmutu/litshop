package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"litshop/src/config/mysql"
	"time"
)

type Modeler interface {
	schema.Tabler
	Connection() string
	Ms() map[string]interface{}
}

type Model struct {
	ID        uint64         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	DB *gorm.DB
}

func (*Model) Connection() string {
	return "default"
}

func (m *Model) InitDB() {
	m.DB = mysql.GormClientByConn(m.Connection())
}

func (m *Model) Ms() map[string]interface{} {
	b, _ := json.Marshal(&m)
	var ms map[string]interface{}
	_ = json.Unmarshal(b, &ms)
	return ms
}
