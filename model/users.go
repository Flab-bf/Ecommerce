package model

import "time"

type UserMassage struct {
	Uid          int       `gorm:"primaryKey;autoIncrement" json:"-"`
	Account      int       `gorm:"not null;unique" json:"-" query:"account"`
	Password     string    `gorm:"type:varchar(225);not null" json:"-" query:"password"`
	NickName     string    `gorm:"not null" json:"nickName"`
	CreatedAt    time.Time `gorm:"type:timestamp;comment:创建时间" json:"-"`
	UpPassword   time.Time `gorm:"type:timestamp" json:"-"`
	Gender       int       `gorm:"type:int" json:"gender"`
	Birthday     string    `gorm:"type:varchar(255);null" json:"birthday"`
	Avatar       string    `gorm:"type:string" json:"avatar"`
	Email        string    `gorm:"type:varchar(255)" json:"email"`
	Introduction string    `gorm:"type:text" json:"introduction"`
	QQ           string    `gorm:"type:varchar(255)" json:"qq"`
	Tel          string    `gorm:"type:varchar(255)" json:"tel"`
}

type UserChangePassword struct {
	Account     int    `json:"account"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
type UserToken struct {
	Uid          int    `gorm:"primaryKey" json:"uid"`
	Token        string `gorm:"type:varchar(512)" json:"token"`
	RefreshToken string `gorm:"type:varchar(512)" json:"refreshToken"`
}
