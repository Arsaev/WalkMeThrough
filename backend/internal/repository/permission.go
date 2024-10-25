package repository

import "github.com/Arsaev/WalkMeThrough/backend/internal/entity"

// Permission will be hardcoded in and returnd by struct that implements this interface.
// Its defined this way since single actions are not dynamic but predefined by this codebase.
type PermissionRepository interface {
	// Get - fetches the permission struct from the database for the given ID
	Get(id string) (*entity.Permission, error)
	// List - fetches all the permissions from the database
	// List of permissions wont be huge thus no pagination is needed
	List() ([]*entity.Permission, error)
}
