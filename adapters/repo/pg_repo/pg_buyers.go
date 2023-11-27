package pgrepo

import (
	"context"
	"github.com/karilho/general-market-GO/domain"
	"github.com/vingarcia/ksql"
)

func upsertBuyer(ctx context.Context, db ksql.Provider, buyer domain.Buyers) (buyerId int, _ error) {
	//now := time.Now()
	//buyer.UpdatedAt = &now
	err := db.Patch(ctx, domain.BuyersTable, &buyer)
	if err != nil {
		err = db.Insert(ctx, domain.BuyersTable, &buyer)
		if err != nil {
			return 0, domain.InternalErr("unexpected error when saving buyer", map[string]interface{}{
				"buyer": buyer,
				"error": err.Error(),
			})
		}
	}

	if err != nil {
		return 0, domain.InternalErr("unexpected error when saving buyer", map[string]interface{}{
			"buyer": buyer,
			"error": err.Error(),
		})
	}

	return buyer.BuyerID, nil
}

func getBuyer(ctx context.Context, db ksql.Provider, buyerID int) (map[string]any, error) {
	var row struct {
		Buyers   domain.Buyers   `tablename:"b"`
		UserData domain.UserData `tablename:"user_data"`
	}
	err := db.QueryOne(ctx, &row, "FROM buyers as b JOIN user_data ON b.user_data_id = user_data.user_data_id WHERE b.user_data_id = $1", buyerID)
	if err == ksql.ErrRecordNotFound {
		return map[string]interface{}{}, domain.NotFoundErr("no buyer found with provided id", map[string]interface{}{
			"buyer_id": buyerID,
		})
	}
	if err != nil {
		return map[string]interface{}{}, domain.InternalErr("unexpected error when fetching buyer", map[string]interface{}{
			"buyer_id": buyerID,
			"error":    err.Error(),
		})
	}

	return map[string]interface{}{
		"buyer": row,
	}, nil
}
