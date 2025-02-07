package dao

import (
	"ecommerce/model"
	"ecommerce/utils"
	"errors"
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

func IsAccountAndPassword(password string, account int) error {
	var selectPassword model.UserMassage
	var inputAccount model.UserMassage
	isAccount := DB.Model(&model.UserMassage{}).Select("account").
		Where("account=?", account).First(&inputAccount)
	if isAccount.Error != nil {
		return isAccount.Error //账号不存在
	}
	if inputAccount.Account != account {
		return errors.New("账号不存在")
	}
	result := DB.Model(&model.UserMassage{}).Select("password").
		Where("account=?", account).First(&selectPassword)
	if result.Error != nil {
		return result.Error //查询失败
	}
	if selectPassword.Password != password {
		return errors.New("密码错误") //密码错误
	}
	return nil
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
	result := DB.Model(&model.UserMassage{}).Where("uid=?", req.Uid).Updates(req)
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

func PostTokenJwt(uid int) (string, string) {
	var userToken = model.UserToken{
		Uid: uid,
	}
	token, _ := utils.SetAccessToken(uid)
	refreshToken, _ := utils.SetRefreshToken(uid)
	userToken.Token = token
	userToken.RefreshToken = refreshToken
	var count int64
	DB.Model(&model.UserToken{}).Where("uid=?", uid).Count(&count)
	if count == 0 {
		result := DB.Create(&userToken)
		if result.Error != nil {
			return "", ""
		}
	} else {
		result := DB.Model(&model.UserToken{}).Where("uid=?", uid).Updates(&userToken)
		if result.Error != nil {
			return "", ""
		}
	}
	return token, refreshToken
}

func PutTokenJwt(uid int) string {
	var userToken = model.UserToken{
		Uid: uid,
	}
	token, _ := utils.SetAccessToken(uid)
	userToken.Token = token
	result := DB.Model(&model.UserToken{}).Where("uid=?", uid).Update("token", token)
	if result.Error != nil {
		return ""
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
