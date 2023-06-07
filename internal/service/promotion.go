package service

import (
	"context"
	"encoding/json"
	"storage-app/ent"
	"storage-app/internal/model"
	"storage-app/internal/repository"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type PromotionService struct {
	repo *repository.PromotionRepository
	rdb  *redis.Client // Redis client
}

func NewPromotionService(repo *repository.PromotionRepository, rdb *redis.Client) *PromotionService {
	return &PromotionService{
		repo: repo,
		rdb:  rdb,
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
	// Try to get the result from Redis first
	result, err := s.rdb.Get(strconv.Itoa(id)).Result()

	if err == redis.Nil {
		// If the result is not in Redis, get it from the database
		promotion, err := s.repo.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}

		// Then store the result in Redis for next time
		promotionJSON, err := json.Marshal(promotion)
		if err != nil {
			return nil, err
		}

		err = s.rdb.Set(strconv.Itoa(id), promotionJSON, 1*time.Hour).Err()
		if err != nil {
			return nil, err
		}

		return promotion, nil
	} else if err != nil {
		return nil, err
	}

	// If the result was in Redis, deserialize it into a Promotion
	var promotion model.Promotion
	err = json.Unmarshal([]byte(result), &promotion)
	if err != nil {
		return nil, err
	}

	return &promotion, nil
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
