package ebook

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct {
	database *pgxpool.Pool
}

func NewPostgresRepository(database *pgxpool.Pool) Repository {
	return &postgresRepository{
		database: database,
	}
}

func (repository *postgresRepository) FindAll(
	ctx context.Context,
) ([]Ebook, error) {
	const query = `
		SELECT
			e.id,
			e.title,
			e.slug,
			e.author,
			c.name,
			e.description,
			e.price,
			COALESCE(e.cover_path, '')
		FROM ebooks AS e
		INNER JOIN categories AS c
			ON c.id = e.category_id
		WHERE e.status = 'published'
		ORDER BY e.published_at DESC NULLS LAST, e.id DESC
	`

	rows, err := repository.database.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query ebooks: %w", err)
	}
	defer rows.Close()

	ebooks := make([]Ebook, 0)

	for rows.Next() {
		var item Ebook
		var coverPath string

		err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Slug,
			&item.Author,
			&item.Category,
			&item.Description,
			&item.Price,
			&coverPath,
		)
		if err != nil {
			return nil, fmt.Errorf("scan ebook: %w", err)
		}

		item.IsFree = item.Price == 0

		if coverPath != "" {
			item.CoverURL = "/media/" + coverPath
		}

		ebooks = append(ebooks, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate ebooks: %w", err)
	}

	return ebooks, nil
}
