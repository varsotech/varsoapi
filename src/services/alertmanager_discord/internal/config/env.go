package config

import (
	_ "embed"

	"github.com/varsotech/varsoapi/src/common/config"
)

var (
	AlertsDiscordWebhook = config.RequireSensitiveEnv("ALERTS_DISCORD_WEBHOOK", config.AppEnv)
)
