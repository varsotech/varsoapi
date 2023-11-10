-- reverse: create index "post_uuid" to table: "posts"
DROP INDEX "post_uuid";
-- reverse: create index "post_total_votes" to table: "posts"
DROP INDEX "post_total_votes";
-- reverse: create index "post_author_user_uuid" to table: "posts"
DROP INDEX "post_author_user_uuid";
-- reverse: create "posts" table
DROP TABLE "posts";
-- reverse: create index "comment_uuid" to table: "comments"
DROP INDEX "comment_uuid";
-- reverse: create index "comment_user_uuid" to table: "comments"
DROP INDEX "comment_user_uuid";
-- reverse: create "comments" table
DROP TABLE "comments";
