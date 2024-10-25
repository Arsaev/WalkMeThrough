package repository

import (
	"context"

	"github.com/Arsaev/WalkMeThrough/backend/internal/entity"
)

type PasswordRepository interface {
	// Create - saves the password struct in the database
	Create(ctx context.Context, password *entity.Password) error
	// GetByUserID - fetches the password for a user based on the user ID
	GetByUserID(ctx context.Context, userID string) (*entity.Password, error)
}
