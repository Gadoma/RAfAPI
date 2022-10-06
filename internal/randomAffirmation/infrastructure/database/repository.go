package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/gadoma/rafapi/internal/common/infrastructure/database"
	"github.com/gadoma/rafapi/internal/randomAffirmation/domain"
	"github.com/oklog/ulid/v2"
)

var _ domain.RandomAffirmationRepository = (*RandomAffirmationRepository)(nil)

type RandomAffirmationRepository struct {
	db *database.DB
}

func NewRandomAffirmationRepository(db *database.DB) *RandomAffirmationRepository {
	return &RandomAffirmationRepository{db: db}
}

func (r *RandomAffirmationRepository) GetRandomAffirmations(ctx context.Context, categoryIds []ulid.ULID) ([]*domain.RandomAffirmation, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func(tx *database.Tx) {
		err := tx.Rollback()
		if err != nil {
			panic(err)
		}
	}(tx)

	return getRandomAffirmations(ctx, tx, categoryIds)
}

func getRandomAffirmations(ctx context.Context, tx *database.Tx, categoryIds []ulid.ULID) ([]*domain.RandomAffirmation, error) {
	convertedIds := make([]any, 0)
	whereCondition := ""
	limit := "LIMIT 10"

	if len(categoryIds) > 0 {
		placeholders := make([]string, 0)

		for _, i := range categoryIds {
			convertedIds = append(convertedIds, i.String())
			placeholders = append(placeholders, "?")
		}

		whereCondition = "WHERE category_id IN(" + strings.Join(placeholders, ",") + ")"
		limit = ""
	}

	rows, err := tx.QueryContext(ctx,
		fmt.Sprintf(`
		SELECT
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
		%s
		GROUP BY ra.category_id
		%s`, whereCondition, limit),
		convertedIds...,
	)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	randomAffirmations := make([]*domain.RandomAffirmation, 0)
	for rows.Next() {
		var randomAffirmation domain.RandomAffirmation
		if err := rows.Scan(&randomAffirmation.Text); err != nil {
			return nil, err
		}
		randomAffirmations = append(randomAffirmations, &randomAffirmation)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return randomAffirmations, nil
}
