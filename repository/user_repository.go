package repository

import (
	"errors"
	"fmt"
	"go-postgres-test-1/model"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(username, password, email string) (err model.Error)
	GetAllUsers() (users []model.User, err model.Error)
	GetUser(id uint) (user model.User, err model.Error)
	UpdateUser(id uint, newEmail, newPassword string) (err model.Error)
	DeleteUser(id uint) (err model.Error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(username, password, email string) (err model.Error) {
	user := model.User{Username: username, Password: password, Email: email}
	result := r.db.Create(&user)
	if result.Error != nil {
		fmt.Println("Error creating user:", result.Error.Error())
		if isDuplicateKeyError(result.Error) {
			return model.Error{StatusCode: http.StatusConflict, Message: "Username already taken"}
		} else {
			return model.Error{StatusCode: http.StatusInternalServerError, Message: "Error creating user"}
		}
	}
	return err
}

func (r *userRepository) GetAllUsers() (users []model.User, err model.Error) {
	result := r.db.Find(&users)
	if result.Error != nil {
		fmt.Println("Error fetching users:", result.Error.Error())
		return []model.User{}, model.Error{StatusCode: http.StatusInternalServerError, Message: "Error fetching users"}
	}
	return users, err
}

func (r *userRepository) GetUser(id uint) (user model.User, err model.Error) {
	result := r.db.First(&user, id)
	if result.Error != nil {
		fmt.Println("Error fetching user:", result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, model.Error{StatusCode: http.StatusNotFound, Message: "User not found"}
		} else {
			return model.User{}, model.Error{StatusCode: http.StatusInternalServerError, Message: "Error fetching user"}
		}
	}
	return user, err
}

func (r *userRepository) UpdateUser(id uint, newEmail, newPassword string) (err model.Error) {
	user, err := r.GetUser(id)
	if err.StatusCode != 0 {
		return err
	}

	if newEmail != "" {
		user.Email = newEmail
	}
	if newPassword != "" {
		user.Password = newPassword
	}

	result := r.db.Save(&user)
	if result.Error != nil {
		fmt.Println("Error updating user:", result.Error.Error())
		return model.Error{StatusCode: http.StatusInternalServerError, Message: "Error updating user"}
	}
	return err
}

func (r *userRepository) DeleteUser(id uint) (err model.Error) {
	_, err = r.GetUser(id)
	if err.StatusCode != 0 {
		return err
	}

	result := r.db.Delete(&model.User{}, id)
	if result.Error != nil {
		fmt.Println("Error deleting user:", result.Error.Error())
		return model.Error{StatusCode: http.StatusInternalServerError, Message: "Error deleting user"}

	}
	return err
}

func isDuplicateKeyError(err error) bool {
	if err, ok := err.(*pgconn.PgError); ok {
		return err.Code == "23505" // Duplicate key violation
	}
	return false
}
