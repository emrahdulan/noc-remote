CREATE TABLE IF NOT EXISTS results(
   time TIMESTAMPTZ NOT NULL,
   task_id INT NOT NULL,
   data JSONB NOT NULL
);

SELECT create_hypertable('results', by_range('time'));

CREATE UNIQUE INDEX ix_taskid_time ON results (task_id, time DESC);
