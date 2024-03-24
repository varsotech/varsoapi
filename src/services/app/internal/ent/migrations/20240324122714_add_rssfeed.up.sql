-- modify "organizations" table
ALTER TABLE "organizations" DROP COLUMN "rss_feed_url";
-- create "rss_feeds" table
CREATE TABLE "rss_feeds" ("uuid" uuid NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "rss_feed_url" character varying NOT NULL, "organization_feeds" uuid NULL, PRIMARY KEY ("uuid"), CONSTRAINT "rss_feeds_organizations_feeds" FOREIGN KEY ("organization_feeds") REFERENCES "organizations" ("uuid") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create index "rss_feeds_rss_feed_url_key" to table: "rss_feeds"
CREATE UNIQUE INDEX "rss_feeds_rss_feed_url_key" ON "rss_feeds" ("rss_feed_url");
-- create index "rssfeed_rss_feed_url" to table: "rss_feeds"
CREATE UNIQUE INDEX "rssfeed_rss_feed_url" ON "rss_feeds" ("rss_feed_url");
-- create index "rssfeed_uuid" to table: "rss_feeds"
CREATE UNIQUE INDEX "rssfeed_uuid" ON "rss_feeds" ("uuid");
-- create "news_items" table
CREATE TABLE "news_items" ("uuid" uuid NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "rss_guid" character varying NOT NULL, "title" character varying NOT NULL, "description" character varying NOT NULL, "content" character varying NOT NULL, "link" character varying NOT NULL, "links" jsonb NOT NULL, "item_publish_time" timestamptz NULL, "item_update_time" timestamptz NULL, "image_url" character varying NOT NULL, "image_title" character varying NOT NULL, "categories" jsonb NOT NULL, "rss_feed_items" uuid NULL, PRIMARY KEY ("uuid"), CONSTRAINT "news_items_rss_feeds_items" FOREIGN KEY ("rss_feed_items") REFERENCES "rss_feeds" ("uuid") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create index "news_items_rss_guid_key" to table: "news_items"
CREATE UNIQUE INDEX "news_items_rss_guid_key" ON "news_items" ("rss_guid");
-- create index "newsitem_rss_guid" to table: "news_items"
CREATE UNIQUE INDEX "newsitem_rss_guid" ON "news_items" ("rss_guid");
-- create index "newsitem_uuid" to table: "news_items"
CREATE UNIQUE INDEX "newsitem_uuid" ON "news_items" ("uuid");
-- create "persons" table
CREATE TABLE "persons" ("uuid" uuid NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "email" character varying NOT NULL, "name" character varying NOT NULL, "news_item_authors" uuid NULL, PRIMARY KEY ("uuid"), CONSTRAINT "persons_news_items_authors" FOREIGN KEY ("news_item_authors") REFERENCES "news_items" ("uuid") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create index "person_email" to table: "persons"
CREATE UNIQUE INDEX "person_email" ON "persons" ("email");
-- create index "person_uuid" to table: "persons"
CREATE UNIQUE INDEX "person_uuid" ON "persons" ("uuid");
-- create index "persons_email_key" to table: "persons"
CREATE UNIQUE INDEX "persons_email_key" ON "persons" ("email");
