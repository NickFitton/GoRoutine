package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	Email       string    `binding:"required" json:"email"`
	CreatedAt   time.Time `json:"createdAt"`
	LastLoginAt time.Time `json:"lastLoginAt"`
}

var users = []User{}

func GetAll() []User {
	return users
}

func Create(user User) User {
	user.Id = uuid.New()
	user.CreatedAt = time.Now()
	user.LastLoginAt = time.Now()
	users = append(users, user)
	return user
}

func FindById(id uuid.UUID) (User, bool) {
	for _, user := range users {
		if user.Id == id {
			return user, true
		}
	}
	return User{}, false
}

func Delete(userId uuid.UUID) bool {
	for i, user := range users {
		if user.Id == userId {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
