package service

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"restapi/common"
	"restapi/entity"
	"restapi/model"
	"restapi/str"
)

func Register(w http.ResponseWriter, r *http.Request) {
	rq := initRq(r)
	if valid(rq, w) == nil {
		create(rq)
		common.Res(w, str.RegisterSuccessfully, nil)
	}
}

func initRq(r *http.Request) *model.RegisterReq {
	var rq model.RegisterReq
	json.NewDecoder(r.Body).Decode(&rq)
	return &rq
}

func valid(rq *model.RegisterReq, w http.ResponseWriter) error {
	v := validator.New()
	err := v.Struct(rq)
	if err != nil {
		common.Res(w, str.Empty, err)
		return err
	}
	return nil
}

func create(rq *model.RegisterReq){
	db := common.DBConnect()
	db.SingularTable(true)
	defer db.Close()
	db.Create(&entity.UserInfo{
		FirstName: rq.FirstName,
		LastName:  rq.LastName,
		Age:       rq.Age,
	})
	db.Commit()
}