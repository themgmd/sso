package types

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id         uuid.UUID   `json:"id"`
	GivenName  string      `json:"given_name"`
	FamilyName string      `json:"family_name"`
	Patronymic null.String `json:"patronymic"`
	Email      string      `json:"email"`
	Password   string      `json:"-"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  null.Time   `json:"deleted_at"`
}

func (u User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
