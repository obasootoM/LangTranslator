CREATE TABLE "client" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "second_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone_number" int NOT NULL,
  "language" varchar NOT NULL,
  "time" timestamp NOT NULL DEFAULT 'now()',
  "password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL,
  "updated_at" timestamptz DEFAULT 'now()',
  "created_at" timestamptz DEFAULT 'now()'
);