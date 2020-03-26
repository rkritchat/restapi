package service

import (
	"net/http"
	"restapi/common"
	"restapi/entity"
	"restapi/model"
	"restapi/str"
)

func (h *Handlers) InquiryAll(w http.ResponseWriter, r *http.Request) {
	find := h.db.Find(&[]entity.UserInfo{})
	if find.Error != nil {
		common.RespErr(w, find.Error)
		return
	}
	common.RespJson(w,
		&model.InquiryAllRes{
			UserInfo:  find.Value.(*[]entity.UserInfo),
			CommonRes: common.InitRes(str.InquiryAllUserInfoSuccessfully, nil),
		})
}
