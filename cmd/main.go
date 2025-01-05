package main

import (
	"ecommerce/dao"
	"ecommerce/router"
)

func main() {
	dao.ConnectDB()
	h:=router.NewRouter()
	h.Spin()
}
