package adapters

import (
	model "api/models/users/store"
	"encoding/json"
	"net/http"
)

func Store(r *http.Request) model.User {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	user.SetCreated()
	user.SetUpdated()
	return user
}
