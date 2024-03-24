package person

import (
	"context"

	"github.com/varsotech/varsoapi/src/services/app/internal/ent"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/person"
)

func UpsertBulk(ctx context.Context, personEmailToName map[string]string) ([]*build.Person, error) {
	var creators []*build.PersonCreate
	for email, name := range personEmailToName {
		creators = append(creators, ent.Database.Person.Create().SetEmail(email).SetName(name))
	}

	err := ent.Database.Person.CreateBulk(creators...).OnConflictColumns(person.FieldEmail).UpdateNewValues().Exec(ctx)
	if err != nil {
		return nil, err
	}

	var emails []string
	for email := range personEmailToName {
		emails = append(emails, email)
	}

	return ent.Database.Person.Query().Where(person.EmailIn(emails...)).All(ctx)
}
