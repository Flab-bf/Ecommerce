package middleWares

import (
	"context"
	"ecommerce/dao"
	"ecommerce/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"strings"
)

func JwtAuthMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(404, "miss authorization")
			ctx.Abort()
			return
		}
		if strings.HasPrefix(authHeader, "Bearer ") {
			authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		}
		myClaims, err := utils.ParseToken(authHeader)
		if err != nil {
			ctx.JSON(404, "token 解析错误")
			ctx.Abort()
			return
		}
		ok, err := dao.IsLegalUser(myClaims.Uid)
		if !ok {
			ctx.JSON(100, "非法用户")
			ctx.Abort()
			return
		}
		ctx.Set("uid", myClaims.Uid)
		ctx.Next(c)
	}
}
