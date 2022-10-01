package database

import (
	"context"
	"time"

	"github.com/gadoma/rafapi/internal/domain"
)

var _ domain.AffirmationRepository = (*AffirmationRepository)(nil)

type AffirmationRepository struct {
	db *DB
}

func NewAffirmationRepository(db *DB) *AffirmationRepository {
	return &AffirmationRepository{db: db}
}

func (r *AffirmationRepository) GetAffirmations(ctx context.Context) ([]*domain.Affirmation, int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()

	return getAffirmations(ctx, tx)
}

func (r *AffirmationRepository) GetAffirmation(ctx context.Context, id int) (*domain.Affirmation, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	return getAffirmation(ctx, tx, id)
}

func (r *AffirmationRepository) CreateAffirmation(ctx context.Context, au *domain.AffirmationUpdate) (int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	id, err := createAffirmation(ctx, tx, au.Text, au.CategoryId)

	if err != nil {
		return 0, err
	}

	tx.Commit()

	return id, nil
}

func (r *AffirmationRepository) UpdateAffirmation(ctx context.Context, id int, au *domain.AffirmationUpdate) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := updateAffirmation(ctx, tx, id, au.Text, au.CategoryId); err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (r *AffirmationRepository) DeleteAffirmation(ctx context.Context, id int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteAffirmation(ctx, tx, id); err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func getAffirmations(ctx context.Context, tx *Tx) ([]*domain.Affirmation, int, error) {
	n := 0
	rows, err := tx.QueryContext(ctx,
		`SELECT 
			id, 
			text,
			category_id,
			created_at,
			updated_at,
			COUNT(*) OVER()
		FROM 
			affirmations
		ORDER BY 
			id ASC`,
	)

	if err != nil {
		return nil, n, err
	}

	defer rows.Close()

	affirmations := make([]*domain.Affirmation, 0)
	for rows.Next() {
		var affirmation domain.Affirmation
		if err := rows.Scan(
			&affirmation.Id,
			&affirmation.Text,
			&affirmation.CategoryId,
			(*StringTime)(&affirmation.CreatedAt),
			(*StringTime)(&affirmation.UpdatedAt),
			&n,
		); err != nil {
			return nil, 0, err
		}

		affirmations = append(affirmations, &affirmation)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return affirmations, n, nil
}

func getAffirmation(ctx context.Context, tx *Tx, id int) (*domain.Affirmation, error) {
	rows, err := tx.QueryContext(ctx,
		`SELECT 
			id, 
			text,
			category_id,
			created_at,
			updated_at
		FROM 
			affirmations
		WHERE 
			id = ?
		ORDER BY 
			id ASC`,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var affirmation domain.Affirmation
		if err := rows.Scan(
			&affirmation.Id,
			&affirmation.Text,
			&affirmation.CategoryId,
			(*StringTime)(&affirmation.CreatedAt),
			(*StringTime)(&affirmation.UpdatedAt),
		); err != nil {
			return nil, err
		}
		return &affirmation, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

func createAffirmation(ctx context.Context, tx *Tx, text string, categoryId int) (int, error) {
	createdAt := tx.now.Format(time.RFC3339)
	updatedAt := createdAt

	result, err := tx.ExecContext(ctx,
		`INSERT INTO 
		affirmations(
			text,
			category_id,
			created_at,
			updated_at
		)
		VALUES(
			?, 
			?, 
			?, 
			?
		)`,
		text,
		categoryId,
		createdAt,
		updatedAt,
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

func updateAffirmation(ctx context.Context, tx *Tx, id int, text string, categoryId int) error {
	updatedAt := tx.now.Format(time.RFC3339)

	if _, err := tx.ExecContext(ctx,
		`UPDATE 
			affirmations 
		SET
			text = ?,
			category_id = ?,
			updated_at = ?
		WHERE 
			id = ?	
		`,
		text,
		categoryId,
		updatedAt,
		id,
	); err != nil {
		return err
	}

	return nil
}

func deleteAffirmation(ctx context.Context, tx *Tx, id int) error {
	if _, err := tx.ExecContext(ctx,
		`DELETE FROM 
			affirmations 
		WHERE 
			id = ?
		`,
		id,
	); err != nil {
		return err
	}

	return nil
}
