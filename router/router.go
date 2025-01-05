package router

import "github.com/cloudwego/hertz/pkg/app/server"

func NewRouter() *server.Hertz {
	h := server.New()
	return h
}
