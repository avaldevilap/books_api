package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

func UnmarshalBook(data []byte) (Book, error) {
	var r Book
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Book) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Book struct to describe book object.
type Book struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	UserID     uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Title      string    `db:"title" json:"title" validate:"required,lte=255"`
	Author     string    `db:"author" json:"author" validate:"required,lte=255"`
	BookStatus int       `db:"book_status" json:"book_status" validate:"required,len=1"`
	BookAttrs  BookAttrs `db:"book_attrs" json:"book_attrs" validate:"required,dive"`
}

// BookAttrs struct to describe book attributes.
type BookAttrs struct {
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
}
