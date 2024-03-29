package buyers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/karilho/general-market-GO/adapters/cloud"
	"github.com/karilho/general-market-GO/adapters/repo"
	"github.com/karilho/general-market-GO/domain"
	"os"
)

type Service struct {
	buyersRepo     repo.Buyers
	storageservice domain.StorageService
}

func NewBuyerService(buyersRepo repo.Buyers, service *cloud.S3StorageService) Service {
	return Service{
		buyersRepo:     buyersRepo,
		storageservice: service,
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

	buyerJSON, err := json.Marshal(buyer)
	if err != nil {
		return 0, err
	}

	fileName := fmt.Sprintf("buyer%d.json", buyer.UserDataID)
	createdFile, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}
	defer os.Remove(createdFile.Name())

	if _, err := createdFile.Write(buyerJSON); err != nil {
		return 0, err
	}
	if err := createdFile.Close(); err != nil {
		return 0, err
	}

	fmt.Println("O json é: ", string(buyerJSON))

	err = s.storageservice.UploadFileToBucket("my-new-bucket-test-general-market", createdFile.Name())
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (s Service) GetBuyer(ctx context.Context, buyerID int) (domain.Buyers, error) {
	buyer, err := s.buyersRepo.GetBuyer(ctx, buyerID)
	if err != nil {
		return domain.Buyers{}, err
	}
	return buyer, nil
}
