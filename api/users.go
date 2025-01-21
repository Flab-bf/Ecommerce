package api

import (
	"context"
	"ecommerce/model"
	"ecommerce/service"
	"github.com/cloudwego/hertz/pkg/app"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var req model.UserMassage
	err := c.Bind(&req)
	if err != nil {
		c.JSON(100, "待定")
		return
	}
	err = service.RegisterUser(&req)
	if err != nil {
		c.JSON(100, "witing")
		return
	}
	c.JSON(200, "ok")
}

func UserLogin(ctx context.Context, c *app.RequestContext) {
	var req model.UserMassage
	err := c.Bind(&req)
	if err != nil {
		c.JSON(100, "witing")
		return
	}
	is := service.LoginUser(&req)
	if is == 0 {
		c.JSON(100, "witing")
	}
	c.JSON(200, "ok")
}
