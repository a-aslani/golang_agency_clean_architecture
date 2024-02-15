CREATE TABLE IF NOT EXISTS discovery_sessions
(
    id              TEXT UNIQUE  NOT NULL,
    name            VARCHAR(100) NOT NULL,
    email           VARCHAR(120) NOT NULL,
    project_details TEXT         NOT NULL,
    date            TIMESTAMP    NOT NULL,
    created         TIMESTAMP    NOT NULL,
    updated         TIMESTAMP    NULL,
    PRIMARY KEY (id)
);