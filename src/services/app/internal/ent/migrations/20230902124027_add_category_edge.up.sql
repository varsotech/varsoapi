-- modify "tasks" table
ALTER TABLE "tasks" ADD COLUMN "task_category" bigint NULL, ADD CONSTRAINT "tasks_categories_category" FOREIGN KEY ("task_category") REFERENCES "categories" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
