CREATE TABLE IF NOT EXISTS alerts (
  id uuid PRIMARY KEY,
  name VARCHAR(250) NULL,
  priority VARCHAR(50) NULL,
  source VARCHAR(250) NULL,
  message VARCHAR(250) NULL,
  integration_id UUID NULL,
  responder_ids UUID[] NULL,
  colleted_at TIMESTAMP NOT NULL DEFAULT now(),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX alerts_priority ON alerts (priority);
CREATE UNIQUE INDEX alerts_name ON alerts (name);
CREATE UNIQUE INDEX alerts_source ON alerts (source);
CREATE UNIQUE INDEX alerts_created_at ON alerts (created_at);
