-- modify "persons" table
ALTER TABLE "persons" DROP COLUMN "news_item_authors";
-- create "news_item_authors" table
CREATE TABLE "news_item_authors" ("news_item_id" uuid NOT NULL, "person_id" uuid NOT NULL, PRIMARY KEY ("news_item_id", "person_id"), CONSTRAINT "news_item_authors_news_item_id" FOREIGN KEY ("news_item_id") REFERENCES "news_items" ("uuid") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "news_item_authors_person_id" FOREIGN KEY ("person_id") REFERENCES "persons" ("uuid") ON UPDATE NO ACTION ON DELETE CASCADE);
