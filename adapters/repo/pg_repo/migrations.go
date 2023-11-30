package pgrepo

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"net/http"
)

func MigrateDB(migrationsDir http.FileSystem, DBurl string) {
	log.Info("Starting migrations...")

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

	log.Info(fmt.Sprintf("Applied %d migrations!\n", n))
}
