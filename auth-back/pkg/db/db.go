package db

import (
	"auth/auth-back/internal/cfg"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserInfo struct {
	Id    uint64 `json:"id" gorm:"primaryKey"`
	Email string `json:"email"`
	Pass  string `json:"password"`
}

var DataBase *gorm.DB

func InitDb() {
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", cfg.Config.User, cfg.Config.Pass, "5432", cfg.Config.Dbname)
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9920 sslmode=disable TimeZone=Asia/Shanghai")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error initialize DB %s \n", err.Error())
		return
	}
	fmt.Println("DB initialize!")
	DataBase = db
}

func AutoMigrate() {
	err := DataBase.AutoMigrate(&UserInfo{})
	if err != nil {
		fmt.Printf("Error migrate DB %s \n", err.Error())
		return
	}
}
