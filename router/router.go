package router

import (
	"ecommerce/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func NewRouter() *server.Hertz {
	h := server.New()
	user := h.Group("/user")
	//用户注册
	user.POST("/register", api.UserRegister)
	//登录
	user.GET("/token", api.UserLogin)
	//刷新token
	user.GET("/token/refresh")
	//修改密码
	user.PUT("/password")
	//获取用户信息
	user.GET("/info/:user_id")
	//修改用户信息
	user.PUT("/info")

	product := h.Group("/product")
	//获取商品列表
	product.GET("/list")
	//加入购物车
	product.PUT("/addCart")
	//获取购物车列表
	product.GET("/crat")
	//获取商品详情
	product.GET("/info/:product_id")
	//获取相应标签的商品列表
	product.GET("/:type")

	comment := h.Group("/comment")
	//获取评论
	comment.GET("/:product_id")
	//评论
	comment.POST("/:product_id") //post用于创建新资源
	//删除评论
	comment.DELETE("/:comment_id")
	//更新评论
	comment.PUT("/:comment_id")
	//点踩评论
	comment.POST("/praise")

	//下单
	h.POST("/operate/order")

	//搜索商品
	h.GET("/:name/search")

	return h
}
