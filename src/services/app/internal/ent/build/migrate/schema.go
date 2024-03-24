// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// NewsItemsColumns holds the columns for the "news_items" table.
	NewsItemsColumns = []*schema.Column{
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "rss_guid", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "link", Type: field.TypeString},
		{Name: "links", Type: field.TypeJSON},
		{Name: "item_publish_time", Type: field.TypeTime, Nullable: true},
		{Name: "item_update_time", Type: field.TypeTime, Nullable: true},
		{Name: "image_url", Type: field.TypeString},
		{Name: "image_title", Type: field.TypeString},
		{Name: "categories", Type: field.TypeJSON},
		{Name: "news_item_feed", Type: field.TypeUUID, Nullable: true},
	}
	// NewsItemsTable holds the schema information for the "news_items" table.
	NewsItemsTable = &schema.Table{
		Name:       "news_items",
		Columns:    NewsItemsColumns,
		PrimaryKey: []*schema.Column{NewsItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "news_items_rss_feeds_feed",
				Columns:    []*schema.Column{NewsItemsColumns[14]},
				RefColumns: []*schema.Column{RssFeedsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "newsitem_uuid",
				Unique:  true,
				Columns: []*schema.Column{NewsItemsColumns[0]},
			},
			{
				Name:    "newsitem_rss_guid",
				Unique:  true,
				Columns: []*schema.Column{NewsItemsColumns[3]},
			},
		},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "unique_name", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "website_url", Type: field.TypeString},
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
	// PersonsColumns holds the columns for the "persons" table.
	PersonsColumns = []*schema.Column{
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
	}
	// PersonsTable holds the schema information for the "persons" table.
	PersonsTable = &schema.Table{
		Name:       "persons",
		Columns:    PersonsColumns,
		PrimaryKey: []*schema.Column{PersonsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "person_uuid",
				Unique:  true,
				Columns: []*schema.Column{PersonsColumns[0]},
			},
			{
				Name:    "person_email",
				Unique:  true,
				Columns: []*schema.Column{PersonsColumns[3]},
			},
		},
	}
	// RssFeedsColumns holds the columns for the "rss_feeds" table.
	RssFeedsColumns = []*schema.Column{
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "rss_feed_url", Type: field.TypeString, Unique: true},
		{Name: "organization_feeds", Type: field.TypeUUID, Nullable: true},
	}
	// RssFeedsTable holds the schema information for the "rss_feeds" table.
	RssFeedsTable = &schema.Table{
		Name:       "rss_feeds",
		Columns:    RssFeedsColumns,
		PrimaryKey: []*schema.Column{RssFeedsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rss_feeds_organizations_feeds",
				Columns:    []*schema.Column{RssFeedsColumns[4]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "rssfeed_uuid",
				Unique:  true,
				Columns: []*schema.Column{RssFeedsColumns[0]},
			},
			{
				Name:    "rssfeed_rss_feed_url",
				Unique:  true,
				Columns: []*schema.Column{RssFeedsColumns[3]},
			},
		},
	}
	// NewsItemAuthorsColumns holds the columns for the "news_item_authors" table.
	NewsItemAuthorsColumns = []*schema.Column{
		{Name: "news_item_id", Type: field.TypeUUID},
		{Name: "person_id", Type: field.TypeUUID},
	}
	// NewsItemAuthorsTable holds the schema information for the "news_item_authors" table.
	NewsItemAuthorsTable = &schema.Table{
		Name:       "news_item_authors",
		Columns:    NewsItemAuthorsColumns,
		PrimaryKey: []*schema.Column{NewsItemAuthorsColumns[0], NewsItemAuthorsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "news_item_authors_news_item_id",
				Columns:    []*schema.Column{NewsItemAuthorsColumns[0]},
				RefColumns: []*schema.Column{NewsItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "news_item_authors_person_id",
				Columns:    []*schema.Column{NewsItemAuthorsColumns[1]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		NewsItemsTable,
		OrganizationsTable,
		PersonsTable,
		RssFeedsTable,
		NewsItemAuthorsTable,
	}
)

func init() {
	NewsItemsTable.ForeignKeys[0].RefTable = RssFeedsTable
	RssFeedsTable.ForeignKeys[0].RefTable = OrganizationsTable
	NewsItemAuthorsTable.ForeignKeys[0].RefTable = NewsItemsTable
	NewsItemAuthorsTable.ForeignKeys[1].RefTable = PersonsTable
}
