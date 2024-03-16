-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "uuid" uuid NOT NULL, "email" character varying NULL, "username" character varying NULL, "password" character varying NULL, "salt" character varying NULL, "discord_user_id" character varying NULL, PRIMARY KEY ("id"));
-- create index "users_uuid_key" to table: "users"
CREATE UNIQUE INDEX "users_uuid_key" ON "users" ("uuid");
-- create index "user_uuid" to table: "users"
CREATE UNIQUE INDEX "user_uuid" ON "users" ("uuid");
-- create index "user_email" to table: "users"
CREATE UNIQUE INDEX "user_email" ON "users" ("email");
-- create index "user_username" to table: "users"
CREATE UNIQUE INDEX "user_username" ON "users" ("username");