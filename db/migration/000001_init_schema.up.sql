CREATE TABLE "client" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "second_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);
CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "email" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("email") REFERENCES "client" ("email");

CREATE TABLE "profile" (
  "id" bigserial PRIMARY KEY, 
  "image" varchar NOT NULL,
  "name" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "address_line" varchar NOT NULL,
  "country" varchar NOT NULL,
  "native_language" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
); 

ALTER TABLE "profile" ADD FOREIGN KEY ("email") REFERENCES "client" ("email");

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY, 
  "source_language" varchar NOT NULL, 
  "target_language" varchar NOT NULL,  
  "translator" varchar NOT NULL, 
  "proof_reader" varchar NOT NULL, 
  "translation_delivary_date" varchar NOT NULL, 
  "proof_reading_delivary_date" varchar NOT NULL,  
  "project_end_date" varchar NOT NULL,  
  "service_level" varchar NOT NULL,  
  "profession" varchar NOT NULL,  
  "translator_category" varchar NOT NULL,  
  "delivary_speed" varchar NOT NULL, 
  "translator_request" varchar NOT NULL,  
  "delivary_address" varchar NOT NULL
);

CREATE TABLE "pages" (
  "id" bigserial PRIMARY KEY,
  "source_language" varchar NOT NULL, 
  "target_language" varchar NOT NULL, 
  "file" varchar NOT NULL,
  "profession" varchar NOT NULL,
  "category" varchar NOT NULL,
  "field" varchar NOT NULL,
  "duration" varchar NOT NULL,
  "additional_service" varchar NOT NULL
);