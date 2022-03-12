package adapters

import (
	model "api/models/users/update"
	"encoding/json"
	"net/http"
)

func Update(r *http.Request, id int) model.User {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	user.SetUpdated()
	user.Id = id
	return user
}
