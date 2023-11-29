package domain

import "github.com/vingarcia/ksql"

var BuyersTable = ksql.NewTable("buyers", "buyer_id")

var UserDataTable = ksql.NewTable("user_data", "user_data_id")

var UsersTable = ksql.NewTable("users", "id")
