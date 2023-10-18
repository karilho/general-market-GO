package users

import (
	"context"
	"github.com/karilho/general-market-GO/adapters/repo"
	"github.com/karilho/general-market-GO/domain"
)

type Service struct {
	usersRepo repo.Users
}

func NewService(usersRepo repo.Users) Service {
	return Service{
		usersRepo: usersRepo,
	}
}

/*
	func (s Service) UpsertUser(ctx context.Context, user domain.User) (userID int, _ error) {
		userID, err := s.usersRepo.UpsertUser(ctx, user)
		if err != nil {
			return 0, err
		}

		//implement log
		return userID, nil
	}
*/
func (s Service) GetUser(ctx context.Context, userID int) (domain.User, error) {
	user, err := s.usersRepo.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
