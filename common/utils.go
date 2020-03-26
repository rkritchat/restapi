package common

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"net/http"
	"restapi/model"
	"restapi/str"
)

func InitRes(s string, e error) model.CommonRes {
	if e != nil {
		return model.CommonRes{Status: "Failed", Desc: e.Error()}
	}
	return model.CommonRes{Status: "Success", Desc: s}
}

func RespCommon(w http.ResponseWriter, s string, e error) {
	w.Header().Set(str.ContentType, str.ApplicationJson)
	json.NewEncoder(w).Encode(InitRes(s, e))
}

func RespJson(w http.ResponseWriter, rs interface{}) {
	w.Header().Set(str.ContentType, str.ApplicationJson)
	json.NewEncoder(w).Encode(&rs)
}

func RespErr(w http.ResponseWriter, e error) {
	w.Header().Set(str.ContentType, str.ApplicationJson)
	json.NewEncoder(w).Encode(InitRes("", e))
}

func DBConnect() gorm.DB {
	db, err := gorm.Open("mysql", viper.GetString("db"))
	if err != nil {
		panic(err.Error())
	}
	db.SingularTable(true)
	return *db
}

func LoadConfig() {
	viper.SetConfigName(str.Config)
	viper.AddConfigPath(str.Dot)
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
}
