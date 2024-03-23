package organization

import (
	"github.com/varsotech/varsoapi/src/services/app/client/models"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build"
)

func TranslateOrganizations(posts []*build.Organization) []*models.Organization {
	var translated []*models.Organization

	for _, post := range posts {
		translated = append(translated, TranslateOrganization(post))
	}

	return translated
}

func TranslateOrganization(org *build.Organization) *models.Organization {
	return &models.Organization{
		Uuid:       org.ID.String(),
		Name:       org.Name,
		RssFeedUrl: org.RssFeedURL,
	}
}
