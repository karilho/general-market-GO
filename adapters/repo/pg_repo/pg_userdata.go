package pgrepo

import (
	"context"
	"github.com/karilho/general-market-GO/domain"
	"github.com/vingarcia/ksql"
)

func upsertData(ctx context.Context, db ksql.Provider, data domain.UserData) (dataId int, _ error) {
	err := db.Patch(ctx, domain.UserDataTable, &data)

	if err != nil {
		err = db.Insert(ctx, domain.UserDataTable, &data)
		if err != nil {
			return 0, domain.InternalErr("unexpected error when saving data info", map[string]interface{}{
				"data":  data,
				"error": err.Error(),
			})
		}
	}
	return data.UserDataID, nil
}
