package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Arsaev/WalkMeThrough/backend/internal/entity"
	"github.com/Arsaev/WalkMeThrough/backend/internal/repository"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

// PermissionsUsecase - struct implements user permissions logic
// like CRUD on user groups and there permissions, single permissions for users etc.
type PermissionsUsecase struct {
	groupstore repository.GroupRepository
	permstore  repository.PermissionRepository
	usergroup  repository.UserGroupRepository
	userperm   repository.UserPermissionRepository
	userstore  repository.UserRepository

	logger *log.Logger
}

// PermissionsUCConfig
type PermissionsUCConfig struct {
	GroupStore repository.GroupRepository
	PermStore  repository.PermissionRepository
	UserGroup  repository.UserGroupRepository
	UserPerm   repository.UserPermissionRepository
	UserStore  repository.UserRepository

	Logger *log.Logger
}

// Validate - validates the input fields of the struct
func (cfg PermissionsUCConfig) Validate() error {
	// check if store interfaces are not nil
	if cfg.GroupStore == nil {
		return fmt.Errorf("group store is nil")
	}
	if cfg.PermStore == nil {
		return fmt.Errorf("permission store is nil")
	}
	if cfg.UserGroup == nil {
		return fmt.Errorf("user group store is nil")
	}
	if cfg.UserPerm == nil {
		return fmt.Errorf("user permission store is nil")
	}

	if cfg.Logger == nil {
		cfg.Logger = log.New()
	}

	return nil
}

// NewPermissionsUsecase - returns a new instance of PermissionsUsecase
func NewPermissionsUsecase(cfg PermissionsUCConfig) *PermissionsUsecase {
	// validate the input fields
	if err := cfg.Validate(); err != nil {
		log.Fatalf("PermissionsUsecase config validation failed: %v", err)
	}

	return &PermissionsUsecase{
		groupstore: cfg.GroupStore,
		permstore:  cfg.PermStore,
		usergroup:  cfg.UserGroup,
		userperm:   cfg.UserPerm,
	}
}

