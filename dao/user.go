package dao

import (
	"ecommerce/model"
	"ecommerce/utils"
	"time"
)

func CreateUser(req *model.UserMassage) error {
	result := DB.Create(req)
	return result.Error
}

func IsRepeatUser(req *model.UserMassage) (bool, error) {
	var count int64
	result := DB.Model(&model.UserMassage{}).Where("account=?", req.Account).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func IsAccountAndPassword(password string, account int) int {
	var selectPassword model.UserMassage
	var inputAccount model.UserMassage
	isAccount := DB.Model(&model.UserMassage{}).Select("account").Where("account=?", account).First(&inputAccount)
	if isAccount != nil {
		return -1 //账号不存在
	}
	if inputAccount.Account != account {
		return -1
	}
	result := DB.Model(&model.UserMassage{}).Select("password").Where("account=?", account).First(&selectPassword)
	if result.Error != nil {
		return -1 //查询失败
	}
	if selectPassword.Password != password {
		return 0 //密码错误
	}
	return 1
}

func UpdatePassword(req *model.UserMassage, new string) error {
	result := DB.Where("account=?", req.Account).Update("password", new)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserInfo(account int) (model.UserMassage, error) {
	var userInfo model.UserMassage
	result := DB.Model(&model.UserMassage{}).Omit("up_password", "password").Where("account=?", account).First(&userInfo)
	if result.Error != nil {
		return userInfo, result.Error
	}
	return userInfo, nil
}

func PostTokenJwt(uid int64) {
	var userToken = model.UserToken{
		Uid: uid,
	}
	token, _ := utils.SetTokenJwt(uid, time.Minute*30)
	userToken.Token = token
	result := DB.Create(&userToken)
	if result.Error != nil {
	}
}
