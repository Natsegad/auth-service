package register

import (
	"auth/auth-back/internal/service/user"
	pkgjwt "auth/auth-back/pkg/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthPage(c *gin.Context) {
	fmt.Println("Auth page !")
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Printf("Error ParseForm %s \n", err.Error())
		return
	}

	email := c.Request.PostForm["email"][0]
	pass := c.Request.PostForm["password"][0]
	haveUser, data := user.CheckHaveUser(email, pass)
	if haveUser == true {
		claim, err := pkgjwt.ParseJwt(data.Token)
		if claim != nil && err == nil {
			c.JSON(200, data.Token)
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("User with email = %s pass = %s not found ! \n", email, pass)
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "User not found pls press sign-up",
	})
}
