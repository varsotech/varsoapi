-- create index "category_user_uuid_title" to table: "categories"
CREATE UNIQUE INDEX "category_user_uuid_title" ON "categories" ("user_uuid", "title");
