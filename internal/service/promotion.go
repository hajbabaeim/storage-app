package service

import (
	"context"
	"storage-app/ent"
	"storage-app/internal/model"
	"storage-app/internal/repository"
)

type PromotionService struct {
	repo *repository.PromotionRepository
}

func NewPromotionService(repo *repository.PromotionRepository) *PromotionService {
	return &PromotionService{
		repo: repo,
	}
}

func (s *PromotionService) BatchInsert(ctx context.Context, batch []*model.Promotion) error {
	return s.repo.BatchInsert(ctx, batch)
}

func (s *PromotionService) BeginTransaction(ctx context.Context) (*ent.Tx, error) {
	return s.repo.BeginTransaction(ctx)
}

func (s *PromotionService) RollbackTransaction(ctx context.Context, tx *ent.Tx) error {
	return s.repo.RollbackTransaction(ctx, tx)
}

func (s *PromotionService) DeleteAll(ctx context.Context) (int, error) {
	return s.repo.DeleteAll(ctx)
}

// InsertBatch inserts multiple promotions into the database.
func (s *PromotionService) InsertBatch(ctx context.Context, promotions []*model.Promotion) error {
	for _, promotion := range promotions {
		if err := s.repo.Insert(ctx, promotion); err != nil {
			return err
		}
	}
	return nil
}

func (svc *PromotionService) GetPromotion(ctx context.Context, id int) (*model.Promotion, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *PromotionService) AddPromotion(ctx context.Context, promotion *model.Promotion) error {
	return svc.repo.Insert(ctx, promotion)
}

func (s *PromotionService) GetByID(ctx context.Context, id int) (*model.Promotion, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PromotionService) DeleteByID(ctx context.Context, id int) (int, error) {
	return s.repo.DeleteByID(ctx, id)
}

func (s *PromotionService) Insert(ctx context.Context, promotion *model.Promotion) error {
	return s.repo.Insert(ctx, promotion)
}

func (s *PromotionService) CommitTransaction(ctx context.Context, tx *ent.Tx) error {
	return s.repo.CommitTransaction(ctx, tx)
}
