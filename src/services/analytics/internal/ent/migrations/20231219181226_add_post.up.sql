-- create "comments" table
CREATE TABLE "comments" ("uuid" uuid NOT NULL, "user_uuid" uuid NOT NULL, "text" character varying NOT NULL, "was_edited" boolean NOT NULL DEFAULT false, "total_votes" bigint NOT NULL, "upvotes" bigint NOT NULL, "downvotes" bigint NOT NULL, PRIMARY KEY ("uuid"));
-- create index "comment_user_uuid" to table: "comments"
CREATE INDEX "comment_user_uuid" ON "comments" ("user_uuid");
-- create index "comment_uuid" to table: "comments"
CREATE UNIQUE INDEX "comment_uuid" ON "comments" ("uuid");
-- create "posts" table
CREATE TABLE "posts" ("uuid" uuid NOT NULL, "author_user_uuid" uuid NOT NULL, "title" character varying NOT NULL, "cover_image" character varying NOT NULL, "total_votes" bigint NOT NULL, "upvotes" bigint NOT NULL, "downvotes" bigint NOT NULL, "post_comments" uuid NULL, PRIMARY KEY ("uuid"), CONSTRAINT "posts_comments_comments" FOREIGN KEY ("post_comments") REFERENCES "comments" ("uuid") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create index "post_author_user_uuid" to table: "posts"
CREATE INDEX "post_author_user_uuid" ON "posts" ("author_user_uuid");
-- create index "post_total_votes" to table: "posts"
CREATE INDEX "post_total_votes" ON "posts" ("total_votes");
-- create index "post_uuid" to table: "posts"
CREATE UNIQUE INDEX "post_uuid" ON "posts" ("uuid");
