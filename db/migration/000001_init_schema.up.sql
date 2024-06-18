CREATE TABLE "users" (
    "user_id" bigserial PRIMARY KEY,
    "username" varchar UNIQUE NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password_hash" text NOT NULL,
    "created_at" timestamptz DEFAULT (now())
);
CREATE TABLE "posts" (
    "post_id" bigserial PRIMARY KEY,
    "title" varchar NOT NULL,
    "content" text NOT NULL,
    "author_id" bigint NOT NULL,
    "published_at" timestamptz DEFAULT (now()),
    "updated_at" timestamptz
);
CREATE TABLE "comments" (
    "comment_id" bigserial PRIMARY KEY,
    "post_id" bigint NOT NULL,
    "author_id" bigint NOT NULL,
    "content" text NOT NULL,
    "created_at" timestamptz DEFAULT (now())
);
CREATE TABLE "likes" (
    "like_id" bigserial PRIMARY KEY,
    "post_id" bigint NOT NULL,
    "author_id" bigint NOT NULL,
    "created_at" timestamptz DEFAULT (now())
);
CREATE TABLE "categories" (
    "category_id" bigserial PRIMARY KEY,
    "name" varchar UNIQUE NOT NULL,
    "created_at" timestamptz DEFAULT (now())
);
CREATE TABLE "tags" (
    "tag_id" bigserial PRIMARY KEY,
    "name" varchar UNIQUE NOT NULL,
    "created_at" timestamptz DEFAULT (now())
);
CREATE TABLE "post_categories" ("post_id" bigint UNIQUE, "category_id" bigint UNIQUE);
CREATE TABLE "post_Tags" ("post_id" bigint UNIQUE, "tag_id" bigint UNIQUE);
CREATE INDEX ON "posts" ("author_id");
CREATE INDEX ON "comments" ("post_id");
CREATE INDEX ON "comments" ("author_id");
CREATE INDEX ON "likes" ("post_id");
CREATE INDEX ON "likes" ("author_id");
CREATE INDEX ON "post_categories" ("post_id");
CREATE INDEX ON "post_categories" ("category_id");
CREATE INDEX ON "post_Tags" ("post_id");
CREATE INDEX ON "post_Tags" ("tag_id");
ALTER TABLE
    "posts"
ADD
    FOREIGN KEY ("author_id") REFERENCES "users" ("user_id");
ALTER TABLE
    "comments"
ADD
    FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id");
ALTER TABLE
    "comments"
ADD
    FOREIGN KEY ("author_id") REFERENCES "users" ("user_id");
ALTER TABLE
    "likes"
ADD
    FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id");
ALTER TABLE
    "likes"
ADD
    FOREIGN KEY ("author_id") REFERENCES "users" ("user_id");
CREATE TABLE "posts_post_categories" (
        "posts_post_id" bigserial,
        "post_categories_post_id" bigint,
        PRIMARY KEY ("posts_post_id", "post_categories_post_id")
    );
ALTER TABLE
    "posts_post_categories"
ADD
    FOREIGN KEY ("posts_post_id") REFERENCES "posts" ("post_id");
ALTER TABLE
    "posts_post_categories"
ADD
    FOREIGN KEY ("post_categories_post_id") REFERENCES "post_categories" ("post_id");
ALTER TABLE
    "post_categories"
ADD
    FOREIGN KEY ("category_id") REFERENCES "categories" ("category_id");
CREATE TABLE "posts_post_Tags" (
        "posts_post_id" bigserial,
        "post_Tags_post_id" bigint,
        PRIMARY KEY ("posts_post_id", "post_Tags_post_id")
    );
ALTER TABLE
    "posts_post_Tags"
ADD
    FOREIGN KEY ("posts_post_id") REFERENCES "posts" ("post_id");
ALTER TABLE
    "posts_post_Tags"
ADD
    FOREIGN KEY ("post_Tags_post_id") REFERENCES "post_Tags" ("post_id");
ALTER TABLE
    "post_Tags"
ADD
    FOREIGN KEY ("tag_id") REFERENCES "tags" ("tag_id");