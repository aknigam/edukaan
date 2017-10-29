package controllers

import (
	"edukaan/common"
	"net/http"
)

type NoteController struct {
}

var Note NoteController

func (controller *NoteController) AddNote(w http.ResponseWriter, r *http.Request) {

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	common.Info.Println("Call successfull", body)

	w.WriteHeader(http.StatusOK)
	// location header should also be set as per the REST standards
	return
}
