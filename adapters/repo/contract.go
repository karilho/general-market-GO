package repo

import (
	"context"
	"github.com/karilho/general-market-GO/domain"
)

type Users interface {
	GetUser(ctx context.Context, userID int) (domain.User, error)
	UpsertUser(ctx context.Context, user domain.User) (userID int, err error)
}

type Buyers interface {
	UpsertBuyer(ctx context.Context, buyer domain.Buyers) (buyerID int, err error)
	UpsertUserData(ctx context.Context, userData domain.UserData) (userDataID int, err error)
	GetBuyer(ctx context.Context, buyerID int) (domain.Buyers, error)
}
