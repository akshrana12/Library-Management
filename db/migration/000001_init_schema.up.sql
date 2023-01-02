CREATE TABLE
    IF NOT EXISTS "books" (
        "book_id" INTEGER PRIMARY KEY,
        "book_name" VARCHAR,
        "book_author" VARCHAR,
        "book_genre" VARCHAR
    );

CREATE TABLE
    IF NOT EXISTS "users" (
        "student_id" INTEGER PRIMARY KEY,
        "student_name" VARCHAR,
        "email_id" VARCHAR UNIQUE
    );

CREATE TABLE
    IF NOT EXISTS "book_issue" (
        "user_id" INTEGER,
        "book_id" INTEGER
    );