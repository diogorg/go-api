package repositories

import (
	db "api/db"
	modelShow "api/models/users/show"
	modelStore "api/models/users/store"
	modelUpdate "api/models/users/update"
)

type UserPotgresRepository struct {
}

func (c UserPotgresRepository) GetAll() []modelShow.User {
	users := []modelShow.User{}
	db.Find(&users)
	return users
}

func (c UserPotgresRepository) FindById(id int) modelShow.User {
	var user modelShow.User
	db.First(&user, id)
	return user
}

func (c UserPotgresRepository) Store(user *modelStore.User) {
	db.Create(user)
}

func (c UserPotgresRepository) Delete(id int) {
	var user modelShow.User
	db.Delete(&user, id)
}

func (c UserPotgresRepository) Update(user *modelUpdate.User) {
	db.Save(user)
}
