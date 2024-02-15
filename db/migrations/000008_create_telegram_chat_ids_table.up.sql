CREATE TABLE IF NOT EXISTS telegram_chat_ids (
    id TEXT UNIQUE NOT NULL,
    chat_id INTEGER NOT NULL,
    role_id text NOT NULL,
    CONSTRAINT fk_chat_id_role FOREIGN KEY (role_id) REFERENCES roles (id) ON UPDATE CASCADE
);