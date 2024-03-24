-- reverse: create "news_item_authors" table
DROP TABLE "news_item_authors";
-- reverse: modify "persons" table
ALTER TABLE "persons" ADD COLUMN "news_item_authors" uuid NULL;
