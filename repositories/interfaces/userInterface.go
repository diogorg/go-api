package interfaces

import (
	modelShow "api/models/users/show"
	modelStore "api/models/users/store"
	modelUpdate "api/models/users/update"
)

type UserRepository interface {
	GetAll() []modelShow.User
	FindById(id int) modelShow.User
	Store(user *modelStore.User)
	Delete(id int)
	Update(user *modelUpdate.User)
}
