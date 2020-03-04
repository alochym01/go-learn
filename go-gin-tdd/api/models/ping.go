package models

import (
	"github.com/jinzhu/gorm"
)

type PingModel struct {
	gorm.Model
	Ping string
}

func Ping(ping string) string {
	return "pong"
}
