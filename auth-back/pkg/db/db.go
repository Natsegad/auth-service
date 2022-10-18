package db

import (
	"auth/auth-back/internal/cfg"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserInfo struct {
	Id           uint64 `json:"id" gorm:"primaryKey"`
	Email        string `json:"email"`
	Pass         string `json:"password"`
	Token        string `json:"jwt"`
	RefreshToken string `json:"refresh_jwt"`
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

func AddUser(user UserInfo) {
	ok := DataBase.Create(&user)
	if ok.Error != nil {
		fmt.Printf("Error AddUser %s \n", ok.Error.Error())
		return
	}

	fmt.Printf("User added to database %s \n", user.Email)
}

func DeleteUser(id uint64) error {
	user := UserInfo{}
	ok := DataBase.Find(&user, id)
	if ok.Error != nil {
		fmt.Printf("Error AddUser %s \n", ok.Error.Error())
		return ok.Error
	}

	DataBase.Delete(&user)

	return nil
}

func GetUser(id uint64) (UserInfo, error) {
	user := UserInfo{}
	ok := DataBase.Find(&user, id)
	if ok.Error != nil {
		fmt.Printf("Error AddUser %s \n", ok.Error.Error())
		return user, ok.Error
	}

	return user, nil
}

func GetAllUsers() ([]UserInfo, error) {
	var users []UserInfo
	ok := DataBase.Find(&users)
	if ok.Error != nil {
		fmt.Printf("Error AddUser %s \n", ok.Error.Error())
		return nil, ok.Error
	}

	return users, nil
}

func GetUserByEmail(email string) UserInfo {
	ret := UserInfo{}
	users, err := GetAllUsers()
	if err != nil {
		fmt.Printf("Error GetAllUsers in GetUserByEmail \n")
		return ret
	}

	for _, v := range users {
		if v.Email == email {
			return v
		}
	}

	return ret
}

func UpdateUser(user UserInfo) error {
	userBase := UserInfo{}
	ok := DataBase.Find(&userBase, user.Id)
	if ok.Error != nil {
		fmt.Printf("Error AddUser %s \n", ok.Error.Error())
		return ok.Error
	}

	userBase.Email = user.Email
	userBase.Pass = user.Pass
	userBase.Token = user.Token

	DataBase.Save(&userBase)

	return nil
}
