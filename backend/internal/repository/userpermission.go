package repository

import (
	"context"

	"github.com/Arsaev/WalkMeThrough/backend/internal/entity"
)

// UserPermissionRepository - interface defines methods to manage atomic user permissions on persistence layer
type UserPermissionRepository interface {
	// Create - saves the user permission struct in the database
	Create(ctx context.Context, userPermission *entity.UserPermission) error
	// ListByUserID - fetches all the permissions that a user is assigned to
	ListByUserID(ctx context.Context, userID string, limit, offset int64) ([]*entity.Permission, error)
	// GetByUserAndPermission - fetches a user permission by user and permission ID
	GetByUserAndPermission(ctx context.Context, userID, permissionID string) (*entity.UserPermission, error)
	// Delete - deletes the user permission struct from the database for the given ID
	// (unique id of user permission record, not user or permission IDs)
	Delete(ctx context.Context, id string) error
}
