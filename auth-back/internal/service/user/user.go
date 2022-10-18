package user

import (
	"auth/auth-back/internal/domain/auth"
	"auth/auth-back/internal/global"
	"auth/auth-back/internal/service"
	"auth/auth-back/pkg/db"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/google/uuid"
)

func GenUserId() uint64 {
	return uint64(uuid.New().ID())
}

// Нужно передать Body с json
func GetValidUserJson(body io.ReadCloser) auth.ValidUserJsonReq {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	bndJson := auth.ValidUserJsonReq{}
	err = json.Unmarshal(data, &bndJson)
	if err != nil {
		fmt.Println(err.Error())
	}

	return bndJson
}

// Проверяет есть ли юзер в базе
func CheckHaveUser(email, pass string) (bool, *db.UserInfo) {
	user := db.GetUserByEmail(email)
	if user.Id == 0 {
		return false, nil
	} else {

		decrptPass := service.EncryptDecrypt(user.Pass, global.XorKey)
		if decrptPass == pass {
			return true, &user
		}
	}
	// data, err := service.ReadUsers()
	// if err != nil {
	// 	fmt.Printf("Error ReadUsers %s \n", err.Error())
	// 	return false, nil
	// }

	// for _, v := range data.Users {
	// 	v.Password = service.EncryptDecrypt(v.Password, global.XorKey)
	// 	if v.Email == email && v.Password == pass {
	// 		return true, &v
	// 	}
	// }

	return false, nil
}
