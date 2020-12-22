package model

import "time"

type MediaType = int

type Media struct {
	Model
	Type     MediaType     `json:"type" gorm:"type:integer(10)"`
	Filename string        `json:"filename" gorm:"type:varchar(255)"`
	FullPath string        `json:"full_path" gorm:"type:varchar(255)"`
	Ext      string        `json:"ext" gorm:"type:varchar(255)"`
	Size     uint64        `json:"size"`
	Mime     string        `json:"mime" gorm:"type:varchar(255)"`
	Duration time.Duration `json:"duration"`
}

func (*Media) TableName() string {
	return "media"
}

func (*Media) Connection() string {
	return "default"
}
