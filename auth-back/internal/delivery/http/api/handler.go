package api

import (
	"auth/auth-back/internal/service"
	serviceuser "auth/auth-back/internal/service/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ValidUser(c *gin.Context) {
	bndJson := serviceuser.GetValidUserJson(c.Request.Body)
	if bndJson.Id == 0 {
		fmt.Printf("Error user.GetValidUserJson return 0 \n")
		c.JSON(400, gin.H{
			"found": false,
			"id":    0,
		})
		return
	}

	usersData, err := service.ReadUsers()
	if err != nil {
		fmt.Printf("Error %s \n", err.Error())
		return
	}

	user := usersData.Users[bndJson.Id]
	if user.Id != 0 {
		c.JSON(200, gin.H{
			"found": true,
			"id":    user.Id,
		})
		return
	}

	c.JSON(400, gin.H{
		"found": false,
		"id":    0,
	})
}
