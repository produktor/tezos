package migrations

import (
	"fmt"
	"link_api/internal/config"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func Init(config *config.Config) error {
	db, err := goose.OpenDBWithDriver("postgres", config.MigrationsDSN)
	if err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if err := goose.Run("up", db, "migrations"); err != nil {
		return err
	}

	return nil
}
