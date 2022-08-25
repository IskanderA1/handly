CREATE TYPE "event_type" AS ENUM (
  'info',
  'error',
  'success',
  'undefind'
);

CREATE TABLE "projects" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "token" varchar UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "events" (
  "id" bigserial PRIMARY KEY,
  "project_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "event_type" event_type NOT NULL DEFAULT 'info'
);

CREATE TABLE "logs" (
  "id" bigserial PRIMARY KEY,
  "project_id" bigint NOT NULL,
  "event_id" bigint NOT NULL,
  "account_id" varchar NOT NULL,
  "data" varchar,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "accounts" (
  "account_id" varchar PRIMARY KEY,
  "project_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "token" varchar UNIQUE NOT NULL,
  "last_update_at" timestamp NOT NULL DEFAULT 'now()',
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "projects" ("name");

CREATE INDEX ON "events" ("name");

CREATE INDEX ON "events" ("id", "project_id");

CREATE INDEX ON "logs" ("project_id", "event_id", "account_id");

CREATE INDEX ON "accounts" ("name");

CREATE INDEX ON "accounts" ("account_id", "project_id");

ALTER TABLE "events" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");
