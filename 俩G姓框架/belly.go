package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Register (c *gin.Context){
	var f Bang
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(200,gin.H{"code":10001,"message":"绑定失败"})
		return
	}

	var user Mon
	db.Model(&Mon{}).Where("id = ?",f.Id).First(&user)
	if user.Id!=0{
		c.JSON(http.StatusOK, gin.H{"code": 10001, "message": "用户已注册"})
		return
	}

	user = Mon{
        Password : f.Password,
        Username : f.Username,
	}
	user.Id,_ =strconv.Atoi(f.Id)
	if err :=  db.Model(&Mon{}).Create(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "message": "数据库报错","question":"三个参数都有值吗"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": "注册成功!"})
}

func Login (c *gin.Context){
	var f Bang
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(200,gin.H{"code":10001,"message":"绑定失败"})
		return
	}

	var user Mon
	user.Id,_=strconv.Atoi(f.Id)
	user.Password = f.Password

	db.Model(&Mon{}).Where(&Mon{Id: user.Id, Password: user.Password}).First(&user)
	if user.Username == "" {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "message": "密码错误"})
		return
	}

	//Token=json_web_token.Create(user.Username,user.Id)
	claims := &Bang{
		Id: f.Id,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(3600)).Unix()
	var err error
	SingedToken, err = CreateToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": "登录成功!","token":SingedToken})
}

func Query (c *gin.Context){
	var f Bang
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(200,gin.H{"code":10001,"message":"绑定失败"})
		return
	}

	var user Mon
	user.Id,_=strconv.Atoi(f.Id)

	db.Model(&Mon{}).Where("id = ?",user.Id).First(&user)
	if user.Username == "" {
		c.JSON(http.StatusOK, gin.H{"code": 10001, "message": "并无此人"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 10000, "id": user.Id,"username":user.Username,"password":user.Password})
}

func Update (c *gin.Context){
	var f Bang
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(200,gin.H{"code":10001,"message":"绑定失败"})
		return
	}

	var user Mon
	user = Mon{//没在里面的参数会为空，所以id的赋值放到后面
		Password : f.Password,
		Username : f.Username,
	}
	id,_:=c.Get("uid")
	fmt.Println(id)
	user.Id,_=strconv.Atoi(id.(string))
	fmt.Println(user.Id)

	if user.Username != ""{db.Model(Mon{}).Where("id = ?",user.Id).Update("username", user.Username)}
	if user.Password != ""{db.Model(Mon{}).Where("id = ?",user.Id).Update("password", user.Password)}

	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": "更新成功!"})
}