package reguser

import (
	"auth/auth-back/internal/domain/auth"
	pkgjwt "auth/auth-back/pkg/jwt"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GenResponse(id uint64) auth.AuthUserRes {
	ret := auth.AuthUserRes{}
	ret.Id = id
	ret.Token = pkgjwt.GenerateJwtById(id)
	return ret
}

func GenResponseJWT(id uint64, token string) auth.AuthUserRes {
	ret := auth.AuthUserRes{}
	ret.Id = id
	ret.Token = token
	return ret
}

func WriteUserInfo(info auth.AuthUserReq) {
	file, err := ioutil.ReadFile("user.json")
	if err != nil {
		fmt.Printf("Dont write use info %s %s \n", info.Email, err.Error())
		return
	}
	usersJson := auth.UsersJson{}
	if len(file) != 0 {
		err = json.Unmarshal(file, &usersJson)
		if err != nil {
			fmt.Printf("err = json.Unmarshal(file,&usersJson) %s \n", err.Error())
		}
	} else {
		usersJson.Users = make(map[uint64]auth.AuthUserReq)
	}

	usersJson.Users[info.Id] = info

	jsonStr, err := json.Marshal(&usersJson)

	err = ioutil.WriteFile("user.json", jsonStr, 0)
	if err != nil {
		fmt.Printf("Dont write use info %s %s \n", info.Email, err.Error())
		return
	}
}
