CREATE TABLE "client" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "second_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone_number" varchar NOT NULL,
  "language" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "time" varchar NOT NULL,
  "password" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);
CREATE TABLE "translator" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "second_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);


