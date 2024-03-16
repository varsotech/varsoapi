-- create "categories" table
CREATE TABLE "categories" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "uuid" uuid NOT NULL, "user_uuid" uuid NOT NULL, "title" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "categories_uuid_key" to table: "categories"
CREATE UNIQUE INDEX "categories_uuid_key" ON "categories" ("uuid");
-- create index "category_user_uuid" to table: "categories"
CREATE INDEX "category_user_uuid" ON "categories" ("user_uuid");
-- create index "category_uuid" to table: "categories"
CREATE UNIQUE INDEX "category_uuid" ON "categories" ("uuid");
-- create "tasks" table
CREATE TABLE "tasks" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "uuid" uuid NOT NULL, "user_uuid" uuid NOT NULL, "title" character varying NOT NULL, "duration_seconds" integer NOT NULL, PRIMARY KEY ("id"));
-- create index "task_user_uuid" to table: "tasks"
CREATE INDEX "task_user_uuid" ON "tasks" ("user_uuid");
-- create index "task_uuid" to table: "tasks"
CREATE UNIQUE INDEX "task_uuid" ON "tasks" ("uuid");
-- create index "tasks_uuid_key" to table: "tasks"
CREATE UNIQUE INDEX "tasks_uuid_key" ON "tasks" ("uuid");