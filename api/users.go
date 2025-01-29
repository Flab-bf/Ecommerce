package api

import (
	"context"
	"ecommerce/dao"
	"ecommerce/model"
	"ecommerce/service"
	"ecommerce/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var req model.UserChangePassword
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, utils.ErrorResponse(10001, "参数错误"))
		return
	}
	err = service.RegisterUser(&req)
	if err != nil {
		c.JSON(409, utils.ErrorResponse(20001, "用户已存在"))
		return
	}
	c.JSON(201, utils.SuccessResponse(nil))
}

func UserLogin(ctx context.Context, c *app.RequestContext) {
	var umsg model.UserChangePassword
	var req model.UserMassage
	err := c.Bind(&umsg)
	if err != nil {
		c.JSON(400, utils.ErrorResponse(10001, "参数错误"))
		return
	}
	req.Account = umsg.Account
	req.Password = umsg.Password
	err, t, rt := service.LoginUser(&req)
	if err != nil {
		c.JSON(401, utils.ErrorResponse(20003, "用户名或密码错误"))
		return
	}
	c.JSON(200, utils.SuccessResponse(map[string]interface{}{
		"refreshToken": rt,
		"token":        t,
	}))
}

func RefreshToken(ctx context.Context, c *app.RequestContext) {
	var refreshToken model.UserToken
	err := c.Bind(&refreshToken)
	if err != nil {
		c.JSON(400, utils.ErrorResponse(10002, "参数错误"))
		return
	}
	myClaims, err := utils.ParseRefreshToken(refreshToken.RefreshToken)
	if err != nil {
		c.JSON(500, utils.ErrorResponse(10002, "解析错误"))
		return
	}
	t := dao.PutTokenJwt(myClaims.Uid)
	if t == "" {
		c.JSON(401, utils.ErrorResponse(20003, "登录过期"))
		return
	}
	c.JSON(200, utils.SuccessResponse(map[string]interface{}{
		"refresh_token": refreshToken.RefreshToken,
		"token":         t,
	}))
}

func UpdatePassword(ctx context.Context, c *app.RequestContext) {
	var req model.UserChangePassword
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, utils.ErrorResponse(10001, "参数错误"))
		return
	}
	err = service.ChangePassword(&req)
	if err != nil {
		c.JSON(500, utils.ErrorResponse(20005, "密码修改成功"))
		return
	}
	c.JSON(200, utils.SuccessResponse(nil))
}

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var req model.UserMassage
	accountString := c.Param("user_id")
	var err error
	req.Account, err = strconv.Atoi(accountString)
	if err != nil {
		c.JSON(500, utils.ErrorResponse(10002, "服务器错误"))
		return
	}
	info, err := service.GetUserInfo(req.Account)
	if err != nil {
		c.JSON(404, utils.ErrorResponse(20006, "未能找到用户信息"))
		return
	}
	c.JSON(200, utils.SuccessResponse(map[string]interface{}{
		"date": info,
	}))
}

func ChangeUserInfo(ctx context.Context, c *app.RequestContext) {
	var req model.UserMassage
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, utils.ErrorResponse(10001, "参数错误"))
		return
	}
	uid, _ := c.Get("uid")
	req.Uid = uid.(int)
	err = service.ChangeUserInfo(&req)
	if err != nil {
		c.JSON(400, utils.ErrorResponse(20004, "用户信息更新失败"))
		return
	}
	c.JSON(200, utils.SuccessResponse(nil))
}
