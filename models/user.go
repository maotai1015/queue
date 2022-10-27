package models

import "time"

type User struct {
	ID				int64		`json:"id" gorm:"primary_key"`
	Phone			string		`json:"phone" gorm:"comment:手机号"`
	CreateTime		time.Time	`json:"create_time" gorm:"comment:创建时间"`
}
func (User) TableName() string {
	return "user"
}

func SaveInfo(phone string) (err error) {
	user := User{Phone: phone, CreateTime: time.Now()}
	err = db.Create(user).Error
	return
}