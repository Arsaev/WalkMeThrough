package entity

// Group represents a permission level in the system
// Group contains a list of permissions that are granted to users in the group
type Group struct {
	ID          string        `json:"id"` // unique identifier for the group
	Name        string        `json:"name"`
	Permissions []*Permission `json:"permissions"` // list of permissions that the group has

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

// UserGroup is a group that user are assigned to.
// This will allow to fetch all the groups that a user is assigned to and all the users that are assigned to a group
type UserGroup struct {
	ID      string `json:"id"` // unique identifier for the user group
	UserID  string `json:"user_id"`
	GroupID string `json:"group_id"`

	CreatedAt int64 `json:"created_at"`
}

// UserPermuission is a permission that user are assigned to.
// These are specific permissions that are assigned to a user beyond the group permissions
type UserPermission struct {
	ID           string `json:"id"` // unique identifier for the user permission
	UserID       string `json:"user_id"`
	PermissionID string `json:"permission_id"`

	CreatedAt int64 `json:"created_at"`
}

// Permission represent atomic action that can be performed in the system
type Permission struct {
	ID   string `json:"id"` // unique identifier for the permission
	Name string `json:"name"`
}
