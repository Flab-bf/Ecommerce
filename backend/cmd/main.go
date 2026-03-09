package main

import (
	"ecommerce/backend/dao"
	"ecommerce/backend/router"
)

func main() {
	dao.ConnectDB()
	h := router.NewRouter()
	h.Spin()
}
