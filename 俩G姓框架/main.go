package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router:=gin.Default()
	router.Use(cors.Default())

	router.POST("/Mon/register",Register)
	router.POST("/Mon/login", Login)
	router.POST("/Mon/query", Query)
	router.POST("/Mon/update",Meddle,Update)

	_ = router.Run(":8080")
}

type Bang struct {
	jwt.StandardClaims
	Id       string `form:"id" json:"id" xml:"idr"`
	Username     string `form:"username" json:"username" xml:"username"`
	Password string `form:"password" json:"password" xml:"password"`
}

func Meddle(c *gin.Context) {//中间件
	claim, err := CheckAction(SingedToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	c.Set("uid", claim.Id)
	c.Next()
	return
}