CREATE TABLE IF NOT EXISTS public.alerts (
  id CHAR(50) PRIMARY KEY,
  name VARCHAR(250) NULL,
  priority VARCHAR(50) NULL,
  source VARCHAR(250) NULL,
  message VARCHAR(250) NULL,
  integration_id CHAR(50) NULL,
  responder_ids CHAR(50)[] NULL,
  colleted_at TIMESTAMP NOT NULL DEFAULT now(),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE INDEX alerts_priority ON public.alerts (priority);
CREATE INDEX alerts_name ON public.alerts (name);
CREATE INDEX alerts_source ON public.alerts (source);
CREATE INDEX alerts_created_at ON public.alerts (created_at);
