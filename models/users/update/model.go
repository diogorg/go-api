package update

import (
	"api/support"
	"database/sql"
)

type User struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Login     string         `json:"login"`
	Password  string         `json:"password"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

func (u *User) SetUpdated() {
	u.UpdatedAt = sql.NullString{String: support.CurrentDateTime(), Valid: true}
}
