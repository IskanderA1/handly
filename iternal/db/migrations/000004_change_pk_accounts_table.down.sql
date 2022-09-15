DROP TABLE IF EXISTS "logs";

DROP TABLE IF EXISTS "users";

CREATE TABLE "accounts" (
  "account_id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "uuid" varchar UNIQUE NOT NULL,
  "last_update_at" timestamp NOT NULL DEFAULT 'now()',
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "logs" (
  "id" bigserial PRIMARY KEY,
  "project_id" bigint NOT NULL,
  "event_id" bigint NOT NULL,
  "account_id" varchar NOT NULL,
  "data" varchar,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE "logs" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("account_id");