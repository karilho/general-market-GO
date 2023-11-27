package pgrepo

import (
	"context"
	"github.com/karilho/general-market-GO/domain"
	"github.com/vingarcia/ksql"
)

/*
// This example function shows how we can have a transaction that reuses other repository
// functions in order to avoid code duplication.
//
// Note that the `db` argument might be either a transaction or a normal db connection pool.
//
// If `db` is a normal db connection pool `db.Transaction()` starts a new transaction,
// if not, `db.Transaction()` will just reuse the existing transaction.
//
// The same is valid for all the other private repo functions that we call, none of them
// need to know whether the input db is a transaction or not, making it possible
// to decide whether they will be used inside a transaction or not when we call them.


func changeUserEmail(ctx context.Context, db ksql.Provider, userID int, newEmail string) error {
	return db.Transaction(ctx, func(db ksql.Provider) error {
		user, err := getUser(ctx, db, userID)
		if err != nil {
			return err
		}

		// If there is nothing to do, just return:
		if user.Email == newEmail {
			return nil
		}

		_, err = getUserByEmail(ctx, db, newEmail)
		if err != ksql.ErrRecordNotFound {
			return fmt.Errorf("can't change user email to '%s': this email is already used by other user", newEmail)
		}
		if err != nil {
			return err
		}

		user.Email = newEmail
		_, err = upsertUser(ctx, db, user)
		return err
	})
}

*/

// Keeping the implementation deatached like this and passing the database provider interface
// as an argument allows you to include several diferent calls in a same transaction.
func upsertUser(ctx context.Context, db ksql.Provider, user domain.User) (userID int, _ error) {
	//now := time.Now()
	//user.UpdatedAt = &now
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

	if err != nil {
		return 0, domain.InternalErr("unexpected error when saving user", map[string]interface{}{
			"user":  user,
			"error": err.Error(),
		})
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

func getUserByEmail(ctx context.Context, db ksql.Provider, email string) (domain.User, error) {
	var user domain.User
	err := db.QueryOne(ctx, &user, "FROM users WHERE email = $1", email)
	if err == ksql.ErrRecordNotFound {
		return domain.User{}, domain.NotFoundErr("no user found with provided email", map[string]interface{}{
			"email": email,
		})
	}
	if err != nil {
		return domain.User{}, domain.InternalErr("unexpected error when fetching user by email", map[string]interface{}{
			"email": email,
			"error": err.Error(),
		})
	}

	return user, nil
}
