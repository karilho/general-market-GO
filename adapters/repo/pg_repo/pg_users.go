package pgrepo

import (
	"context"
	"github.com/karilho/general-market-GO/domain"
	"github.com/vingarcia/ksql"
)

func upsertUser(ctx context.Context, db ksql.Provider, user domain.User) (userID int, _ error) {
	err := db.Patch(ctx, domain.UsersTable, &user)
	if err != nil {
		err = db.Insert(ctx, domain.UsersTable, &user)
		if err != nil {
			return 0, domain.InternalErr("unexpected error when saving user", map[string]interface{}{
				"user":  user,
				"error": err.Error(),
			})
		}
	}

	return user.ID, nil
}

func getUser(ctx context.Context, db ksql.Provider, userID int) (domain.User, error) {
	var user domain.User
	err := db.QueryOne(ctx, &user, "FROM users WHERE id = $1", userID)
	if err == ksql.ErrRecordNotFound {
		return domain.User{}, domain.NotFoundErr("no user found with provided id", map[string]interface{}{
			"user_id": userID,
		})
	}
	if err != nil {
		return domain.User{}, domain.InternalErr("unexpected error when fetching user", map[string]interface{}{
			"user_id": userID,
			"error":   err.Error(),
		})
	}

	return user, nil
}
