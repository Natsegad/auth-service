package app

import (
	"auth/auth-back/internal/cfg"
	authapi "auth/auth-back/internal/delivery/http/api"
	autreg "auth/auth-back/internal/delivery/http/sign-in"
	authlogin "auth/auth-back/internal/delivery/http/sigup"
	"auth/auth-back/pkg/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	cfg.InitCfg()
	db.InitDb()
	db.AutoMigrate()
	fmt.Println(cfg.Config.Host)

	r := gin.Default()

	r.Handle("POST", "/auth", autreg.AuthPage)
	r.Handle("POST", "/login", authlogin.SignUpPage)
	r.Handle("GET", "/user-valid", authapi.ValidUser)

	r.Run(":8080")
}
