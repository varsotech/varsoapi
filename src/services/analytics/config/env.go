package config

import (
	_ "embed"

	"github.com/varsotech/varsoapi/src/common/config"
)

var (
	AnalyticsDatabaseName = config.RequireEnv("ANALYTICS_DATABASE_NAME")
)
