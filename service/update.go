package service

import (
	"encoding/json"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"net/http"
	"restapi/common"
	"restapi/entity"
	"restapi/model"
	"restapi/rule"
	"restapi/str"
)

type UpdateTask struct {
	err error
	model.UpdateReq
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	task := (&UpdateTask{}).initRq(r)
	task.Validate()
	if task.err != nil {
		common.RespErr(w, task.err)
		return
	}
	task.Update(h.db)
	common.RespCommon(w, str.UpdateSuccessfully, task.err)
}

func (t *UpdateTask) initRq(r *http.Request) *UpdateTask {
	json.NewDecoder(r.Body).Decode(&t)
	return t
}

func (t *UpdateTask) Validate() {
	t.err = validation.ValidateStruct(t,
		validation.Field(&t.Id, rule.Id...),
		validation.Field(&t.FirstName, rule.FirstName...),
		validation.Field(&t.LastName, rule.LastName...),
		validation.Field(&t.Age, rule.Age...))
}

func (t *UpdateTask) Update(db *gorm.DB) {
	rs := db.First(&entity.UserInfo{Id: t.Id})
	if rs.Error != nil {
		t.err = errors.New(str.UserIDIsNotFound)
		return
	}
	t.err = rs.Update(&entity.UserInfo{
		FirstName: t.FirstName,
		LastName:  t.LastName,
		Age:       t.Age}).Error
}
