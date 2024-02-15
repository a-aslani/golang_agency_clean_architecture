CREATE TABLE IF NOT EXISTS contact_us
(
    id         TEXT UNIQUE  NOT NULL,
    name       VARCHAR(100) NOT NULL,
    email      VARCHAR(120) NOT NULL,
    message    TEXT         NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    PRIMARY KEY (id)
);