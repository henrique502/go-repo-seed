BEGIN;

DROP TABLE IF EXISTS alerts;

DROP UNIQUE INDEX alerts_priority;
DROP UNIQUE INDEX alerts_name;
DROP UNIQUE INDEX alerts_source;
DROP UNIQUE INDEX alerts_created_at;

COMMIT;
