CREATE TABLE IF NOT EXISTS publishers (
  id   UUID        NOT NULL DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (name)
);
