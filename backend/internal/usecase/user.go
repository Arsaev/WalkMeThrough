package usecase

import (
	log "github.com/sirupsen/logrus"

	"github.com/Arsaev/WalkMeThrough/backend/internal/repository"
)

// UserUsecase implements the business logic for the user entity
// like creating a user, updating a user, deleting a user, authenticating a user, etc.
type UserUseCase struct {
	userStore repository.UserRepository

	// permissionUC - use case for permission entity
	permissionUC *PermissionsUsecase

	logger *log.Logger
}

// NewUserUseCase creates a new user use case
func NewUserUseCase(userStore repository.UserRepository, permissionUC *PermissionsUsecase, logger *log.Logger) *UserUseCase {
	if logger == nil {
		logger = log.New()
	}

	if userStore == nil {
		logger.Fatal("userStore is nil")
	}

	if permissionUC == nil {
		logger.Fatal("permissionUC is nil")
	}

	return &UserUseCase{
		userStore:    userStore,
		permissionUC: permissionUC,
		logger:       logger,
	}
}
