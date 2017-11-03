package repository

import (
	"edukaan/common"
	"edukaan/models"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type NoteRepository struct {
}

// Create a new noteSnippet
// Refer: https://stackoverflow.com/questions/7151261/append-to-a-file-in-go

func (repo *NoteRepository) AddNote(noteSnippet *models.NoteSnippet) (err error) {

	file, err := os.OpenFile("/Users/a.nigam/Documents/workspace/gospace/src/edukaan/temp.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		common.Error.Println("NoteSnippet could not be created ", err)
		return
	}

	defer file.Close()

	if _, err = file.WriteString("\n"); err != nil {
		panic(err)
	}

	if _, err = file.WriteString(noteSnippet.PageUrl + "\n"); err != nil {
		panic(err)
	}
	if _, err = file.WriteString(noteSnippet.Content + "\n"); err != nil {
		panic(err)
	}

	return
}
