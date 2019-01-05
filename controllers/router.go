package controllers

import (
	"PaperSystem/models"
	"PaperSystem/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type Mainrouter struct {
	router *gin.Engine
	db *xorm.Engine
}
var Router *gin.Engine
func (self *Mainrouter) Initialize(r *gin.Engine) {
	store:=cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	self.router=r
	Router=r
	self.db=models.Db

	self.router.GET("/",self.IndexHandler)
	self.router.GET("/login",self.LoginViewHandler)
	self.router.POST("/logout",self.LogoutHandler)
	self.router.POST("/login",self.LoginHandler)
	self.router.POST("/signup",self.SignupHandler)
	self.router.GET("/signup",self.SignupViewHandler)
	self.router.POST("/newPaper",self.NewPaperHandler)
	self.router.GET("/newPaper",self.PaperViewHandler)
	self.router.GET("/papers/:Uuid",self.GetOnePaper)
	self.router.GET("/error/:msg",self.ErrorHandler)
	//self.router.POST("/authenticate",self.AuthenticateHandler)
	_ = self.router.Run(utils.APP_Address)

}




