package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/varsotech/varsoapi/src/common/config"
	"github.com/varsotech/varsoapi/src/common/maingroup"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent"
	"github.com/varsotech/varsoapi/src/services/auth/internal/modules/role"
	"github.com/varsotech/varsoapi/src/services/auth/internal/modules/user"
	"github.com/varsotech/varsoapi/src/services/auth/internal/router"
)

func main() {
	logrus.Infof("Starting auth service in %s mode", config.AppEnv)
	defer logrus.Infof("Finished gracefully shutting down")

	ctx := context.Background()

	err := ent.Connect(ctx)
	if err != nil {
		logrus.WithError(err).Panicf("failed connecting to database")
	}

	err = role.Initialize(ctx)
	if err != nil {
		logrus.WithError(err).Panicf("failed to initialize role module")
	}

	err = user.Initialize(ctx)
	if err != nil {
		logrus.WithError(err).Panicf("failed to initialize user module")
	}

	mainGroup := maingroup.New(ctx)

	mainGroup.Go(router.ListenAndServe)

	logrus.Panic(mainGroup.Wait())
}
