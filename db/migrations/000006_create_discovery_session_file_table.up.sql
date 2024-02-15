CREATE TABLE IF NOT EXISTS discovery_session_file
(
    discovery_session_id TEXT NOT NULL,
    file_id              TEXT NOT NULL,
    PRIMARY KEY (discovery_session_id, file_id),
    CONSTRAINT fk_discovery_session_file_discovery_session FOREIGN KEY (discovery_session_id) REFERENCES discovery_sessions (id) ON UPDATE CASCADE,
    CONSTRAINT fk_discovery_session_file_file FOREIGN KEY (file_id) REFERENCES files (id) ON UPDATE CASCADE
);