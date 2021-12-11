package global

import (
	"encoding/json"
	"github.com/kanyuanzhi/tialloy/global"
	"io/ioutil"
)

var CustomObject *CustomObj

type CustomObj struct {
	*global.Obj
	MysqlUsername string `json:"mysql_username,omitempty"`
	MysqlPassword string `json:"mysql_password,omitempty"`
	MysqlHost     string `json:"mysql_host,omitempty"`
	MysqlPort     uint   `json:"mysql_port,omitempty"`
	MysqlDbname   string `json:"mysql_dbname,omitempty"`
}

func (c *CustomObj) Reload() {
	data, err := ioutil.ReadFile("conf/tialloy.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &CustomObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	CustomObject = &CustomObj{
		Obj:           global.Object,
		MysqlUsername: "admin",
		MysqlPassword: "c3i123456",
		MysqlHost:     "127.0.0.1",
		MysqlPort:     3306,
		MysqlDbname:   "web_service",
	}
	CustomObject.Reload()
	DB = initDB()
}
