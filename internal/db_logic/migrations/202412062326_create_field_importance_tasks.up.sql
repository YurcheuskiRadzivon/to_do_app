ALTER TABLE "Task"
    ADD COLUMN importance INT DEFAULT 1;

UPDATE "Task"
SET importance = 1
WHERE importance IS NULL;

ALTER TABLE "Task"
    ALTER COLUMN importance SET NOT NULL;
