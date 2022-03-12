create table users(
    "id" serial primary key,
    "name" varchar,
    "email" varchar,
    "login" varchar,
    "password" varchar,
    "created_at" timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp
);