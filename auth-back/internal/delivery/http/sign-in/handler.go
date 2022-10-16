package register

import (
	"auth/auth-back/internal/service/reguser"
	"auth/auth-back/internal/service/user"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthPage(c *gin.Context) {
	userId := user.GenUserId()
	response := reguser.GenResponse(userId)

	jsonUser, err := json.Marshal(&response)
	if err != nil {
		fmt.Printf("Error %s \n", err.Error())
		panic(nil)
	}

	c.JSON(200, string(jsonUser))
}
