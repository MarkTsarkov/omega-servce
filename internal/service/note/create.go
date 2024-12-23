package note

import (
	"context"
	"github.com/marktsarkov/omega-service/internal/entity"
	"strings"
)

func (s *serv) Create(ctx context.Context, note *entity.Note) (int64, error) {
	title := []string{note.Title, "MADE WITH 1xbet"}
	note.Title = strings.Join(title, " ")

	body := []string{note.Body, "BIG PRIZES: 1xbet"}
	note.Body = strings.Join(body, " ")

	id, err := s.repo.Create(ctx, note)
	if err != nil {
		return 0, err
	}

	return id, nil
}
