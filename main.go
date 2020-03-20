package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"restapi/path"
	"restapi/service"
	"restapi/str"
)

func main() {
	config()
	r := mux.NewRouter()
	r.HandleFunc(path.RegisterPath, service.Register).Methods(http.MethodPost)
	//http.Handle(path.ApiPrefixPath, r)
    http.ListenAndServe(":8080", r)
}

func config() {
	viper.SetConfigName(str.Config)
	viper.AddConfigPath(str.Dot)
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
}

