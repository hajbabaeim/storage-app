package repository

import (
	"context"
	"errors"

	"storage-app/ent"
	"storage-app/ent/promotion"
	"storage-app/internal/model"

	_ "github.com/lib/pq"
)

type PromotionRepository struct {
	client *ent.Client
}

func NewPostgresDb(dbUri string) (*ent.Client, error) {
	client, err := ent.Open("postgres", dbUri)
	if err != nil {
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	return client, nil
}

func NewPromotionRepository(client *ent.Client) *PromotionRepository {
	return &PromotionRepository{
		client: client,
	}
}

func (r *PromotionRepository) GetByID(ctx context.Context, id string) (*model.Promotion, error) {
	p, err := r.client.Promotion.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("promotion not found")
		}
		return nil, err
	}

	return &model.Promotion{
		ID:             p.ID,
		Price:          p.Price,
		ExpirationDate: p.ExpirationDate,
	}, nil
}

func (r *PromotionRepository) Insert(ctx context.Context, m *model.Promotion) error {
	// Check if promotion with the same ID already exists
	exists, err := r.client.Promotion.Query().Where(
		promotion.ID(m.ID),
		promotion.Price(m.Price),
		promotion.ExpirationDate(m.ExpirationDate),
	).
		Exist(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil // Return nil if the promotion already exists
	}
	_, err = r.client.Promotion.
		Create().
		SetID(m.ID).
		SetPrice(m.Price).
		SetExpirationDate(m.ExpirationDate).
		Save(ctx)

	return err
}
