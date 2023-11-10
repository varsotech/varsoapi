package ent

import (
	"context"
	"embed"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	common_config "github.com/varsotech/varsoapi/src/common/config"
	"github.com/varsotech/varsoapi/src/common/database"
	"github.com/varsotech/varsoapi/src/services/app/config"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build"
)

//go:embed migrations/*.sql
var fs embed.FS

var Database *build.Client

func Connect(ctx context.Context) error {
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return errors.Wrapf(err, "failed creating new iofs for migrations dir")
	}

	err = database.ApplyMigrations(ctx, d, config.AppDatabaseName)
	if err != nil {
		return errors.Wrapf(err, "failed applying migrations")
	}

	connectionString := common_config.PostgresConnectionString(config.AppDatabaseName)

	client, err := build.Open("postgres", connectionString)
	if err != nil {
		return errors.Wrapf(err, "failed opening connection to postgres")
	}

	Database = client
	return nil
}
