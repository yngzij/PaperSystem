package controllers

import (
	"PaperSystem/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func (router *Mainrouter) IndexHandler(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("userSession")

	fmt.Println(v)
	var user *models.User
	papers, err := models.GetAllPaper()
	for k, v := range papers {
		fmt.Println(k, v)
	}
	if err != nil {
		c.JSON(http.StatusOK, &models.Response{Success: false, Message: "Get Data Error "})
		return
	}

	if v != nil {
		user_temp, _ := models.UserByUUID(v.(string))
		user = &user_temp
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":   user,
		"papers": papers,
	})

}

func (r *Mainrouter) LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userSession")
	session.Save()
	c.Redirect(301, "/")
}

func (router *Mainrouter) ErrorHandler(c *gin.Context) {
	msg := c.Param("msg")
	c.HTML(200, "alert_msg.html", gin.H{
		"msg": msg,
	})
}

func (router *Mainrouter) GetOnePaper(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("userSession")
	uuid := c.Param("Uuid")
	var paper *models.Paper
	if v != nil {
		paper_temp, _ := models.PaperByUUID(uuid)
		fmt.Println(paper_temp)
		paper = &paper_temp
		c.HTML(http.StatusOK, "paper_show.html", gin.H{
			"paper": paper,
		})
	} else {
		c.Redirect(302, "/login")
	}
}

func (router *Mainrouter) NewPaperHandler(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("content")
	var images []string
	for i := 1; i <= 4; i++ {
		image:="image_"+strconv.Itoa(i)

		file, header, err := c.Request.FormFile(image)

		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			fmt.Println(err)
			return
		}

		//文件的名称
		filename := header.Filename
		fmt.Println(file, err, filename)
		//创建文件
		imgpath := "static/img/" + filename
		out, err := os.Create(imgpath)
		//注意此处的 static/uploadfile/ 不是/static/uploadfile/
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		imgpath="/"+imgpath
		images=append(images, imgpath)
		fmt.Println(images)
		_ = out.Close()
	}
	c.String(http.StatusCreated, "upload successful")

	paper := models.Paper{
		Title:     title,
		Body:      body,
		Uuid:      models.CreateUUID(),
		CreatedAt: time.Now(),
		Image_1:     images[0],
		Image_2:     images[1],
		Image_3:     images[2],
		Image_4:     images[3],
	}

	if err := paper.Create(); err != nil {
		c.JSON(http.StatusOK, &models.Response{Success: false, Message: "reg error"})
	}
	c.Redirect(http.StatusMovedPermanently, "/newPaper")

}

func (router *Mainrouter) SignupHandler(c *gin.Context) {
	accuont := c.PostForm("useraccount")
	pwd := c.PostForm("password")
	name := c.PostForm("username")
	cryptpwd := models.Encrypt(pwd)

	matched, _ := regexp.MatchString("^[a-zA-Z0-9][a-zA-Z0-9_.-]+$", accuont)

	if !matched {
		c.JSON(http.StatusOK, &models.Response{Success: false, Message: "Username is invalid!"})
		return
	}

	user := models.User{
		Name:      name,
		Password:  cryptpwd,
		Account:   accuont,
		CreatedAt: time.Now(),
		Uuid:      models.CreateUUID(),
	}

	if err := user.Create(); err != nil {
		c.JSON(http.StatusOK, &models.Response{Success: false, Message: "reg error"})
	}
	c.Redirect(http.StatusMovedPermanently, "/login")
}
