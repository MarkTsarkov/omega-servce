package handlers

import (
	"context"
	pb "github.com/marktsarkov/omega-service/internal/controller/grpc"
	"github.com/marktsarkov/omega-service/internal/entity"
)

func (s pb.NoteServer) GetByID(ctx context.Context, id int64) (note *entity.Note, err error) {

}
