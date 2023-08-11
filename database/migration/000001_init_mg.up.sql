CREATE TABLE IF NOT EXISTS "users" (
    id SERIAL PRIMARY KEY,
    firstname varchar(255) NOT NULL,
    lastname varchar(255) NOT NULL,
    nickname varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS "posts" (
    id SERIAL PRIMARY KEY,
    title varchar(255) NOT NULL,
    picture varchar(255) NOT NULL,
    user_id int NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS "interests" (
    id SERIAL PRIMARY KEY,
    title varchar(255) NOT NULL,
    category varchar(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS "user_interests" (
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    interest_id int NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS "user_pictures" (
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    thumbnail varchar(255) NOT NULL,
    original varchar(255) NOT NULL,
    large varchar(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp
);