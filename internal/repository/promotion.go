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

func (r *PromotionRepository) GetByID(ctx context.Context, id int) (*model.Promotion, error) {
	p, err := r.client.Promotion.
		Query().
		Where(promotion.ID(id)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("promotion not found")
		}
		return nil, err
	}

	return &model.Promotion{
		ID:             p.ID,
		PID:            p.Pid,
		Price:          p.Price,
		ExpirationDate: p.ExpirationDate,
	}, nil
}

func (r *PromotionRepository) Insert(ctx context.Context, m *model.Promotion) error {
	// Check if promotion with the same info already exists
	exists, err := r.client.Promotion.Query().Where(
		promotion.Pid(m.PID),
		promotion.ID(m.ID),
		promotion.Price(m.Price),
	).Exist(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil // Return nil if the promotion already exists
	}

	_, err = r.client.Promotion.
		Create().
		SetID(m.ID).
		SetPid(m.PID).
		SetPrice(m.Price).
		SetExpirationDate(m.ExpirationDate).
		Save(ctx)

	return err
}

func (r *PromotionRepository) DeleteByID(ctx context.Context, id int) (int, error) {
	return r.client.Promotion.Delete().Where(promotion.ID(id)).Exec(ctx)
}

func (r *PromotionRepository) CommitTransaction(ctx context.Context, tx *ent.Tx) error {
	return tx.Commit()
}

func (r *PromotionRepository) BatchInsert(ctx context.Context, batch []*model.Promotion) error {
	for _, p := range batch {
		_, err := r.client.Promotion.
			Create().
			SetID(p.ID).
			SetPid(p.PID).
			SetPrice(p.Price).
			SetExpirationDate(p.ExpirationDate).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PromotionRepository) BeginTransaction(ctx context.Context) (*ent.Tx, error) {
	return r.client.Tx(ctx)
}

func (r *PromotionRepository) RollbackTransaction(ctx context.Context, tx *ent.Tx) error {
	return tx.Rollback()
}

func (r *PromotionRepository) DeleteAll(ctx context.Context) (int, error) {
	return r.client.Promotion.Delete().Exec(ctx)
}
