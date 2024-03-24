package main

import (
	"context"

	"github.com/sirupsen/logrus"
	common_config "github.com/varsotech/varsoapi/src/common/config"
	"github.com/varsotech/varsoapi/src/common/maingroup"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/news"
	"github.com/varsotech/varsoapi/src/services/app/internal/router"
)

func main() {
	logrus.Infof("Starting app service in %s mode", common_config.AppEnv)
	defer logrus.Infof("Finished gracefully shutting down")

	ctx := context.Background()

	err := ent.Connect(ctx)
	if err != nil {
		logrus.WithError(err).Panicf("failed connecting to database")
	}

	err = news.Initialize(ctx)
	if err != nil {
		logrus.WithError(err).Panicf("failed initializing organization module")
	}

	mainGroup := maingroup.New(ctx)

	mainGroup.Go(router.ListenAndServe)
	mainGroup.Go(news.SyncRSSFeeds)

	logrus.Panic(mainGroup.Wait())
}
