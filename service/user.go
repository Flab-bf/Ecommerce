package service

import (
	"ecommerce/dao"
	"ecommerce/model"
)

func RegisterUser(req *model.UserMassage) error {
	is, err := dao.IsRepeatUser(req)
	if err != nil {
		return err //账号重复
	}
	if is {
		return err
	}
	err = dao.CreateUser(req)
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(req *model.UserMassage) (int, string) {
	is := dao.IsAccountAndPassword(req.Password, req.Account)
	if is != 1 {
		return is, ""
	}
	var err error
	req.Uid, err = dao.FindUidFromAccount(req.Account)
	if err != nil {

	}
	token := dao.PostTokenJwt(req.Uid)
	return 1, token
}

func ChangePassword(req *model.UserChangePassword) error {
	is := dao.IsAccountAndPassword(req.Password, req.Account)
	if is == 0 {
		//daiding******************************
		//**********************************
	}
	err := dao.UpdatePassword(req)
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
