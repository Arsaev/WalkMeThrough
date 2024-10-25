package repository

import (
	"context"

	"github.com/Arsaev/WalkMeThrough/backend/internal/entity"
)

// UserRepository - interface defines methods to work with user on persistence layer
type UserRepository interface {
	// Upsert - creates or updates a user in the database based on the user ID
	Upsert(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	List(ctx context.Context, limit, offset int64) ([]*entity.User, error)
}
