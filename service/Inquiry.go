package service

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"restapi/common"
	"restapi/entity"
	"restapi/model"
	"restapi/str"
	"strconv"
)

type InquiryTask struct {
	id    int
	err   error
	value model.InquiryRes
}

func (h *Handlers) Inquiry(w http.ResponseWriter, r *http.Request) {
	task := (&InquiryTask{}).initReqId(r)
	if task.err != nil {
		common.RespErr(w, task.err)
		return
	}
	rs := h.db.Find(&entity.UserInfo{Id: task.id})
	task.isFound(rs)
	task.initRes(w)
}

func (t *InquiryTask) initReqId(r *http.Request) *InquiryTask {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	t.id = id
	t.err = err
	return t
}

func (t *InquiryTask) isFound(rs *gorm.DB) {
	if rs.Error != nil {
		t.err = errors.New(str.InvalidUserID)
		return
	}
	t.value.UserInfo = rs.Value.(*entity.UserInfo)
}

func (t *InquiryTask) initRes(w http.ResponseWriter) {
	t.value.CommonRes = common.InitRes(str.InquirySuccessfully, t.err)
	common.RespJson(w, &t.value)
}
