package repository

import (
	"context"

	"github.com/Arsaev/WalkMeThrough/backend/internal/entity"
)

// UserGroupRepository - interface defines methods to manage user and group relations on persistence layer
type UserGroupRepository interface {
	// Create - saves the user group struct in the database
	Create(ctx context.Context, userGroup *entity.UserGroup) error

	// GetByUserAndGroup attempt to fetch a user group by user and group IDs
	// if there is no user group with the given user and group IDs, it should return nil and nil
	GetByUserAndGroup(ctx context.Context, userID, groupID string) (*entity.UserGroup, error)
	// ListByUserID - fetches all the groups that a user is assigned to
	ListByUserID(ctx context.Context, userID string, limit, offset int64) ([]*entity.Group, error)
	// ListByGroupID - fetches all user groups with given group ID
	ListByGroupID(ctx context.Context, groupID string, limit, offset int64) ([]*entity.User, error)
	// Delete - deletes the user group struct from the database for the given ID
	// (unique id of user group, not user or group IDs)
	Delete(ctx context.Context, id string) error
}
