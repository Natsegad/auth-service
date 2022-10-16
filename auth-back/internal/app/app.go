package app

import (
	authapi "auth/auth-back/internal/delivery/http/api"
	autreg "auth/auth-back/internal/delivery/http/register"
	authlogin "auth/auth-back/internal/delivery/http/sigin"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.Handle("GET", "/auth", autreg.AuthPage)
	r.Handle("POST", "/login", authlogin.SignUpPage)
	r.Handle("GET", "/user-valid", authapi.ValidUser)

	r.Run(":8080")
}
