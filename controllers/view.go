package controllers

import (
	"PaperSystem/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)


func (router *Mainrouter) LoginViewHandler(c *gin.Context){
	c.HTML(http.StatusOK,"login.html",gin.H{
	})
}

func (router *Mainrouter)SignupViewHandler(c *gin.Context) {
	c.HTML(http.StatusOK,"signup.html",gin.H{
	})
}


func (router *Mainrouter)PaperViewHandler(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("userSession")
	var user *models.User
	if v != nil {
		user_temp,_:= models.UserByUUID(v.(string))
		user=&user_temp
		c.HTML(http.StatusOK,"paper.html",gin.H{
			"user":user,
		})
	}else{
		c.Redirect(302,"/")
	}

}


func (r *Mainrouter)LoginHandler(c *gin.Context) {
	fmt.Println("postpost sig")
	account := c.PostForm("useraccount")
	pwd := c.PostForm("password")
	print(account,pwd)
	//cryptpwd:=models.Encrypt(pwd)
	user, err := models.UserByAccount(account)
	if err != nil {
		//c.JSON(http.StatusOK, models.Response{Success: false, Message: "Error !"})
		c.HTML(200,"/error/account not find !",gin.H{
		})
	}

	if user.Password == models.Encrypt(pwd) {
		//store, _ := redis.NewStore(10, "tcp", "119.23.214.1223:6379", "ta", []byte("_cookie"))
		session:=sessions.Default(c)
		store:=cookie.NewStore([]byte("secret"))
		r.router.Use(sessions.Sessions("mySession", store))

		session.Set("userSession",user.Uuid)
		session.Save()
		c.Redirect(http.StatusMovedPermanently,"/")
	}else {
		c.Redirect(http.StatusMovedPermanently, "/error/password error ")
	}
}

