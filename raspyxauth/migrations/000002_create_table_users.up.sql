CREATE SCHEMA IF NOT EXISTS auth;

ALTER SCHEMA auth OWNER TO postgres;

COMMENT ON SCHEMA auth IS 'Schema for authentication/authorization';

CREATE TABLE IF NOT EXISTS auth.users (
    uuid VARCHAR(36) PRIMARY KEY,
    username VARCHAR(256) NOT NULL UNIQUE,
    password_hash VARCHAR(256) NOT NULL,
    access_level INT NOT NULL DEFAULT 0
);

ALTER TABLE auth.users OWNER TO postgres;
COMMENT ON TABLE auth.users IS 'Table with users';
GRANT ALL ON TABLE auth.users TO postgres;