// CreateGroup - creates a new group
func (uc *PermissionsUsecase) CreateGroup(ctx context.Context, name string, permissions []*entity.Permission) (*entity.Group, error) {
	logentry := log.WithContext(ctx).WithFields(log.Fields{
		"func":    "CreateGroup",
		"Package": "usecase.permissions",
	})

	if name == "" {
		logentry.Error("group name is empty")
		return nil, fmt.Errorf("group name is empty")
	}

	// create a new group
	group := &entity.Group{
		ID:          uuid.New().String(),
		Name:        name,
		Permissions: permissions,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	// save the group in the database
	if err := uc.groupstore.Create(ctx, group); err != nil {
		logentry.WithError(err).Error("failed to save group in the database")
		return nil, err
	}

	return group, nil
}

// GetGroup - returns a group by its id
func (uc *PermissionsUsecase) GetGroup(ctx context.Context, id string) (*entity.Group, error) {
	logentry := log.WithContext(ctx).WithFields(log.Fields{
		"func":    "GetGroup",
		"Package": "usecase.permissions",
	})

	if id == "" {
		logentry.Error("group id is empty")
		return nil, fmt.Errorf("group id is empty")
	}

	// get the group from the database
	group, err := uc.groupstore.Get(ctx, id)
	if err != nil {
		logentry.WithError(err).Error("failed to get group from the database")
		return nil, err
	}

	return group, nil
}

// UpdateGroup - updates a group by its id
func (uc *PermissionsUsecase) Update(ctx context.Context, group *entity.Group) (*entity.Group, error) {
	logentry := log.WithContext(ctx).WithFields(log.Fields{
		"func":    "UpdateGroup",
		"Package": "usecase.permissions",
	})

	if group == nil {
		logentry.Error("group is nil")
		return nil, fmt.Errorf("group is nil")
	}

	if group.ID == "" {
		logentry.Error("group id is empty")
		return nil, fmt.Errorf("group id is empty")
	}

	// fetche the group from the database
	oldGroup, err := uc.groupstore.Get(ctx, group.ID)
	if err != nil {
		logentry.WithError(err).Error("failed to get group from the database")
		return nil, err
	}

	group.CreatedAt = oldGroup.CreatedAt
	group.UpdatedAt = time.Now().Unix()

	// update the group in the database
	if err := uc.groupstore.Update(ctx, group); err != nil {
		logentry.WithError(err).Error("failed to update group in the database")
		return nil, err
	}

	return group, nil
}

// AssignUserToGroup - creates user group relation record
func (uc *PermissionsUsecase) AssignUserToGroup(ctx context.Context, userID, groupID string) (*entity.UserGroup, error) {
	logentry := log.WithContext(ctx).WithFields(log.Fields{
		"func":    "AssignUserToGroup",
		"Package": "usecase.permissions",
	})

	if userID == "" {
		logentry.Error("user id is empty")
		return nil, fmt.Errorf("user id is empty")
	}

	if groupID == "" {
		logentry.Error("group id is empty")
		return nil, fmt.Errorf("group id is empty")
	}

	// check if user group relation already exists
	fetchedUserGroup, err := uc.usergroup.GetByUserAndGroup(ctx, userID, groupID)
	if err == nil {
		logentry.Error("failed to check if user already belongs to the group")
		return nil, fmt.Errorf("failed to check if user already belongs to the group")
	}

	if fetchedUserGroup != nil {
		logentry.Error("user already belongs to the group")
		return fetchedUserGroup, nil // this makes the function idempotent
	}

	// check if user exists
	_, err := uc.userstore.Get(ctx, userID)
	if user == nil {
		// todo: check if the error is user not found error or some other error
		logentry.Error("failed to check if user exists")
		return nil, fmt.Errorf("failed to check if user exists")
	}

	// check if group exists
	_, err := uc.groupstore.Get(ctx, groupID)
	if group == nil {
		logentry.Error("failed to check if group exists")
		return nil, fmt.Errorf("failed to check if group exists")
	}

	// create a new user group relation
	usergroup := &entity.UserGroup{
		ID:        uuid.New().String(),
		UserID:    userID,
		GroupID:   groupID,
		CreatedAt: time.Now().Unix(),
	}

	err = uc.usergroup.Create(ctx, usergroup)
	if err != nil {
		logentry.WithError(err).Error("failed to create user group relation")
		return nil, fmt.Errorf("failed to create user group relation, %w" err)
	}

	return usergroup, nil
}

// RemoveUserFromGroup - removes user group relation record
func (uc *PermissionsUsecase) RemoveUserFromGroup(ctx context.Context, userID, groupID string) error {
	logentry := log.WithContext(ctx).WithFields(log.Fields{
		"func":    "RemoveUserFromGroup",
		"Package": "usecase.permissions",
	})

	if userID == "" {
		logentry.Error("user id is empty")
		return fmt.Errorf("user id is empty")
	}

	if groupID == "" {
		logentry.Error("group id is empty")
		return fmt.Errorf("group id is empty")
	}

	// check if user group relation exists
	usergroup, err := uc.usergroup.GetByUserAndGroup(ctx, userID, groupID)
	if err != nil {
		logentry.WithError(err).Error("failed to check if user group relation exists")
		return fmt.Errorf("failed to check if user group relation exists")
	}

	if usergroup == nil {
		logentry.Error("user group relation does not exist")
		return nil // this makes the function idempotent
	}

	// delete the user group relation
	err = uc.usergroup.Delete(ctx, usergroup.ID)
	if err != nil {
		logentry.WithError(err).Error("failed to delete user group relation")
		return fmt.Errorf("failed to delete user group relation")
	}

	return nil
}

// ListUserGroups - returns all the groups that a user is assigned to
func (uc *PermissionsUsecase) ListUserGroups(ctx context.Context, userID string, limit, offset int64) ([]*entity.Group, error) {
	logentry := log.WithContext(ctx).WithFields(log.Fields{
		"func":    "ListUserGroups",
		"Package": "usecase.permissions",
	})

	if userID == "" {
		logentry.Error("user id is empty")
		return nil, fmt.Errorf("user id is empty")
	}

	// fetch all the groups that a user is assigned to
	groups, err := uc.usergroup.ListByUserID(ctx, userID, limit, offset)
	if err != nil {
		logentry.WithError(err).Error("failed to fetch user groups")
		return nil, fmt.Errorf("failed to fetch user groups")
	}

	return groups, nil
}

// ListUserPermissions - returns all the permissions that a user has
// todo pagination 
func (uc *PermissionsUsecase) ListUserPermissions(ctx context.Context, userID string, limit, offset int64) ([]*entity.Permission, error) {
	logentry := log.WithContext(ctx).WithFields(log.Fields{
		"func":    "ListUserPermissions",
		"Package": "usecase.permissions",
	})

	if userID == "" {
		logentry.Error("user id is empty")
		return nil, fmt.Errorf("user id is empty")
	}

	// fetch all the permissions that a user has
	permissions, err := uc.userperm.ListByUserID(ctx, userID, limit, offset)
	if err != nil {
		logentry.WithError(err).Error("failed to fetch user permissions")
		return nil, fmt.Errorf("failed to fetch user permissions")
	}

	return permissions, nil
}

// GetUserPermissions - returns all the permissions that a user has combined from groups and user permissions 
func (uc *PermissionsUsecase) GetUserPermissions(ctx context.Context, userID string) ([]*entity.Permission, error) {
	logentry := log.WithContext(ctx).WithFields(log.Fields{
		"func":    "GetUserPermissions",
		"Package": "usecase.permissions",
	})

	if userID == "" {
		logentry.Error("user id is empty")
		return nil, fmt.Errorf("user id is empty")
	}

	// fetch all the permissions that a user has
	permissions, err := uc.userperm.ListByUserID(ctx, userID, limit, offset)
	if err != nil {
		logentry.WithError(err).Error("failed to fetch user permissions")
		return nil, fmt.Errorf("failed to fetch user permissions")
	}

	// fetch all the groups that a user is assigned to
	groups, err := uc.usergroup.ListByUserID(ctx, userID, limit, offset)
	if err != nil {
		logentry.WithError(err).Error("failed to fetch user groups")
		return nil, fmt.Errorf("failed to fetch user groups")
	}

	// make combined list of permissions
	for _, group := range groups {
		permissions = append(permissions, group.Permissions...)
	}

	return permissions, nil
}


// CanUserPerformAction - checks if a user has a permission to perform an action. Action or id should be defined
// just for demo purposes, in real world, RBAC should be implemented using tool like casbin
func (uc *PermissionsUsecase) CanUserPerformAction(ctx context.Context, userID, action, id string) (bool, error) {
	logentry := log.WithContext(ctx).WithFields(log.Fields{
		"func":    "CanUserPerformAction",
		"Package": "usecase.permissions",
	})

	if userID == "" {
		logentry.Error("user id is empty")
		return false, fmt.Errorf("user id is empty")
	}

	if action == "" && id == "" {
		logentry.Error("action or id should be defined")
		return false, fmt.Errorf("action or id should be defined")
	}

	// fetch all the permissions that a user has
	permissions, err := uc.GetUserPermissions(ctx, userID)
	if err != nil {
		logentry.WithError(err).Error("failed to fetch user permissions")
		return false, fmt.Errorf("failed to fetch user permissions")
	}

	// check if the user has the permission
	for _, permission := range permissions {
		if permission.Name == action || permission.ID == id {
			return true, nil
		}
	}

	return false, nil
}


// AssignUserPermission - assigns a permission to a user
func (uc *PermissionsUsecase) AssignUserPermission(ctx context.Context, userID, permissionID string) (*entity.UserPermission, error) {
	logentry := log.WithContext(ctx).WithFields(log.Fields{
		"func":    "AssignUserPermission",
		"Package": "usecase.permissions",
	})

	if userID == "" {
		logentry.Error("user id is empty")
		return nil, fmt.Errorf("user id is empty")
	}

	if permissionID == "" {
		logentry.Error("permission id is empty")
		return nil, fmt.Errorf("permission id is empty")
	}

	// check if user permission relation already exists
	fetchedUserPermission, err := uc.userperm.GetByUserAndPermission(ctx, userID, permissionID)
	if err == nil {
		logentry.Error("failed to check if user already has the permission")
		return nil, fmt.Errorf("failed to check if user already has the permission")
	}

	if fetchedUserPermission != nil {
		logentry.Error("user already has the permission")
		return fetchedUserPermission, nil // this makes the function idempotent
	}

	// check if user exists
	_, err := uc.userstore.Get(ctx, userID)
	if err != nil {
		logentry.Error("failed to check if user exists")
		return nil, fmt.Errorf("failed to check if user exists")
	}

	// permission exists since they are hardcoded


	// create a new user permission relation
	userpermission := &entity.UserPermission{
		ID:           uuid.New().String(),
		UserID:       userID,
		PermissionID: permissionID,
		CreatedAt:    time.Now().Unix(),
	}

	err = uc.userperm.Create(ctx, userpermission)
	if err != nil {
		logentry.WithError(err).Error("failed to create user permission relation")
		return nil, fmt.Errorf("failed to create user permission relation, %w" err)
	}

	return userpermission, nil
}


// RemoveUserPermission - removes user permission relation record
func (uc *PermissionsUsecase) RemoveUserPermission(ctx context.Context, userID, permissionID string) error {
	logentry := log.WithContext(ctx).WithFields(log.Fields{
		"func":    "RemoveUserPermission",
		"Package": "usecase.permissions",
	})

	if userID == "" {
		logentry.Error("user id is empty")
		return fmt.Errorf("user id is empty")
	}

	if permissionID == "" {
		logentry.Error("permission id is empty")
		return fmt.Errorf("permission id is empty")
	}

	// check if user permission relation exists
	userpermission, err := uc.userperm.GetByUserAndPermission(ctx, userID, permissionID)
	if err != nil {
		logentry.WithError(err).Error("failed to check if user permission relation exists")
		return fmt.Errorf("failed to check if user permission relation exists")
	}

	if userpermission == nil {
		logentry.Error("user permission relation does not exist")
		return nil // this makes the function idempotent
	}

	// delete the user permission relation
	err = uc.userperm.Delete(ctx, userpermission.ID)
	if err != nil {
		logentry.WithError(err).Error("failed to delete user permission relation")
		return fmt.Errorf("failed to delete user permission relation")
	}

	return nil
}