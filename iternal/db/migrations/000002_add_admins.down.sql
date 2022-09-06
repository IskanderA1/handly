ALTER TABLE "accounts" RENAME COLUMN "uuid" TO "token";

ALTER TABLE "accounts" ADD COLUMN "project_id";

DROP TABLE IF EXISTS "admins";
