package api

import (
	"context"
	"ecommerce/service"
	"ecommerce/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func GetProductList(ctx context.Context, c *app.RequestContext) {
	date, err := service.ProductList()
	if err != nil {
		c.JSON(500, utils.ErrorResponse(10002, "服务器发生错误"))
		return
	}
	c.JSON(200, utils.SuccessResponse(date))
}

func AddCart(ctx context.Context, c *app.RequestContext) {
	var productId map[string]int
	err := c.Bind(&productId)
	if err != nil {
		c.JSON(400, utils.ErrorResponse(10001, "参数错误"))
		return
	}
	uid, ok := c.Get("uid")
	if !ok {
		c.JSON(500, utils.ErrorResponse(10002, "发生了意料之外的错误"))
		return
	}
	userId, ok := uid.(int)
	if !ok {
		c.JSON(500, utils.ErrorResponse(10002, "发生了意料之外的错误"))
		return
	}
	err = service.AddCart(productId["product_id"], userId)
	if err != nil {
		c.JSON(400, utils.ErrorResponse(30002, "添加失败"))
		return
	}
	c.JSON(201, utils.SuccessResponse(nil))
}

func CartInfo(ctx context.Context, c *app.RequestContext) {
	uid, ok := c.Get("uid")
	if !ok {
		c.JSON(500, utils.ErrorResponse(10002, "发生了意料之外的错误"))
		return
	}
	userId, ok := uid.(int)
	if !ok {
		c.JSON(500, utils.ErrorResponse(10002, "发生了意料之外的错误"))
		return
	}
	date, err := service.GetCarts(userId)
	if err != nil {
		c.JSON(500, utils.ErrorResponse(30003, "获取失败"))
		return
	}
	c.JSON(200, utils.SuccessResponse(date))
}
