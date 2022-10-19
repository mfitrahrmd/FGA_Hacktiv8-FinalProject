CREATE TABLE users(
    id TEXT PRIMARY KEY NOT NULL,
    username VARCHAR UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password TEXT NOT NULL,
    age INT,
    created_at DATE NOT NULL,
    updated_at DATE
);