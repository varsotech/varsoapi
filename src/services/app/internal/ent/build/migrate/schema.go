// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "unique_name", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "website_url", Type: field.TypeString},
		{Name: "rss_feed_url", Type: field.TypeString},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "organization_uuid",
				Unique:  true,
				Columns: []*schema.Column{OrganizationsColumns[0]},
			},
			{
				Name:    "organization_unique_name",
				Unique:  true,
				Columns: []*schema.Column{OrganizationsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrganizationsTable,
	}
)

func init() {
}
