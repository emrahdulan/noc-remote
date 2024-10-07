DROP trigger task_updated_at_trigger on tasks;
DROP function task_updated_at();

DROP INDEX IF EXISTS tasks_id;
DROP INDEX IF EXISTS tasks_cid;
DROP INDEX IF EXISTS tasks_is_active;

DROP TABLE IF EXISTS tasks;