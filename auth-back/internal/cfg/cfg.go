package cfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Cfg struct {
	Host   string `json:"host"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Dbname string `json:"dbname"`
	Table  string `json:"table"`
}

const path = "../cfg.json"

var Config Cfg

func InitCfg() {
	ret := Cfg{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal(data, &ret)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Config = ret
}
