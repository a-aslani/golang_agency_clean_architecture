CREATE TABLE IF NOT EXISTS roles (
    id TEXT UNIQUE NOT NULL,
    code VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL
);

INSERT INTO roles(id, code, name) VALUES ('2a8fd6f2-33c9-4254-a818-ee15cb640ede', 'CEO', 'CEO');
INSERT INTO roles(id, code, name) VALUES ('6459d91f-cc79-4609-a7a3-c59fe2839615', 'CTO', 'CTO');
INSERT INTO roles(id, code, name) VALUES ('73f7199d-a872-4f7f-81b6-f77192375244', 'PM', 'Project Manager');
INSERT INTO roles(id, code, name) VALUES ('63d332bf-3402-4c9f-affc-0042af05ab9a', 'MARKETING', 'Marketing');
INSERT INTO roles(id, code, name) VALUES ('4a6369e9-b7c8-435b-b54a-4fe11c8832ad', 'DEVELOPER', 'Developer');
INSERT INTO roles(id, code, name) VALUES ('034910b8-e14c-4f3e-952a-1d76e2082867', 'DESIGNER', 'UI/UX Designer');
INSERT INTO roles(id, code, name) VALUES ('f496a25f-f2c5-4d2d-aa6c-3f30d421ce51', 'SUPPORT', 'Support');