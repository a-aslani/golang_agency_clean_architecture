CREATE TABLE IF NOT EXISTS contact_us_file
(
    contact_us_id TEXT        NOT NULL,
    file_id       TEXT        NOT NULL,
    PRIMARY KEY (contact_us_id, file_id),
    CONSTRAINT fk_contact_us_file_contact_us FOREIGN KEY (contact_us_id) REFERENCES contact_us (id) ON UPDATE CASCADE,
    CONSTRAINT fk_contact_us_file_file FOREIGN KEY (file_id) REFERENCES files (id) ON UPDATE CASCADE
);