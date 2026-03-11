package router

import (
	api2 "ecommerce/backend/api"
	"ecommerce/backend/middleWares"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"os"
	"path/filepath"
	"time"
)

func NewRouter() *server.Hertz {
	//使用https
	//h := server.Default(server.WithTLS(utils.ConnectHttps()), server.WithHostPorts(":8080"))

	//使用http
	h := server.Default(server.WithHostPorts(":8080"))
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 开发环境允许所有来源，生产可指定前端IP:端口
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许所有请求方法
		AllowHeaders:     []string{"*"},                                       // 包含JWT的Authorization头
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,           // 允许前端携带cookie/token
		MaxAge:           12 * time.Hour, // 预检请求缓存，减少OPTIONS请求
	}))

	rootDir, _ := os.Getwd()
	staticDir := filepath.Join(rootDir, "../../frontend")
	h.Static("/", staticDir)

	user := h.Group("/user")
	user.Use(middleWares.JwtAuthMiddleware())
	//用户注册
	h.POST("/user/register", api2.UserRegister)
	//登录
	h.POST("/user/token", api2.UserLogin)
	//刷新token
	//需要前端在token过期前发送刷新请求
	user.POST("/token/refresh", api2.RefreshToken)
	//修改密码
	user.PUT("/password", api2.UpdatePassword)
	//获取用户信息
	user.GET("/info/:user_id", api2.GetUserInfo)
	//修改用户信息
	user.PUT("/info", api2.ChangeUserInfo)

	product := h.Group("/product")
	product.Use(middleWares.JwtAuthMiddleware())
	//获取商品列表
	h.GET("/product/list", api2.GetProductList)
	//搜索商品
	h.GET("/product/search", api2.SearchProduct)
	//加入购物车
	product.PUT("/addCart", api2.AddCart)
	//获取购物车列表
	product.GET("/cart", api2.CartInfo)
	//获取商品详情
	product.GET("/info/:product_id", api2.GetInfoFromId)
	//获取相应标签的商品列表
	product.GET("/:type", api2.GetInfoFromType)

	comment := h.Group("/comment")
	comment.Use(middleWares.JwtAuthMiddleware())
	//获取评论
	h.GET("/comment/:product_id", api2.GetComment)
	//评论,回复(有无父评论ID)
	comment.POST("/:product_id", api2.Comment)
	//删除评论
	comment.DELETE("/:comment_id", api2.DeleteComment)
	//更新评论
	comment.PUT("/:comment_id", api2.UpdateComment)

	//点踩评论
	comment.POST("/praise", api2.PraiseOrNot)

	//下单
	h.POST("/operate/order", middleWares.JwtAuthMiddleware(), api2.GetOrder)

	return h
}
