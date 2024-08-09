create type book_access as enum ('public', 'private');

CREATE TABLE users
(
    user_id  SERIAL PRIMARY KEY,
    username VARCHAR(20) UNIQUE NOT NULL,
    password TEXT               NOT NULL,
    email    TEXT               NOT NULL
);

CREATE TABLE books
(
    book_id     SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    author      VARCHAR(255) NOT NULL,
    issue_year  INTEGER      NOT NULL,
    user_id     INTEGER      NOT NULL,
    file_name   TEXT         NOT NULL,
    book_access BOOK_ACCESS  NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);