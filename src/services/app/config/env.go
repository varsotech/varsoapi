package config

import (
	_ "embed"

	"github.com/varsotech/varsoapi/src/common/config"
)

var (
	AppDatabaseName = config.RequireEnv("APP_DATABASE_NAME")
)
