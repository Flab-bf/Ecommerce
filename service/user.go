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
