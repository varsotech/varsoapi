package config

import (
	_ "embed"

	"github.com/varsotech/varsoapi/src/common/config"
)

var (
	FileServerDatabaseName = config.RequireEnv("FILESERVER_DATABASE_NAME")
)
