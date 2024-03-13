package interfaces

import (
	"user/service/pkg/domain"
	"user/service/pkg/utils/models"
)

type UserUseCase interface {
	UsersSignUp(user models.UserSignUp) (domain.TokenUser, error)
	UsersLogin(user models.UserLogin) (domain.TokenUser, error)
}
