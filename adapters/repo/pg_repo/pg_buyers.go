package pgrepo

import (
	"context"
	"github.com/karilho/general-market-GO/domain"
	"github.com/vingarcia/ksql"
)

func upsertBuyer(ctx context.Context, db ksql.Provider, buyer domain.Buyers) (buyerId int, _ error) {
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

	return buyer.BuyerID, nil
}

func getBuyer(ctx context.Context, db ksql.Provider, buyerID int) (domain.Buyers, error) {
	var row struct {
		Buyers   domain.Buyers   `tablename:"b"`
		UserData domain.UserData `tablename:"u"`
	}
	err := db.QueryOne(ctx, &row, "FROM buyers as b JOIN user_data as u ON b.user_data_id = u.user_data_id WHERE b.user_data_id = $1", buyerID)
	if err == ksql.ErrRecordNotFound {
		return domain.Buyers{}, domain.NotFoundErr("no buyer found with provided id", map[string]interface{}{
			"buyer_id": buyerID,
		})
	}
	if err != nil {
		return domain.Buyers{}, domain.InternalErr("unexpected error when fetching buyer", map[string]interface{}{
			"buyer_id": buyerID,
			"error":    err.Error(),
		})
	}

	row.Buyers.UserData = &row.UserData
	return row.Buyers, nil
}
