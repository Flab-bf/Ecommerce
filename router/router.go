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
	//用户注册
	user.POST("/register", api.UserRegister)
	//登录
	user.POST("/token", api.UserLogin)
	//刷新token
	user.GET("/token/refresh", middleWares.JwtAuthMiddleware(), api.RefreshToken)
	//修改密码
	user.PUT("/password", middleWares.JwtAuthMiddleware(), api.UpdatePassword)
	//获取用户信息
	user.GET("/info/:user_id", middleWares.JwtAuthMiddleware(), api.GetUserInfo)
	//修改用户信息
	user.PUT("/info", middleWares.JwtAuthMiddleware(), api.ChangeUserInfo)

	product := h.Group("/product")
	product.Use(middleWares.JwtAuthMiddleware())
	//获取商品列表
	product.GET("/list", api.GetProductList)
	//搜索商品
	product.GET("/search")
	//加入购物车
	product.PUT("/addCart", api.AddCart)
	//获取购物车列表
	product.GET("/crat", api.CartInfo)
	//获取商品详情
	product.GET("/info/:product_id")
	//获取相应标签的商品列表
	product.GET("/:type")

	//comment := h.Group("/comment")
	////获取评论
	//comment.GET("/:product_id")
	////评论
	//comment.POST("/:product_id") //post用于创建新资源
	////删除评论
	//comment.DELETE("/:comment_id")
	////更新评论
	//comment.PUT("/:comment_id")
	////点踩评论
	//comment.POST("/praise")
	//
	////下单
	//h.POST("/operate/order")

	return h
}
