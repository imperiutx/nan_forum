CREATE TABLE IF NOT EXISTS "users" (
                         "id" SERIAL PRIMARY KEY,
                         "user_name" varchar(20) UNIQUE NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "hashed_password" varchar NOT NULL,
                         "is_banned" boolean NOT NULL DEFAULT false,
                         "is_admin" boolean NOT NULL DEFAULT false,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "categories" (
                              "id" SERIAL PRIMARY KEY,
                              "name" varchar NOT NULL,
                              "created_by" varchar NOT NULL,
                              "is_visible" boolean NOT NULL DEFAULT true,
                              "created_at" timestamptz NOT NULL DEFAULT (now()),
                              "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "topics" (
                          "id" SERIAL PRIMARY KEY,
                          "category_id" bigint NOT NULL,
                          "title" varchar(100) NOT NULL,
                          "body" text NOT NULL,
                          "created_by" varchar NOT NULL,
                          "points" bigint NOT NULL DEFAULT 1,
                          "is_visible" boolean NOT NULL DEFAULT true,
                          "created_at" timestamptz NOT NULL DEFAULT (now()),
                          "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "comments" (
                            "id" SERIAL PRIMARY KEY,
                            "topic_id" bigint NOT NULL,
                            "body" text NOT NULL,
                            "points" bigint NOT NULL DEFAULT 1,
                            "created_by" varchar NOT NULL,
                            "is_visible" boolean NOT NULL DEFAULT true,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "replies" (
                           "id" SERIAL PRIMARY KEY,
                           "comment_id" bigint NOT NULL,
                           "body" text NOT NULL,
                           "points" bigint NOT NULL DEFAULT 1,
                           "created_by" varchar NOT NULL,
                           "is_visible" boolean NOT NULL DEFAULT true,
                           "created_at" timestamptz NOT NULL DEFAULT (now()),
                           "updated_at" timestamptz NOT NULL DEFAULT (now())
);

-- ALTER TABLE "users" ADD FOREIGN KEY ("user_name") REFERENCES "categories" ("created_by");
--
-- ALTER TABLE "users" ADD FOREIGN KEY ("user_name") REFERENCES "topics" ("created_by");
--
-- ALTER TABLE "categories" ADD FOREIGN KEY ("id") REFERENCES "topics" ("category_id");
--
-- ALTER TABLE "users" ADD FOREIGN KEY ("user_name") REFERENCES "comments" ("created_by");
--
-- ALTER TABLE "topics" ADD FOREIGN KEY ("id") REFERENCES "comments" ("topic_id");
--
-- ALTER TABLE "users" ADD FOREIGN KEY ("user_name") REFERENCES "replies" ("created_by");
--
-- ALTER TABLE "comments" ADD FOREIGN KEY ("id") REFERENCES "replies" ("comment_id");
