package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"lab11/model"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mailru/easyjson"
)

var db *sql.DB

func main() {
	defer db.Close()
	connStr := "postgres://postgres:postgres@db:5432/contacts_db?sslmode=disable"

	var err error
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	for i := 1; i <= 5; i++ {
		fmt.Printf("Connecting to database (attempt %d/5)...\n", i)
		err = db.Ping()
		if err == nil {
			break
		}

		if i == 5 {
			log.Fatalf("Unable to connect to PostgreSQL after 5 attempts: %v", err)
		}

		time.Sleep(3 * time.Second)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Unnable connect to PostgreSQL: %v", err)
	}

	migrationQuery, err := os.ReadFile("postgres/0001_create_contact.up.sql")
	if err != nil {
		log.Printf("Could not find init migration file: %v", err)
		return
	}

	_, err = db.Exec(string(migrationQuery))
	if err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	fmt.Println("Migration success")

	app := fiber.New(fiber.Config{
		JSONEncoder: func(v interface{}) ([]byte, error) {
			if marshaler, ok := v.(easyjson.Marshaler); ok {
				return easyjson.Marshal(marshaler)
			}
			return json.Marshal(v)
		},
		JSONDecoder: func(data []byte, v interface{}) error {
			if unmarshaler, ok := v.(easyjson.Unmarshaler); ok {
				return easyjson.Unmarshal(data, unmarshaler)
			}
			return json.Unmarshal(data, v)
		},
	})

	app.Get("/contacts", getContacts)
	app.Get("/contacts/:id", getContact)
	app.Post("/contacts", addContacts)
	app.Put("/contacts/:id", updateContact)
	app.Delete("/contacts/:id", deleteContact)

	log.Fatal(app.Listen(":3000"))
}

func getContact(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect ID"})
	}

	var contact model.Contact
	query := "SELECT id, name, phone FROM contacts WHERE id = $1"
	err = db.QueryRow(query, id).Scan(&contact.ID, &contact.Name, &contact.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Couldn't find contact"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}

	return c.JSON(&contact)
}

func getContacts(c fiber.Ctx) error {
	rows, err := db.Query("SELECT id, name, phone FROM contacts ORDER BY id ASC")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Request error"})
	}
	defer rows.Close()

	var contacts []model.Contact

	for rows.Next() {
		var contact model.Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Data read error"})
		}
		contacts = append(contacts, contact)
	}

	if err = rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Rows error"})
	}

	if len(contacts) == 0 {
		return c.JSON([]model.Contact{})
	}
	return c.JSON(contacts)
}

func addContacts(c fiber.Ctx) error {
	var contact model.Contact
	if err := easyjson.Unmarshal(c.Body(), &contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect JSON"})
	}

	query := "INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id"
	err := db.QueryRow(query, contact.Name, contact.Phone).Scan(&contact.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Creating failed"})
	}

	return c.Status(fiber.StatusCreated).JSON(&contact)
}

func updateContact(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect ID"})
	}

	var updatedContact model.Contact
	if err := easyjson.Unmarshal(c.Body(), &updatedContact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect JSON"})
	}

	query := "UPDATE contacts SET name = $1, phone = $2 WHERE id = $3"
	result, err := db.Exec(query, updatedContact.Name, updatedContact.Phone, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Update error"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contact not found"})
	}

	updatedContact.ID = id
	return c.JSON(updatedContact)

}

func deleteContact(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect ID"})
	}

	query := "DELETE FROM contacts WHERE id = $1"
	result, err := db.Exec(query, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Deleting error"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contact not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
