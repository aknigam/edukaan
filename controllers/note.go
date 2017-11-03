package controllers

import (
	"edukaan/common"
	"edukaan/models"
	"edukaan/repository"
	"encoding/json"
	"net/http"
)

type NoteController struct {
	repo repository.NoteRepository
}

var Note NoteController

func (controller *NoteController) AddNote(w http.ResponseWriter, r *http.Request) {

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var noteSnippet models.NoteSnippet
	json.Unmarshal(body, &noteSnippet)

	err := controller.repo.AddNote(&noteSnippet)
	if err != nil {
		common.Error.Println("Order could not be created", err)
		return
	}

	common.Info.Println("Call successfull", body)

	w.WriteHeader(http.StatusOK)
	// location header should also be set as per the REST standards
	return
}
