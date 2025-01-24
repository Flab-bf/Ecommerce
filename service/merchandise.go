package service

import (
	"ecommerce/dao"
	"ecommerce/model"
)

func ProductList() ([]model.Product, error) {
	return dao.GetProductList()
}

func AddCart(id int, userId int) error {
	err, cart := dao.FindProductToCart(id)
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

func SearchInfoFromId(id int) (model.Product, error) {
	info, err := dao.GetProductFromId(id)
	if err != nil {
		return model.Product{}, err
	}
	return info, nil
}

func GetProductFromType(ty string) (model.Product, error) {
	info, err := dao.GetProductFromType(ty)
	if err != nil {
		return model.Product{}, err
	}
	return info, nil
}

func GetProductFromName(name string) (model.Product, error) {
	info, err := dao.GetProductFromName(name)
	if err != nil {
		return model.Product{}, err
	}
	return info, nil
}

func Incart(uid int, info model.Product) {
	in := dao.InCart(info.ProductId, uid)
	if in {
		info.IsAddedCart = true
	} else {
		info.IsAddedCart = false
	}
}
