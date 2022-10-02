package database

import (
	"context"
	"database/sql"

	"github.com/gadoma/rafapi/internal/domain"
)

var _ domain.CategoryRepository = (*CategoryRepository)(nil)

type CategoryRepository struct {
	db *DB
}

func NewCategoryRepository(db *DB) *CategoryRepository {
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

func (r *CategoryRepository) GetCategory(ctx context.Context, id int) (*domain.Category, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	return getCategory(ctx, tx, id)
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, cu *domain.CategoryUpdate) (int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	id, err := createCategory(ctx, tx, cu.Name)

	if err != nil {
		return 0, err
	}

	tx.Commit()

	return id, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, id int, cu *domain.CategoryUpdate) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := updateCategory(ctx, tx, id, cu.Name); err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteCategory(ctx, tx, id); err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func getCategories(ctx context.Context, tx *Tx) ([]*domain.Category, int, error) {
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

func getCategory(ctx context.Context, tx *Tx, id int) (*domain.Category, error) {
	var category domain.Category

	err := tx.QueryRowContext(ctx,
		`SELECT 
			id, 
			name
		FROM 
			categories
		WHERE 
			id = ?`,
		id,
	).Scan(
		&category.Id,
		&category.Name,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func createCategory(ctx context.Context, tx *Tx, name string) (int, error) {
	result, err := tx.ExecContext(ctx,
		`INSERT INTO 
		categories(
			name
		)
		VALUES(
			?
		)`,
		name,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func updateCategory(ctx context.Context, tx *Tx, id int, name string) error {
	if _, err := tx.ExecContext(ctx,
		`UPDATE 
			categories 
		SET
			name = ?
		WHERE 
			id = ?	
		`,
		name,
		id,
	); err != nil {
		return err
	}

	return nil
}

func deleteCategory(ctx context.Context, tx *Tx, id int) error {
	if _, err := tx.ExecContext(ctx,
		`DELETE FROM 
			categories 
		WHERE 
			id = ?
		`,
		id,
	); err != nil {
		return err
	}

	return nil
}
