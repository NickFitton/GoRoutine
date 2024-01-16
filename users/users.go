package users

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type User struct {
	id          uuid.UUID
	email       string `binding:"required"`
	createdAt   time.Time
	lastLoginAt time.Time
}

func New(id uuid.UUID,
	email string,
	createdAt time.Time,
	lastLoginAt time.Time) User {
	newUser := User{
		id:          id,
		email:       email,
		createdAt:   createdAt,
		lastLoginAt: lastLoginAt,
	}
	return newUser
}

var users = []User{}

func GetAll() []User {
	return users
}

func Create(user User) User {
	newUser := User{
		email:       user.email,
		id:          uuid.New(),
		createdAt:   time.Now(),
		lastLoginAt: time.Now(),
	}
	users = append(users, newUser)
	log.Println(newUser)
	return newUser
}

func Delete(emailAddress string) bool {
	for i, user := range users {
		if user.email == emailAddress {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
