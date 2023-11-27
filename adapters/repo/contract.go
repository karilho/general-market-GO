package repo

import (
	"context"
	"github.com/karilho/general-market-GO/domain"
)

// UsersRepo represents the operations we use for
// retrieving a user from a persistent storage
type Users interface {
	GetUser(ctx context.Context, userID int) (domain.User, error)
	UpsertUser(ctx context.Context, user domain.User) (userID int, err error)
	//GetUserByEmail(ctx context.Context, email string) (User, error)
}

type Buyers interface {
	UpsertBuyer(ctx context.Context, buyer domain.Buyers) (buyerID int, err error)
	UpsertUserData(ctx context.Context, userData domain.UserData) (userDataID int, err error)

	GetBuyer(ctx context.Context, buyerID int) (map[string]any, error)
}
