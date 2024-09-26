package repositories

import (
	"context"
	"database/sql"

	"github.com/Rakhulsr/go-rest-api-pzn/model/domain"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category) error
}
