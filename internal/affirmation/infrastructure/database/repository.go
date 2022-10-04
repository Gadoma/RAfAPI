package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gadoma/rafapi/internal/affirmation/domain"
	"github.com/gadoma/rafapi/internal/common/infrastructure/database"
	"github.com/oklog/ulid/v2"
)

var _ domain.AffirmationRepository = (*AffirmationRepository)(nil)

type AffirmationRepository struct {
	db *database.DB
}

func NewAffirmationRepository(db *database.DB) *AffirmationRepository {
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

func (r *AffirmationRepository) GetAffirmation(ctx context.Context, id ulid.ULID) (*domain.Affirmation, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	return getAffirmation(ctx, tx, id)
}

func (r *AffirmationRepository) CreateAffirmation(ctx context.Context, cac *domain.CreateAffirmationCommand) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = createAffirmation(ctx, tx, cac.Id, cac.Text, cac.CategoryId)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (r *AffirmationRepository) UpdateAffirmation(ctx context.Context, id ulid.ULID, uac *domain.UpdateAffirmationCommand) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := updateAffirmation(ctx, tx, id, uac.Text, uac.CategoryId); err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (r *AffirmationRepository) DeleteAffirmation(ctx context.Context, id ulid.ULID) error {
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

func getAffirmations(ctx context.Context, tx *database.Tx) ([]*domain.Affirmation, int, error) {
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
			(*database.StringTime)(&affirmation.CreatedAt),
			(*database.StringTime)(&affirmation.UpdatedAt),
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

func getAffirmation(ctx context.Context, tx *database.Tx, id ulid.ULID) (*domain.Affirmation, error) {
	var affirmation domain.Affirmation

	err := tx.QueryRowContext(ctx,
		`SELECT 
			id, 
			text,
			category_id,
			created_at,
			updated_at
		FROM 
			affirmations
		WHERE 
			id = ?`,
		id.String(),
	).Scan(
		&affirmation.Id,
		&affirmation.Text,
		&affirmation.CategoryId,
		(*database.StringTime)(&affirmation.CreatedAt),
		(*database.StringTime)(&affirmation.UpdatedAt),
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &affirmation, nil
}

func createAffirmation(ctx context.Context, tx *database.Tx, id ulid.ULID, text string, categoryId ulid.ULID) error {
	createdAt := tx.Now.Format(time.RFC3339)
	updatedAt := createdAt

	_, err := tx.ExecContext(ctx,
		`INSERT INTO 
		affirmations(
			id,
			text,
			category_id,
			created_at,
			updated_at
		)
		VALUES(
			?, 
			?, 
			?, 
			?, 
			?
		)`,
		id.String(),
		text,
		categoryId.String(),
		createdAt,
		updatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func updateAffirmation(ctx context.Context, tx *database.Tx, id ulid.ULID, text string, categoryId ulid.ULID) error {
	updatedAt := tx.Now.Format(time.RFC3339)

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
		categoryId.String(),
		updatedAt,
		id.String(),
	); err != nil {
		return err
	}

	return nil
}

func deleteAffirmation(ctx context.Context, tx *database.Tx, id ulid.ULID) error {
	if _, err := tx.ExecContext(ctx,
		`DELETE FROM 
			affirmations 
		WHERE 
			id = ?
		`,
		id.String(),
	); err != nil {
		return err
	}

	return nil
}
