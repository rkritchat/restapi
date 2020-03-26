package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"restapi/common"
	"restapi/path"
	"restapi/service"
)

func main() {
	common.LoadConfig()
	db := common.DBConnect()
	defer db.Close()

	h := service.NewHandle(&db)
	r := addHandler(h)
	panic(http.ListenAndServe(":8080", r))
}

func addHandler(h *service.Handlers) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(path.RegisterPath, h.Register).Methods(http.MethodPost)
	r.HandleFunc(path.UpdatePath, h.Update).Methods(http.MethodPost)
	r.HandleFunc(path.InquiryPath, h.Inquiry).Methods(http.MethodGet)
	r.HandleFunc(path.InquiryAllPath, h.InquiryAll).Methods(http.MethodGet)
	return r
}
