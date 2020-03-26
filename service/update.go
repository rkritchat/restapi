package service

import (
	"encoding/json"
	"net/http"
	"restapi/model"
)

type UpdateReq struct {
	value model.UpdateReq
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	var rq UpdateReq
	rq.initRq(r)
}

func (rq UpdateReq) initRq(r *http.Request) *UpdateReq {
	json.NewDecoder(r.Body).Decode(&rq)
	return &rq
}
