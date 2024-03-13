package usecase

import (
	"errors"
	"fmt"
	"user/service/pkg/domain"
	"user/service/pkg/helper"
	interfaces "user/service/pkg/repository/interface"
	services "user/service/pkg/usecase/interface"
	"user/service/pkg/utils/models"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepository interfaces.UserRepository
}

func NewUserUseCase(repository interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepository: repository,
	}
}

func (ur *userUseCase) UsersSignUp(user models.UserSignUp) (domain.TokenUser, error) {
	email, err := ur.userRepository.CheckUserExistsByEmail(user.Email)
	fmt.Println(email)
	if err != nil {
		return domain.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return domain.TokenUser{}, errors.New("user with this email is already exists")
	}

	phone, err := ur.userRepository.CheckUserExistsByPhone(user.Phone)
	fmt.Println(phone, nil)
	if err != nil {
		return domain.TokenUser{}, errors.New("error with server")
	}
	if phone != nil {
		return domain.TokenUser{}, errors.New("user with this phone is already exists")
	}

	hashPassword, err := helper.PasswordHash(user.Password)
	if err != nil {
		return domain.TokenUser{}, errors.New("error in hashing password")
	}
	user.Password = hashPassword
	userData, err := ur.userRepository.UserSignUp(user)
	if err != nil {
		return domain.TokenUser{}, errors.New("could not add the user")
	}
	accessToken, err := helper.GenerateAccessToken(userData)
	if err != nil {
		return domain.TokenUser{}, errors.New("couldn't create access token due to error")
	}
	refreshToken, err := helper.GenerateRefreshToken(userData)
	if err != nil {
		return domain.TokenUser{}, errors.New("couldn't create refresh token due to error")
	}
	return domain.TokenUser{
		User:         userData,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (ur *userUseCase) UsersLogin(user models.UserLogin) (domain.TokenUser, error) {
	email, err := ur.userRepository.CheckUserExistsByEmail(user.Email)
	if err != nil {
		return domain.TokenUser{}, errors.New("error with server")
	}
	if email == nil {
		return domain.TokenUser{}, errors.New("email doesn't exist")
	}
	userdeatils, err := ur.userRepository.FindUserByEmail(user)
	if err != nil {
		return domain.TokenUser{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userdeatils.Password), []byte(user.Password))
	if err != nil {
		return domain.TokenUser{}, errors.New("password not matching")
	}
	var userDetails models.UserDetails
	err = copier.Copy(&userDetails, &userdeatils)
	if err != nil {
		return domain.TokenUser{}, err
	}
	accessToken, err := helper.GenerateAccessToken(userDetails)
	if err != nil {
		return domain.TokenUser{}, errors.New("couldn't create accesstoken due to internal error")
	}
	refreshToken, err := helper.GenerateRefreshToken(userDetails)
	if err != nil {
		return domain.TokenUser{}, errors.New("counldn't create refreshtoken due to internal error")
	}
	return domain.TokenUser{
		User:         userDetails,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
