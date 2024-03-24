-- reverse: modify "news_items" table
ALTER TABLE "news_items" DROP CONSTRAINT "news_items_rss_feeds_feed", DROP COLUMN "news_item_feed", ADD COLUMN "rss_feed_items" uuid NULL;
