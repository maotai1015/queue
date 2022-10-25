package models

import "time"

type User struct {
	ID				int64		`json:"id" gorm:"primary_key"`
	Phone			string		`json:"phone" gorm:"comment:手机号"`
	CreateTime		time.Time	`json:"create_time" gorm:"comment:创建时间"`
}