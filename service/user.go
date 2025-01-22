package service

import (
	"ecommerce/dao"
	"ecommerce/model"
)

func RegisterUser(req *model.UserMassage) error {
	is, err := dao.IsRepeatUser(req)
	if err != nil || is {
		return err //账号重复
	}
	err = dao.CreateUser(req)
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(req *model.UserMassage) (error, string) {
	err := dao.IsAccountAndPassword(req.Password, req.Account)
	if err != nil {
		return err, ""
	}
	req.Uid, err = dao.FindUidFromAccount(req.Account)
	if err != nil {
		return err, ""
	}
	token := dao.PostTokenJwt(req.Uid)
	return nil, token
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
