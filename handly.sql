CREATE TYPE "event_type" AS ENUM (
  'info',
  'error',
  'success',
  'undefind'
);

CREATE TABLE "projects" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "events" (
  "id" bigserial PRIMARY KEY,
  "project_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "event_type" event_type NOT NULL DEFAULT 'event_type.info'
);

CREATE TABLE "logs" (
  "id" bigserial PRIMARY KEY,
  "project_id" bigint NOT NULL,
  "event_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "data" varchar
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "project_id" bigint NOT NULL,
  "uuid" varchar NOT NULL,
  "name" varchar,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "projects" ("name");

CREATE INDEX ON "events" ("name");

CREATE INDEX ON "events" ("id", "project_id");

CREATE INDEX ON "logs" ("project_id", "event_id", "user_id");

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "users" ("uuid");

CREATE INDEX ON "users" ("id", "project_id");

ALTER TABLE "events" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");
