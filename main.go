package main

import (
	"PaperSystem/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	initRouter()
}

func initRouter() {
	r:=gin.Default()
	r.Static("/static","./static")
	r.LoadHTMLGlob("templates/*")

	mainRouter:=&controllers.Mainrouter{}
	mainRouter.Initialize(r)
}
