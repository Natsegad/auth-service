package user

import (
	"auth/auth-back/internal/domain/auth"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/google/uuid"
)

func GenUserId() uint64 {
	return uint64(uuid.New().ID())
}

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
