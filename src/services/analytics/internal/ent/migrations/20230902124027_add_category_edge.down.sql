-- reverse: modify "tasks" table
ALTER TABLE "tasks" DROP CONSTRAINT "tasks_categories_category", DROP COLUMN "task_category";
