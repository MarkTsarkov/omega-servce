package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/marktsarkov/omega-service/internal/entity"
	"github.com/marktsarkov/omega-service/internal/service"
	"log"
	"strconv"
	"time"
)

func GetNoteById(service service.NoteService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var note *entity.Note

		err := c.BodyParser(&note)
		if err != nil {
			fmt.Println(err, "parser")
			c.Status(fiber.StatusBadRequest).SendString(err.Error()) // Возвращаем ошибку, если не смогли распарсить тело запроса
			return err
		}

		log.Println(note)
		id := note.ID
		log.Println(id)

		note, err = service.GetById(c.Context(), id)
		if err != nil {
			fmt.Println(err, "to note")
			c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // Возвращаем ошибку сервера, если что-то пошло не так
			return err
		}
		response := map[string]string{
			"id":         strconv.FormatInt(id, 10),
			"title":      note.Title,
			"body":       note.Body,
			"created_at": note.CreatedAt.Format(time.DateTime),
			"updated_at": note.UpdatedAt.Format(time.DateTime),
		}

		return c.JSON(response)
	}
}
