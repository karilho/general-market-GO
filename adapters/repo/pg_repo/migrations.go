package pgrepo

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"net/http"
)

func MigrateDB(migrationsDir http.FileSystem, DBurl string) {
	migrationSource := &migrate.HttpFileSystemMigrationSource{
		FileSystem: migrationsDir,
	}
	db, err := sql.Open("postgres", DBurl)
	if err != nil {
		panic(err)
	}

	n, err := migrate.Exec(db, "postgres", migrationSource, migrate.Up)
	if err != nil {
		panic(err)
	}

	//TODO -> Put log instead
	fmt.Printf("Applied %d migrations!\n", n)
}
