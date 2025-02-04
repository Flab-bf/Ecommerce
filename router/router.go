package router

import (
	"ecommerce/api"
	"ecommerce/middleWares"
	"ecommerce/utils"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func NewRouter() *server.Hertz {
	h := server.Default(server.WithTLS(utils.ConnectHttps()), server.WithHostPorts(":8080"))
	user := h.Group("/user")
	user.Use(middleWares.JwtAuthMiddleware())
	//用户注册
	h.POST("/user/register", api.UserRegister)
	//登录
	h.POST("/user/token", api.UserLogin)
	//刷新token
	//需要前端在token过期前发送刷新请求
	user.POST("/token/refresh", api.RefreshToken)
	//修改密码
	user.PUT("/password", api.UpdatePassword)
	//获取用户信息
	user.GET("/info/:user_id", api.GetUserInfo)
	//修改用户信息
	user.PUT("/info", api.ChangeUserInfo)

	product := h.Group("/product")
	product.Use(middleWares.JwtAuthMiddleware())
	//获取商品列表
	h.GET("/product/list", api.GetProductList)
	//搜索商品
	product.GET("/search", api.SearchProduct)
	//加入购物车
	product.PUT("/addCart", api.AddCart)
	//获取购物车列表
	product.GET("/cart", api.CartInfo)
	//获取商品详情
	product.GET("/info/:product_id", api.GetInfoFromId)
	//获取相应标签的商品列表
	product.GET("/:type", api.GetInfoFromType)

	comment := h.Group("/comment")
	comment.Use(middleWares.JwtAuthMiddleware())
	//获取评论
	h.GET("/comment/:product_id", api.GetComment)
	//评论,回复(有无父评论ID)
	comment.POST("/:product_id", api.Comment)
	//删除评论
	comment.DELETE("/:comment_id", api.DeleteComment)
	//更新评论
	comment.PUT("/:comment_id", api.UpdateComment)

	//点踩评论
	comment.POST("/praise", api.PraiseOrNot)

	//下单
	h.POST("/operate/order", middleWares.JwtAuthMiddleware(), api.GetOrder)

	return h
}
