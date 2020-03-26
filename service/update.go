package service

import (
	"encoding/json"
	"net/http"
	"restapi/model"
)

type UpdateTask struct {
	value model.UpdateReq
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	var rq UpdateTask
	rq.initRq(r)
}

func (rq UpdateTask) initRq(r *http.Request) *UpdateTask {
	json.NewDecoder(r.Body).Decode(&rq)
	return &rq
}
