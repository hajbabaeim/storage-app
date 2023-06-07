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

func (s *PromotionService) BeginTransaction(ctx context.Context) (*ent.Tx, error) {
	return s.repo.BeginTransaction(ctx)
}

func (s *PromotionService) CommitTransaction(ctx context.Context, tx *ent.Tx) error {
	return s.repo.CommitTransaction(ctx, tx)
}

func (s *PromotionService) RollbackTransaction(ctx context.Context, tx *ent.Tx) error {
	return s.repo.RollbackTransaction(ctx, tx)
}

// func (s *PromotionService) GetByDBID(ctx context.Context, dbID int64) (*model.Promotion, error) {
// 	return s.repo.GetByDBID(ctx, dbID)
// }
