package login

import (
	serviceauth "auth/auth-back/internal/domain/auth"
	"auth/auth-back/internal/service/reguser"
	"auth/auth-back/internal/service/user"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Printf("Error parseform %s \n", err.Error())
		panic(nil)
	}

	userId := user.GenUserId()

	pass := c.Request.PostForm["password"]
	email := c.Request.PostForm["email"]

	userReq := serviceauth.AuthUserReq{
		Email:    email[0],
		Password: pass[0],
		Id:       userId,
	}

	reguser.WriteUserInfo(userReq)

	response := reguser.GenResponse(userId)

	jsonUser, err := json.Marshal(&response)
	if err != nil {
		fmt.Printf("Error %s \n", err.Error())
		panic(nil)
	}

	fmt.Printf("%s %s \n", email, pass)

	c.JSON(200, string(jsonUser))
}
