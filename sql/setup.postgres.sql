DROP TABLE IF EXISTS mskit_events;

CREATE TABLE mskit_events (
  id SERIAL PRIMARY KEY,
  event_type VARCHAR NOT NULL,
  aggregate_type VARCHAR NOT NULL,
  aggregate_id VARCHAR NOT NULL,
  event_data TEXT NOT NULL DEFAULT ''
);
