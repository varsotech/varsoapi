-- modify "news_items" table
ALTER TABLE "news_items" DROP COLUMN "rss_feed_items", ADD COLUMN "news_item_feed" uuid NULL, ADD CONSTRAINT "news_items_rss_feeds_feed" FOREIGN KEY ("news_item_feed") REFERENCES "rss_feeds" ("uuid") ON UPDATE NO ACTION ON DELETE SET NULL;
