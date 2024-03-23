-- reverse: create index "organization_unique_name" to table: "organizations"
DROP INDEX "organization_unique_name";
-- reverse: modify "organizations" table
ALTER TABLE "organizations" DROP COLUMN "website_url", DROP COLUMN "unique_name";
