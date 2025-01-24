package api

import (
	"context"
	"ecommerce/service"
	"ecommerce/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
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

func GetInfoFromId(ctx context.Context, c *app.RequestContext) {
	pid := c.Param("product_id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(500, utils.ErrorResponse(10002, "发生意外错误"))
		return
	}
	info, err := service.SearchInfoFromId(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse(30003, "未知商品"))
		return
	}
	uid, _ := c.Get("uid")
	userId, ok := uid.(int)
	if !ok {
		c.JSON(500, utils.ErrorResponse(10002, "发生意外错误"))
	}
	service.Incart(userId, info)
	c.JSON(200, utils.SuccessResponse(info))
}

func GetInfoFromType(ctx context.Context, c *app.RequestContext) {
	ptype := c.Param("type")
	info, err := service.GetProductFromType(ptype)
	if err != nil {
		c.JSON(404, utils.ErrorResponse(30003, "未知商品"))
		return
	}
	uid, _ := c.Get("uid")
	userId, ok := uid.(int)
	if !ok {
		c.JSON(500, utils.ErrorResponse(10002, "发生意外错误"))
	}
	service.Incart(userId, info)
	c.JSON(200, utils.SuccessResponse(info))
}

func SearchProduct(c context.Context, ctx *app.RequestContext) {
	pname := ctx.Query("name")
	info, err := service.GetProductFromName(pname)
	if err != nil {
		ctx.JSON(404, utils.ErrorResponse(30003, "Not Found"))
		return
	}
	uid, is := ctx.Get("uid")
	if !is {
		info.IsAddedCart = false
	} else {
		userId, ok := uid.(int)
		if !ok {
			ctx.JSON(500, utils.ErrorResponse(10002, "发生意外错误"))
			return
		}
		service.Incart(userId, info)
	}
	ctx.JSON(200, utils.SuccessResponse(info))
}
