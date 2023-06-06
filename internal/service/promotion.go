package service

import (
	"context"
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

func (svc *PromotionService) GetPromotion(ctx context.Context, id string) (*model.Promotion, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *PromotionService) AddPromotion(ctx context.Context, promotion *model.Promotion) error {
	return svc.repo.Insert(ctx, promotion)
}

func (s *PromotionService) GetByID(ctx context.Context, id string) (*model.Promotion, error) {
	return s.repo.GetByID(ctx, id)
}
