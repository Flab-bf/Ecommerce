package dao

import (
	"ecommerce/model"
	"errors"
)

func GetProductList() ([]model.Product, error) {
	var List []model.Product
	result := DB.Model(&model.Product{}).Find(&List)
	if result.Error != nil {
		return nil, result.Error
	}
	return List, nil
}

func FindProduct(id int) (error, model.Cart) {
	var cart model.Cart
	result := DB.Model(&model.Product{}).Where("product_id=?", id).First(&cart)
	if result.Error != nil {
		return result.Error, model.Cart{}
	}
	result = DB.Model(&model.Product{}).Select("comment_num").Where("product_id=?", id).First(&cart.Num)
	if result.Error != nil {
		return result.Error, model.Cart{}
	}
	return nil, cart

}

func AddCart(cart model.Cart) error {
	var ex int64
	result := DB.Model(&model.Cart{}).Where("user_id=? And product_id=?", cart.UserId, cart.ProductId).Count(&ex)
	if ex != 0 || result.Error != nil {
		return errors.New("商品已存在")
	}
	result = DB.Model(&model.Cart{}).Create(&cart)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func GetCarts(uid int) ([]model.Cart, error) {
	var info []model.Cart
	result := DB.Model(model.Cart{}).Where("user_id=?", uid).Find(&info)
	if result.Error != nil {
		return nil, result.Error
	}
	return info, nil
}
