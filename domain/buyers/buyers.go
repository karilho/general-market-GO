package buyers

import (
	"context"
	"github.com/karilho/general-market-GO/adapters/repo"
	"github.com/karilho/general-market-GO/domain"
)

type Service struct {
	buyersRepo repo.Buyers
}

func NewBuyerService(buyersRepo repo.Buyers) Service {
	return Service{
		buyersRepo: buyersRepo,
	}
}

func (s Service) UpsertUserData(ctx context.Context, buyer domain.UserData) (userDataId int, _ error) {
	generateDataId, err := s.buyersRepo.UpsertUserData(ctx, buyer)
	if err != nil {
		return 0, err
	}

	return generateDataId, nil
}

func (s Service) UpsertBuyer(ctx context.Context, buyer domain.Buyers) (buyerId int, _ error) {

	_, err := s.buyersRepo.UpsertBuyer(ctx, buyer)
	if err != nil {
		return 0, err
	}

	return buyerId, nil
}

func (s Service) GetBuyer(ctx context.Context, buyerID int) (domain.Buyers, error) {
	buyer, err := s.buyersRepo.GetBuyer(ctx, buyerID)
	if err != nil {
		return domain.Buyers{}, err
	}
	return buyer, nil
}
