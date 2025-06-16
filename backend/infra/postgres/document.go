package postgres


import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"errors"
	"context"
	"github.com/Arsaev/WalkMeThrough/backend/internal/entity"
)

// Documents implement repository/documents.go interface for postgresql
type Documents struct {
	conn *sql.DB
	logger *logrus.Logger
}


// Create function implements Documents interfaces by splitting object condition rules
// from object and storing them separated in different table for optimal search
func (d *Documents) Create(ctx context.Context, in *entity.Document) error {
	logentry := d.logger.WithContext(ctx).WithFields(logrus.Fields{
		"func": "Create",
		"package": "/infra/postgres/document.go",
	})

	if in == nil {
		logentry.Error("input is nil")
		return errors.New("input is nil")
	}

	tx, err := d.conn.Begin()
	if err != nil {
		logentry.WithError(err).Error("failed to start transaction")
		return errors.New("failed to create record due to faild transaction begin")
	}

	defer tx.Rollback()
	// save Document
	//
	// save all groups
	//
	// save all rules in each group
	//


	tx.Commit()
	return nil
}
