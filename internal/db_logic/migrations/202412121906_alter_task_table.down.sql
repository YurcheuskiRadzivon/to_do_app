ALTER TABLE "Task"
DROP CONSTRAINT "Task_user_id_fkey";

ALTER TABLE "Task"
ADD CONSTRAINT "Task_user_id_fkey"
FOREIGN KEY (user_id)
REFERENCES "User"(id);