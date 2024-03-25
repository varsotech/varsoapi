package router

import (
	"context"

	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/app/internal/router/routes"
)

func ListenAndServe(ctx context.Context) error {
	router := api.NewRouter()

	// Register routes
	api.RouteGET(router, api.Public, "/api/v1/app", api.SuccessEndpoint)
	api.RouteGET(router, api.Public, "/api/v1/app/organization", routes.GetOrganizations)
	api.RouteGET(router, api.Public, "/api/v1/app/news", routes.GetNews)
	api.RoutePOST(router, api.User, "/api/v1/app/news/item/blur/toggle", routes.BlurToggle)

	return router.ListenAndServe(":5001")
}
