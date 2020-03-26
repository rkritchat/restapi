package model

import "restapi/entity"

type CommonRes struct {
	Status string `json:"status"`
	Desc   string `json:"desc"`
}

type InquiryRes struct {
	CommonRes
	UserInfo *entity.UserInfo `json:"userInfo"`
}

type InquiryAllRes struct {
	CommonRes
	UserInfo *[]entity.UserInfo `json:"userInfos"`
}
