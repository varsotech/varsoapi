package main

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/common/config"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent"
	"github.com/varsotech/varsoapi/src/services/auth/internal/routes/discord_login"
	"github.com/varsotech/varsoapi/src/services/auth/internal/routes/easyregister"
	"github.com/varsotech/varsoapi/src/services/auth/internal/routes/internal_login"
	"github.com/varsotech/varsoapi/src/services/auth/internal/routes/login"
	"github.com/varsotech/varsoapi/src/services/auth/internal/routes/register"
	"github.com/varsotech/varsoapi/src/services/auth/internal/routes/user"
)

func main() {
	logrus.Infof("Starting auth service in %s mode", config.AppEnv)
	defer logrus.Infof("Finished gracefully shutting down")

	err := ent.Connect(context.Background())
	if err != nil {
		logrus.WithError(err).Panicf("failed connecting to database")
	}

	router := api.NewRouter()

	api.RouteGET(router, api.Public, "/", api.SuccessEndpoint)
	api.RoutePOST(router, api.Public, "/api/v1/auth/login", login.Login)
	api.RoutePOST(router, api.Public, "/api/v1/auth/register", register.Register)
	api.RoutePOST(router, api.Public, "/api/v1/auth/register/easy", easyregister.EasyRegister)
	api.RoutePOST(router, api.Public, "/api/v1/auth/discord_login", discord_login.DiscordUserLogin)
	api.RoutePOST(router, api.Public, "/api/v1/auth/internal_login", internal_login.InternalLogin)

	api.RouteGET(router, api.Internal, "/api/v1/auth/user/adminget/:uuid", user.AdminGetUser)
	api.RouteGET(router, api.Internal, "/api/v1/auth/user/admininspect/:discordid", user.AdminInspectUser)
	api.RoutePOST(router, api.Internal, "/api/v1/auth/user/ban/:uuid", user.BanUser)
	api.RoutePOST(router, api.Internal, "/api/v1/auth/user/unban/:uuid", user.UnbanUser)

	log.Fatal(router.ListenAndServe(":5000"))
}
