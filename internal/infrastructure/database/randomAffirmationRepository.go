package database

import (
	"context"
	"strings"

	"github.com/gadoma/rafapi/internal/domain"
)

var _ domain.RandomAffirmationRepository = (*RandomAffirmationRepository)(nil)

type RandomAffirmationRepository struct {
	db *DB
}

func NewRandomAffirmationRepository(db *DB) *RandomAffirmationRepository {
	return &RandomAffirmationRepository{db: db}
}

func (r *RandomAffirmationRepository) GetRandomAffirmations(ctx context.Context, categoryIds []int) ([]*domain.RandomAffirmation, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	return getRandomAffirmations(ctx, tx, categoryIds)
}

func getRandomAffirmations(ctx context.Context, tx *Tx, categoryIds []int) ([]*domain.RandomAffirmation, error) {
	prepare := func(input []int) ([]string, []any) {
		placeholders := make([]string, 0)
		output := make([]any, 0)
		for _, i := range input {
			output = append(output, i)
			placeholders = append(placeholders, "?")
		}
		return placeholders, output
	}

	placeholders, convertedIds := prepare(categoryIds)

	rows, err := tx.QueryContext(ctx,
		`SELECT 
			ra.affirmation_text
		FROM 
		(    
			SELECT  
				c.id as category_id, c.name, a.id, a.text as affirmation_text
			FROM 
				categories c
			INNER JOIN 
				affirmations a 
			ON c.id = a.category_id 
			ORDER BY RANDOM() 
		) AS ra 
		WHERE category_id IN(`+strings.Join(placeholders, ",")+`)
		GROUP BY ra.category_id`,
		convertedIds...,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	randomaffirmations := make([]*domain.RandomAffirmation, 0)
	for rows.Next() {
		var randomaffirmation domain.RandomAffirmation
		if err := rows.Scan(
			&randomaffirmation.Text,
		); err != nil {
			return nil, err
		}

		randomaffirmations = append(randomaffirmations, &randomaffirmation)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return randomaffirmations, nil
}
