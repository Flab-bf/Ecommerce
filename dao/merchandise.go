package dao

import (
	"ecommerce/model"
	"errors"
	"gorm.io/gorm"
)

func GetProductList() ([]model.Product, error) {
	var List []model.Product
	result := DB.Model(&model.Product{}).Find(&List)
	if result.Error != nil {
		return nil, result.Error
	}
	return List, nil
}

func FindProductToCart(id int) (error, model.Cart) {
	var cart model.Cart
	result := DB.Model(&model.Product{}).Omit("user_id").Where("product_id=?", id).First(&cart)
	if result.Error != nil {
		return result.Error, model.Cart{}
	}
	return nil, cart
}

func AddCart(cart model.Cart) error {
	var num int
	result := DB.Model(&model.Cart{}).Select("num").Where("product_id=?", cart.ProductId).First(&num)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	if num == 0 {
		cart.Num = 1
		result = DB.Model(&model.Cart{}).Create(&cart)
	} else {
		cart.Num = num + 1
		result = DB.Model(&model.Cart{}).Where("product_id=?", cart.ProductId).Update("num", cart.Num)
	}
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

func GetProductFromId(pid int) (model.Product, error) {
	var info model.Product
	result := DB.Model(&model.Product{}).Where("product_id=?", pid).Find(&info)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return info, nil
}

func GetProductFromType(ty string) ([]model.Product, error) {
	var info []model.Product
	result := DB.Model(&model.Product{}).Where("type=?", ty).Find(&info)
	if result.Error != nil {
		return []model.Product{}, result.Error
	}
	return info, nil
}

func GetProductFromName(name string) (model.Product, error) {
	var info model.Product
	result := DB.Model(&model.Product{}).Where("name=?", name).Find(&info)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return info, nil
}

func InCart(pid int, uid int) bool {
	var in int64
	DB.Model(&model.Cart{}).Where("product_id=? AND user_id=?", pid, uid).Count(&in)
	if in == 0 {
		return false
	}
	return true
}

func Order(uid int) (int64, error) {
	var oInfo model.OrderInfo
	var order model.Order
	var pInfo []model.Cart
	var praise float64
	result := DB.Model(&model.Cart{}).Where("user_id=?", uid).Find(&pInfo)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0, result.Error
	}
	for _, item := range pInfo {
		praise += item.Price * float64(item.Num)
	}
	order.Uid = uid
	order.Price = praise
	result = DB.Model(&model.Order{}).Create(&order)
	if result.Error != nil {
		return 0, result.Error
	}
	oInfo.OrderId = order.OrderId
	for _, item := range pInfo {
		oInfo.ProductId = item.ProductId
		oInfo.Price = item.Price
		oInfo.Num = item.Num
		result = DB.Model(&model.OrderInfo{}).Create(&oInfo)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, result.Error
		}
	}
	result = DB.Model(&model.Cart{}).Where("user_id=?", uid).Delete(&model.Cart{})
	if result.Error != nil {
		return 0, result.Error
	}
	return order.OrderId, nil
}
