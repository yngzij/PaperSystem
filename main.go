package main

import (
	"PaperSystem/controllers"
	"fmt"
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

func indexHandle(context *gin.Context) {
	fmt.Println(context)
}
/*

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		fmt.Println("get")
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}
*/