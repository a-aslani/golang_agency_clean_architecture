CREATE TABLE IF NOT EXISTS subscribers(
    id TEXT NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    created TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);