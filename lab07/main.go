package main

import (
	"lab07/model"
	"log"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v3"
	"github.com/mailru/easyjson"
)

var (
	notes      = []model.Note{}
	nextID     = 1
	notesMutex sync.RWMutex
)

func main() {
	app := fiber.New()

	app.Get("/", getNotes)
	app.Get("/:id", getNote)
	app.Post("/", createNote)
	app.Put("/:id", updateNote)
	app.Delete("/:id", deleteNote)

	log.Fatal(app.Listen(":3000"))
}

func getNote(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некоректний ID"})
	}

	notesMutex.RLock()
	defer notesMutex.RUnlock()

	for _, note := range notes {
		if note.ID == id {
			return c.JSON(note)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Нотатку не знайдено"})
}

func getNotes(c fiber.Ctx) error {
	notesMutex.RLock()
	defer notesMutex.RUnlock()

	// Якщо нотаток ще немає, повертаємо порожній масив замість null
	if len(notes) == 0 {
		return c.JSON([]model.Note{})
	}
	return c.JSON(notes)
}

func createNote(c fiber.Ctx) error {
	var note model.Note

	// Парсимо тіло запиту за допомогою вбудованого механізму Fiber
	// Завдяки згенерованому коду easyjson, під капотом виконається UnmarshalJSON
	if err := easyjson.Unmarshal(c.Body(), &note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некоректний JSON body"})
	}

	notesMutex.Lock()
	note.ID = nextID
	nextID++
	notes = append(notes, note)
	notesMutex.Unlock()

	return c.Status(fiber.StatusCreated).JSON(note)
}

func updateNote(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некоректний ID"})
	}

	var updatedNote model.Note
	if err := easyjson.Unmarshal(c.Body(), &model.Note{}); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некоректний JSON body"})
	}

	notesMutex.Lock()
	defer notesMutex.Unlock()

	for i, note := range notes {
		if note.ID == id {
			// Зберігаємо оригінальний ID, оновлюємо поля
			updatedNote.ID = id
			notes[i] = updatedNote
			return c.JSON(updatedNote)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Нотатку не знайдено"})

}

func deleteNote(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некоректний ID"})
	}

	notesMutex.Lock()
	defer notesMutex.Unlock()

	for i, note := range notes {
		if note.ID == id {
			// Видаляємо елемент зі зрізу
			notes = append(notes[:i], notes[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent) // 204 No Content
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Нотатку не знайдено"})
}
