package dao

import (
	"ecommerce/model"
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

func IsAccountAndPassword(password string, account string) (int, error) {
	var selectPassword string
	result := DB.Model(&model.UserMassage{}).Where("account=?", account)
	if result.Error != nil {
		return -1, result.Error
	}
	yes := DB.Model(&model.UserMassage{}).Select("password").
		Where("account=?", account).First(&selectPassword)
	if yes.Error != nil || selectPassword != password {
		return 0, yes.Error
	}
	return 1, nil
}

func UpdatePassword(req *model.UserMassage, new string) error {
	result := DB.Where("account=?", req.Account).Update("password", new)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserInfo(account string) (model.UserMassage, error) {
	var userInfo model.UserMassage
	result := DB.Model(&model.UserMassage{}).Omit("up_password", "password").Where("account=?", account).First(&userInfo)
	if result.Error != nil {
		return userInfo, result.Error
	}
	return userInfo, nil
}
