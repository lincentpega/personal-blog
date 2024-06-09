package ports

import (
	"context"
	"github.com/google/uuid"
	"identity/user"
)

type UserRepository interface {
	AddUser(context.Context, *user.User) error
	UpdateUser(context.Context, *user.User) error
	GetUserById(context.Context, uuid.UUID) (*user.User, error)
	DeleteUser(context.Context, uuid.UUID) error
}
