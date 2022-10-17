CREATE TABLE IF NOT EXISTS "sessions" (
"id" uuid PRIMARY KEY,
"user_name" varchar(20) UNIQUE NOT NULL,
"refresh_token" varchar NOT NULL,
"user_agent" varchar NOT NULL,
"client_ip" varchar NOT NULL,
"is_blocked" boolean NOT NULL DEFAULT false,
"expires_at" timestamptz NOT NULL DEFAULT (now()),
"created_at" timestamptz NOT NULL DEFAULT (now())
);
