// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"time"
)
// Representação de tabela
type Product struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Price     int32     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
