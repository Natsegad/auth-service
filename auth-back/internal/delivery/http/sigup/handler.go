package login

import (
	// serviceauth "auth/auth-back/internal/domain/auth"
	"auth/auth-back/internal/global"
	"auth/auth-back/internal/service"

	// "auth/auth-back/internal/service/reguser"
	"auth/auth-back/internal/service/user"
	"auth/auth-back/pkg/db"
	pkgjwt "auth/auth-back/pkg/jwt"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SignUpPage(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Printf("Error parseform %s \n", err.Error())
		panic("Error parseform ")
	}

	userId := user.GenUserId()

	pass := c.Request.PostForm["password"][0]
	email := c.Request.PostForm["email"][0]

	user := db.UserInfo{
		Id:    userId,
		Email: email,
		Pass:  service.EncryptDecrypt(pass, global.XorKey),
		Token: pkgjwt.GenerateJwtById(userId),
	}
	db.AddUser(user)

	// data, err := service.ReadUsers()
	// if err != nil && err.Error() != "unexpected end of JSON input" {
	// 	fmt.Printf("Error read user %s \n", err.Error())
	// 	c.JSON(401, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// for _, v := range data.Users {
	// 	if v.Email == email {
	// 		c.JSON(401, gin.H{
	// 			"email": v.Email,
	// 			"msg":   "user is registered",
	// 		})
	// 		return
	// 	}
	// }

	// userReq := serviceauth.AuthUserReq{
	// 	Email:    email,
	// 	Password: pass,
	// 	Id:       userId,
	// 	Token:    pkgjwt.GenerateJwtById(userId),
	// }

	// reguser.WriteUserInfo(userReq)

	// response := reguser.GenResponseJWT(userId, userReq.Token)

	// jsonUser, err := json.Marshal(&response)
	// if err != nil {
	// 	fmt.Printf("Error %s \n", err.Error())
	// 	panic(err.Error())
	// }

	// fmt.Printf("%s %s \n", email, pass)

	c.JSON(200, user.Token)
}
