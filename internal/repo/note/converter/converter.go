package converter

import (
	"github.com/marktsarkov/omega-service/internal/entity"
	"github.com/marktsarkov/omega-service/internal/repo/note/model"
	"time"
)

func ToNoteFromRepo(note *model.Note) *entity.Note {
	var updateAt time.Time
	if note.UpdatedAt.Valid {
		updateAt = note.UpdatedAt.Time
	}

	return &entity.Note{
		ID:        note.ID,
		Title:     note.Title,
		Body:      note.Body,
		CreatedAt: note.CreatedAt,
		UpdatedAt: updateAt,
	}
}
