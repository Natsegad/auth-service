package app

import (
	autreg "auth/auth-back/internal/delivery/http/register"
	authlogin "auth/auth-back/internal/delivery/login"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.Handle("GET", "/auth", autreg.AuthPage)
	r.Handle("POST", "/login", authlogin.LoginPage)

	r.Run(":8080")
}
