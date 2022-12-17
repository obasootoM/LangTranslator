CREATE TABLE "client" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "second_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "language" varchar NOT NULL,
  "time" timestamp NOT NULL DEFAULT 'now()',
  "password" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);