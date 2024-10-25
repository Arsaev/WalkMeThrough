package repository

import (
	"context"

	"github.com/Arsaev/WalkMeThrough/backend/internal/entity"
)

// GroupRepository - interface defines methods to manage groups on persistence layer
type GroupRepository interface {
	// Create - saves the group struct in the database
	Create(ctx context.Context, group *entity.Group) error
	// Get - fetches the group struct from the database for the given ID
	Get(ctx context.Context, id string) (*entity.Group, error)
	// List - fetches all the groups from the database
	List(ctx context.Context, limit, offset int64) ([]*entity.Group, error)
	// Update - updates the group struct in the database
	Update(ctx context.Context, group *entity.Group) error
	// Delete - deletes the group struct from the database for the given ID
	Delete(ctx context.Context, id string) error
}
