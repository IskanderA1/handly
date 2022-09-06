ALTER TABLE "accounts" RENAME COLUMN "token" TO "uuid";

ALTER TABLE "accounts" DROP COLUMN "project_id";

CREATE TABLE "admins" (
  "username" varchar PRIMARY KEY,
  "password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);