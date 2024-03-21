package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/varsotech/varsoapi/src/common/api"
	common_config "github.com/varsotech/varsoapi/src/common/config"
	"github.com/varsotech/varsoapi/src/services/analytics/internal/ent"
	"github.com/varsotech/varsoapi/src/services/analytics/internal/routes/log"
)

func main() {
	logrus.Infof("Starting analytics service in %s mode", common_config.AppEnv)
	defer logrus.Infof("Finished gracefully shutting down")

	err := ent.Connect(context.Background())
	if err != nil {
		logrus.WithError(err).Panicf("failed connecting to database")
	}

	router := api.NewRouter()

	// Register routes
	api.RouteGET(router, api.Public, "/", api.SuccessEndpoint)

	// Posts
	api.RouteGET(router, api.Public, "/api/v1/analytics/access_log/:ip/:uri", log.AccessLog)

	logrus.Fatal(router.ListenAndServe(":5003"))
}
