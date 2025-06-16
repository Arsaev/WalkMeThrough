package usecase

import (
	"fmt"
)


// Documents usecase implement business logic for managing stored documents
type Documents struct {
	logger *logrus.Logger
	// store access
}


// Create -
func (d *Documents) Create(ctx context.Context, in *entity.Document) (*entity.Document, error) {
	return nil, errors.New("not implemented")
}

// Update
func (d *Documents) Update(ctx context.Context, in *entity.Document) (*entity.Document, error) {
	return nil, errors.New("not implemented")
}


// Search returns list of Documents matching given user data
func (d *Documents) Search(ctx context.Context, in *entity.UserData) ([]*entity.Document, error) {
	//validate data
	//
	// pass to db adapter to
	//
	// return list
	//
	// TODO add pagination
}
