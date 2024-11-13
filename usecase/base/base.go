package base

import (
	"rest/domain"
	"time"
)

type BaseService struct {
	repo repository
	// and also possible adapters
}

func NewUsecase(repo repository) *BaseService {
	return &BaseService{
		repo: repo,
	}
}

func (uc *BaseService) Info() (domain.Base, error) {
	time.Sleep(15 * time.Second)
	return domain.Base{Count: 10}, nil
}
