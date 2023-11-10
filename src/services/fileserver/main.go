package main

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/varsotech/varsoapi/src/common/api"
	common_config "github.com/varsotech/varsoapi/src/common/config"
	"github.com/varsotech/varsoapi/src/services/fileserver/ent"
	"github.com/varsotech/varsoapi/src/services/fileserver/routes/download"
	"github.com/varsotech/varsoapi/src/services/fileserver/routes/stitch"
	"github.com/varsotech/varsoapi/src/services/fileserver/routes/upload"
)

func main() {
	logrus.Infof("Starting fileserver in %s mode", common_config.AppEnv)
	defer logrus.Infof("Finished gracefully shutting down")

	err := ent.Connect(context.Background())
	if err != nil {
		logrus.WithError(err).Panicf("failed connecting to database")
	}

	router := api.NewRouter()

	api.RouteGET(router, api.Public, "/", api.SuccessEndpoint)

	api.RouteGET(router, api.Public, "/api/v1/fileserver/file/public/:filename", download.DownloadPublic)
	api.RoutePOST(router, api.User, "/api/v1/fileserver/file/public", upload.UploadPublic)

	api.RouteGET(router, api.User, "/api/v1/fileserver/file/default/:filename", download.Download)
	api.RoutePOST(router, api.User, "/api/v1/fileserver/file/default", upload.Upload)

	api.RoutePOST(router, api.Internal, "/api/v1/fileserver/stitch", stitch.Stitch)

	log.Fatal(router.ListenAndServe(":5002"))
}
