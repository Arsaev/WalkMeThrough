package repository

import (
	"context"
	"github.com/Arsaev/WalkMeThrough/backend/internal/entity"
)

type Documents interface {
	Create(ctx context.Context, in *entity.Document) (*entity.Document, error)
	Update(ctx context.Context, in *entity.Document) (*entity.Document, error)
	Search(ctx context.Context, in *entity.UserData) ([]*entity.Document, error)
}
