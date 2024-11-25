package note

import (
	repository "github.com/marktsarkov/omega-service/internal/repo"
	"github.com/marktsarkov/omega-service/internal/service"
)

type serv struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) service.NoteService {
	return &serv{repo: repo}
}
