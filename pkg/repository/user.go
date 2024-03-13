package repository

import (
	"errors"
	"user/service/pkg/domain"
	interfaces "user/service/pkg/repository/interface"
	"user/service/pkg/utils/models"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (ur *userRepository) CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := ur.DB.Where(&domain.User{Email: email}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.User{}, res.Error
	}
	return &user, nil
}
func (ur *userRepository) CheckUserExistsByPhone(phone string) (*domain.User, error) {
	var user domain.User
	res := ur.DB.Where(&domain.User{Phone: phone}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.User{}, res.Error
	}
	return &user, nil
}
func (ur *userRepository) UserSignUp(user models.UserSignUp) (models.UserDetails, error) {
	var signupDetail models.UserDetails
	err := ur.DB.Raw(`
		INSERT INTO users(firstname, lastname, email, password, phone)
		VALUES(?, ?, ?, ?, ?)
		RETURNING id, firstname, lastname, email, phone
	`, user.Firstname, user.Lastname, user.Email, user.Password, user.Phone).
		Scan(&signupDetail).Error

	if err != nil {
		return models.UserDetails{}, err
	}
	return signupDetail, nil
}
func (ur *userRepository) FindUserByEmail(user models.UserLogin) (models.UserDetail, error) {
	var userDetails models.UserDetail
	err := ur.DB.Raw("SELECT * FROM users WHERE email=?", user.Email).Scan(&userDetails).Error
	if err != nil {
		return models.UserDetail{}, errors.New("error checking user details")
	}
	return userDetails, nil
}
