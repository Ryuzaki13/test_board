package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine
var connection *sql.DB

const connectionString = "host=127.0.0.1 port=5432 user=postgres password=1234 dbname=board sslmode=disable"

func main() {
	var e error
	connection, e = sql.Open("postgres", connectionString)
	if e != nil {
		fmt.Println(e)
		return
	}

	router = gin.Default()
	router.Static("/assets/", "front/")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", handlerIndex)
	router.GET("/registration", handlerRegistration)
	router.GET("/authorization", handlerAuthorization)
	router.POST("/user/reg", handlerUserRegistration)
	router.POST("/user/auth", handlerUserAuthorization)
	_ = router.Run(":8080")
}

// pkg.go.dev/text/template
func handlerIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"Role": "manager",
	})
}

func handlerRegistration(c *gin.Context) {
	c.HTML(200, "registration.html", gin.H{})
}

func handlerAuthorization(c *gin.Context) {
	c.HTML(200, "authorization.html", gin.H{})
}

func handlerUserRegistration(c *gin.Context) {

	var user User

	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

	e = user.Create()
	if e != nil {
		c.JSON(200, gin.H{
			"Error": "Не удалось зарегистрировать пользователя",
		})
		return
	}

	c.JSON(200, gin.H{
		"Error": nil,
	})
}

func handlerUserAuthorization(c *gin.Context) {
	var user User
	e := c.BindJSON(&user)
	if e != nil {
		c.JSON(200, gin.H{
			"Error": e.Error(),
		})
		return
	}

	e = user.Select()
	if e != nil {
		c.JSON(200, gin.H{
			"Error": "Не удалось авторизоваться",
		})
		return
	}

	c.JSON(200, gin.H{
		"Error": nil,
	})
}
