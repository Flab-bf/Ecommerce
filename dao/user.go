package dao

import (
	"ecommerce/model"
	"ecommerce/utils"
	"time"
)

func CreateUser(req *model.UserMassage) error {
	req.CreatedAt = time.Now()
	req.UpPassword = time.Now()
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
	isAccount := DB.Model(&model.UserMassage{}).Select("account").
		Where("account=?", account).First(&inputAccount)
	if isAccount.Error != nil {
		return -3 //账号不存在
	}
	if inputAccount.Account != account {
		return -2
	}
	result := DB.Model(&model.UserMassage{}).Select("password").
		Where("account=?", account).First(&selectPassword)
	if result.Error != nil {
		return -1 //查询失败
	}
	if selectPassword.Password != password {
		return 0 //密码错误
	}
	return 1
}

func UpdatePassword(req *model.UserChangePassword) error {
	result := DB.Model(&model.UserMassage{}).Where("account=?", req.Account).
		Updates(&model.UserMassage{Password: req.NewPassword, UpPassword: time.Now()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserInfo(account int) (model.UserMassage, error) {
	var userInfo model.UserMassage
	result := DB.Model(&model.UserMassage{}).Omit("up_password", "password").
		Where("account=?", account).First(&userInfo)
	if result.Error != nil {
		return userInfo, result.Error
	}
	return userInfo, nil
}

func PutUserInfo(req *model.UserMassage) error {
	result := DB.Model(&model.UserMassage{}).Where("account=?", req.Account).Updates(req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindUidFromAccount(account int) (int, error) {
	var uid int
	result := DB.Model(&model.UserMassage{}).Select("uid").
		Where("account=?", account).First(&uid)
	if result.Error != nil {
		return 0, result.Error
	}
	return uid, nil
}

func PostTokenJwt(uid int) string {
	var userToken = model.UserToken{
		Uid: uid,
	}
	token, _ := utils.SetTokenJwt(uid, time.Minute*30)
	userToken.Token = token
	var count int64
	DB.Model(&model.UserToken{}).Where("uid=?", uid).Count(&count)
	if count == 0 {
		result := DB.Create(&userToken)
		if result.Error != nil {
		}
		return ""
	}
	result := DB.Model(&model.UserToken{}).Where("uid=?", uid).Update("token", token)
	if result.Error != nil {
		return ""
	}
	return token
}

func PutTokenJwt(uid int) string {
	var userToken = model.UserToken{
		Uid: uid,
	}
	token, _ := utils.RefreshToken(uid)
	userToken.Token = token
	var count int64
	DB.Model(&model.UserToken{}).Where("uid=?", uid).Count(&count)
	if count == 0 {
		result := DB.Create(&userToken)
		if result.Error != nil {
		}
		return ""
	}
	result := DB.Model(&model.UserToken{}).Where("uid=?", uid).Update("token", token)
	if result.Error != nil {

	}
	return token
}

func IsLegalUser(uid int) (bool, error) {
	var count int64
	result := DB.Model(&model.UserToken{}).Where("uid=?", uid).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	if count != 1 {
		return false, nil
	}
	return true, nil
}
