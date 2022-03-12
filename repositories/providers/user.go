package providers

import (
	interfaces "api/repositories/interfaces"
	user "api/repositories/users"
)

func GetUserRespository() interfaces.UserRepository {
	repository := user.UserPotgresRepository{}

	return repository
}
