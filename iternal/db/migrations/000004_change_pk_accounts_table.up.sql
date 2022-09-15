DROP TABLE IF EXISTS "logs";

DROP TABLE IF EXISTS "accounts";

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "project_account_id" varchar,
  "name" varchar,
  "uuid" varchar,
  "last_update_at" timestamp NOT NULL DEFAULT 'now()',
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "logs" (
  "id" bigserial PRIMARY KEY,
  "project_id" bigint NOT NULL,
  "event_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "data" varchar,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE "logs" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
