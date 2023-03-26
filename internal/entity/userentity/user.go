package userentity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID        string    `gorm:"size:36;not null;unique index;primaryKey"`
	Name      string    `gorm:"size:255;"`
	Email     string    `gorm:"size:100;unique"`
	Username  string    `gorm:"size:100;unique"`
	Password  string    `gorm:"size:100;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// Users represent body for get data from user
type UserData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// UserRequest is payload for register user
type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type FindUserRequest struct {
	ID string `json:"id"`
}

// UserRequestValidate is to validate input request
func UserRequestValidate(ur *UserRequest) error {
	err := validation.Errors{
		"name":     validation.Validate(&ur.Name, validation.Required, validation.Length(2, 40)),
		"email":    validation.Validate(&ur.Email, validation.Required),
		"username": validation.Validate(&ur.Username, validation.Required),
		"password": validation.Validate(&ur.Password, validation.Required),
	}

	return err.Filter()
}
