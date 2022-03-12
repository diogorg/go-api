package show

import (
	"database/sql"
)

type User struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Login     string         `json:"login"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
	DeletedAt sql.NullString `json:"deleted_at"`
}
