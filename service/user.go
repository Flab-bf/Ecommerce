package service

import (
	"ecommerce/dao"
	"ecommerce/model"
	"errors"
)

func RegisterUser(req *model.UserChangePassword) error {
	var umsg model.UserMassage
	umsg.Account = req.Account
	umsg.Password = req.Password
	is, err := dao.IsRepeatUser(&umsg)
	if err != nil || is {
		return err //账号重复
	}
	err = dao.CreateUser(&umsg)
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(req *model.UserMassage) (error, string, string) {
	err := dao.IsAccountAndPassword(req.Password, req.Account)
	if err != nil {
		return err, "", ""
	}
	req.Uid, err = dao.FindUidFromAccount(req.Account)
	if err != nil {
		return err, "", ""
	}
	token, refreshToken := dao.PostTokenJwt(req.Uid)
	if refreshToken == "" || token == "" {
		return errors.New("nil token"), "", ""
	}
	return nil, token, refreshToken
}

func ChangePassword(req *model.UserChangePassword) error {
	err := dao.IsAccountAndPassword(req.Password, req.Account)
	if err != nil {
		return err
	}
	err = dao.UpdatePassword(req)
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfo(account int) (model.UserMassage, error) {
	info, err := dao.GetUserInfo(account)
	if err != nil {
		return info, err
	}
	return info, nil
}

func ChangeUserInfo(req *model.UserMassage) error {
	err := dao.PutUserInfo(req)
	if err != nil {
		return err
	}
	return nil
}
