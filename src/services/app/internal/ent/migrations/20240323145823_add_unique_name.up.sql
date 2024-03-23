-- modify "organizations" table
ALTER TABLE "organizations" ADD COLUMN "unique_name" character varying NULL, ADD COLUMN "website_url" character varying NOT NULL;
-- create index "organization_unique_name" to table: "organizations"
CREATE UNIQUE INDEX "organization_unique_name" ON "organizations" ("unique_name");
