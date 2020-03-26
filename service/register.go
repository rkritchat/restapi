package service

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"restapi/common"
	"restapi/entity"
	"restapi/model"
	"restapi/str"
)

type RegisterTask struct {
	err error
	model.RegisterReq
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	task := (&RegisterTask{}).initRq(r)
	if err := task.valid(); err != nil {
		common.RespErr(w, err)
		return
	}
	task.create(h.db)
	common.RespCommon(w, str.RegisterSuccessfully, task.err)
}

func (t *RegisterTask) initRq(r *http.Request) *RegisterTask {
	json.NewDecoder(r.Body).Decode(&t)
	return t
}

func (t RegisterTask) valid() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.FirstName, validation.Required, validation.Length(1, 20)),
		validation.Field(&t.LastName, validation.Required, validation.Length(1, 20)),
		validation.Field(&t.Age, validation.Max(100)))
}

func (t *RegisterTask) create(db *gorm.DB) {
	t.err = db.Create(&entity.UserInfo{
		FirstName: t.FirstName,
		LastName:  t.LastName,
		Age:       t.Age,
	}).Error
}
