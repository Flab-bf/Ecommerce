package api

import (
	"context"
	"ecommerce/dao"
	"ecommerce/model"
	"ecommerce/service"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
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
	is, t := service.LoginUser(&req)
	if is != 1 {
		c.JSON(100, "witing")
		return
	}
	c.JSON(200, t)
}

func RefreshToken(ctx context.Context, c *app.RequestContext) {
	var req model.UserMassage
	err := c.Bind(&req)
	if err != nil {
		c.JSON(100, "witing")
		return
	}
	id, err := dao.FindUidFromAccount(req.Account)
	if err != nil {
		c.JSON(100, "witing")
		return
	}
	t := dao.PutTokenJwt(id)
	c.JSON(200, t)
}

func UpdatePassword(ctx context.Context, c *app.RequestContext) {
	var req model.UserChangePassword
	err := c.Bind(&req)
	if err != nil {
		c.JSON(100, "witing")
		return
	}
	err = service.ChangePassword(&req)
	if err != nil {
		c.JSON(100, "witing")
		return
	}
	c.JSON(200, "ok")
}

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var req model.UserMassage
	accountString := c.Param("user_id")
	var err error
	req.Account, err = strconv.Atoi(accountString)
	if err != nil {

	}
	info, err := service.GetUserInfo(req.Account)
	if err != nil {
		c.JSON(100, "witing")
		return
	}
	c.JSON(200, info)
}

func ChangeUserInfo(ctx context.Context, c *app.RequestContext) {
	var req model.UserMassage
	err := c.Bind(&req)
	if err != nil {
		c.JSON(100, "witing")
		return
	}
	err = service.ChangeUserInfo(&req)
	if err != nil {
		c.JSON(100, "witing")
		return
	}
	c.JSON(200, "ok")
}
