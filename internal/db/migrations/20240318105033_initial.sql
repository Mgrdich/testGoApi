-- Create "movies" table
CREATE TABLE "public"."movies" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "title" character varying(30) NULL, "director" character varying(30) NULL, "release_at" date NULL, "ticket_price" numeric(6,3) NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "person" table
CREATE TABLE "public"."person" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "first_name" character varying(30) NULL, "last_name" character varying(40) NULL, PRIMARY KEY ("id"));
