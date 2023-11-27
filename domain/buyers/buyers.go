package buyers

import (
	"context"
	"fmt"
	"github.com/karilho/general-market-GO/adapters/repo"
	"github.com/karilho/general-market-GO/domain"
)

type Service struct {
	buyersRepo repo.Buyers
}

// Call the repository to use on services, like a @Autowired.
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

	fmt.Printf("generateDataId: %v\n", generateDataId)

	//implement log
	return generateDataId, nil
}

func (s Service) UpsertBuyer(ctx context.Context, buyer domain.Buyers) (buyerId int, _ error) {

	generateBuyerId, err := s.buyersRepo.UpsertBuyer(ctx, buyer)
	if err != nil {
		return 0, err
	}

	fmt.Printf("generateBuyerId: %v\n", generateBuyerId)

	//implement log
	return buyerId, nil
}

func (s Service) GetBuyer(ctx context.Context, buyerID int) (map[string]any, error) {
	buyer, err := s.buyersRepo.GetBuyer(ctx, buyerID)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return buyer, nil
}
