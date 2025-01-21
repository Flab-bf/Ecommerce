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

func LoginUser(req *model.UserMassage) int {
	is := dao.IsAccountAndPassword(req.Password, req.Account)
	if is == -1 {
		return is
	}
	if is == 0 {
		return is
	}
	dao.PostTokenJwt(int64(req.Uid))
	return 1
}
