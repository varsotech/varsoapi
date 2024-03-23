package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/varsotech/varsoapi/src/common/api"
	common_config "github.com/varsotech/varsoapi/src/common/config"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/organization"
	"github.com/varsotech/varsoapi/src/services/app/internal/routes"
)

func main() {
	logrus.Infof("Starting app service in %s mode", common_config.AppEnv)
	defer logrus.Infof("Finished gracefully shutting down")

	ctx := context.Background()

	err := ent.Connect(ctx)
	if err != nil {
		logrus.WithError(err).Panicf("failed connecting to database")
	}

	err = organization.Initialize(ctx)
	if err != nil {
		logrus.WithError(err).Panicf("failed initializing organization module")
	}

	router := api.NewRouter()

	// Register routes
	api.RouteGET(router, api.Public, "/api/v1/app", api.SuccessEndpoint)
	api.RouteGET(router, api.Public, "/api/v1/app/organization", routes.GetOrganizations)
	api.RouteGET(router, api.Public, "/api/v1/app/news", routes.GetNews)

	logrus.Fatal(router.ListenAndServe(":5001"))
}
