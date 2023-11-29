package pgrepo

import (
	"context"
	"github.com/karilho/general-market-GO/domain"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

// Repositories implements the repo.Users interface by using the ksql database.
type Repositories struct {
	db ksql.Provider
}

// New instantiates a new Repositories
func New(ctx context.Context, postgresURL string) (Repositories, error) {
	db, err := kpgx.New(ctx, postgresURL, ksql.Config{})
	if err != nil {
		return Repositories{}, err
	}

	return Repositories{
		db: db,
	}, nil
}

func (r Repositories) UpsertUser(ctx context.Context, user domain.User) (userID int, _ error) {
	return upsertUser(ctx, r.db, user)
}

func (r Repositories) GetUser(ctx context.Context, userID int) (domain.User, error) {
	return getUser(ctx, r.db, userID)
}

func (r Repositories) UpsertBuyer(ctx context.Context, buyer domain.Buyers) (buyerID int, _ error) {
	return upsertBuyer(ctx, r.db, buyer)
}

func (r Repositories) UpsertUserData(ctx context.Context, userData domain.UserData) (userDataID int, _ error) {
	return upsertData(ctx, r.db, userData)
}

func (r Repositories) GetBuyer(ctx context.Context, buyerID int) (domain.Buyers, error) {
	return getBuyer(ctx, r.db, buyerID)
}
