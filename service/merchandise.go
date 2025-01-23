package service

import (
	"ecommerce/dao"
	"ecommerce/model"
)

func ProductList() ([]model.Product, error) {
	return dao.GetProductList()
}

func AddCart(id int, userId int) error {
	err, cart := dao.FindProduct(id)
	if err != nil {
		return err
	}
	cart.UserId = userId
	err = dao.AddCart(cart)
	if err != nil {
		return err
	}
	return nil
}

func GetCarts(uid int) ([]model.Cart, error) {
	info, err := dao.GetCarts(uid)
	if err != nil {
		return nil, err
	}
	return info, nil
}
