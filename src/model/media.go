package model

import "time"

type MediaType = int

type Media struct {
	Model
	Type     MediaType     `json:"type"`
	Filename string        `json:"filename"`
	FullPath string        `json:"full_path"`
	Ext      string        `json:"ext"`
	Size     uint64        `json:"size"`
	Mime     string        `json:"mime"`
	Duration time.Duration `json:"duration"`
}

func (*Media) TableName() string {
	return "media"
}

func (*Media) Connection() string {
	return "default"
}
