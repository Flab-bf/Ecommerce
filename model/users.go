package model

import "time"

type UserMassage struct {
	Account      int       `gorm:"not null;unique" json:"account"`
	Password     string    `gorm:"type:varchar(225);not null" json:"password"`
	NickName     string    `gorm:"not null" json:"nick_name"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpPassword   time.Time `gorm:"type:timestamp;autoUpdateTime"`
	Gender       string    `gorm:"type:int" json:"gender"`
	Birthday     string    `gorm:"type:date" json:"birthday"`
	Avatar       string    `gorm:"type:json"`
	Email        string    `gorm:"type:varchar(255)" json:"email"`
	Introduction string    `gorm:"type:text" json:"introduction"`
	QQ           string    `gorm:"type:varchar(255)" json:"qq"`
	Tel          string    `gorm:"type:varchar(255)" json:"tel"`
}
