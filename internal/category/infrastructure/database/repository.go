package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gadoma/rafapi/internal/category/domain"
	"github.com/gadoma/rafapi/internal/common/infrastructure/database"
	"github.com/oklog/ulid/v2"
)

var _ domain.CategoryRepository = (*CategoryRepository)(nil)

type CategoryRepository struct {
	db *database.DB
}

func NewCategoryRepository(db *database.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetCategories(ctx context.Context) ([]*domain.Category, int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()

	return getCategories(ctx, tx)
}

func (r *CategoryRepository) GetCategory(ctx context.Context, id ulid.ULID) (*domain.Category, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	return getCategory(ctx, tx, id)
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, ccc *domain.CreateCategoryCommand) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = createCategory(ctx, tx, ccc.Id, ccc.Name)

	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, id ulid.ULID, ucc *domain.UpdateCategoryCommand) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := updateCategory(ctx, tx, id, ucc.Name); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id ulid.ULID) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteCategory(ctx, tx, id); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func getCategories(ctx context.Context, tx *database.Tx) ([]*domain.Category, int, error) {
	n := 0
	rows, err := tx.QueryContext(ctx,
		`SELECT 
			id, 
			name,
			COUNT(*) OVER()
		FROM 
			categories
		ORDER BY 
			id ASC`,
	)

	if err != nil {
		return nil, n, err
	}

	defer rows.Close()

	categories := make([]*domain.Category, 0)
	for rows.Next() {
		var category domain.Category
		if err := rows.Scan(
			&category.Id,
			&category.Name,
			&n,
		); err != nil {
			return nil, 0, err
		}

		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return categories, n, nil
}

func getCategory(ctx context.Context, tx *database.Tx, id ulid.ULID) (*domain.Category, error) {
	var category domain.Category

	err := tx.QueryRowContext(ctx,
		`SELECT 
			id, 
			name
		FROM 
			categories
		WHERE 
			id = ?`,
		id.String(),
	).Scan(
		&category.Id,
		&category.Name,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &category, nil
}

func createCategory(ctx context.Context, tx *database.Tx, id ulid.ULID, name string) error {
	_, err := tx.ExecContext(ctx,
		`INSERT INTO 
		categories(
			id,
			name
		)
		VALUES(
			?, 
			?
		)`,
		id.String(),
		name,
	)

	if err != nil {
		return err
	}

	return nil
}

func updateCategory(ctx context.Context, tx *database.Tx, id ulid.ULID, name string) error {
	if _, err := tx.ExecContext(ctx,
		`UPDATE 
			categories 
		SET
			name = ?
		WHERE 
			id = ?	
		`,
		name,
		id.String(),
	); err != nil {
		return err
	}

	return nil
}

func deleteCategory(ctx context.Context, tx *database.Tx, id ulid.ULID) error {
	if _, err := tx.ExecContext(ctx,
		`DELETE FROM 
			categories 
		WHERE 
			id = ?
		`,
		id.String(),
	); err != nil {
		return err
	}

	return nil
}
