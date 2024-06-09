package user

import "github.com/google/uuid"

type User struct {
	id        uuid.UUID
	username  string
	email     string
	password  string
	isDeleted bool
}

func New(username string, email string, password string) *User {
	return &User{id: uuid.New(), username: username, email: email, password: password, isDeleted: false}
}

func Assemble(id uuid.UUID, username string, email string, password string, isDeleted bool) *User {
	return &User{id: id, username: username, email: email, password: password, isDeleted: isDeleted}
}

func (u *User) Id() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) IsDeleted() bool {
	return u.isDeleted
}
