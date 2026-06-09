CREATE SCHEMA IF NOT EXISTS auth;

ALTER SCHEMA auth OWNER TO postgres;

COMMENT ON SCHEMA auth IS 'Schema for authentication/authorization';

CREATE TABLE IF NOT EXISTS auth.refresh_sessions (
    id SERIAL PRIMARY KEY,
    user_uuid VARCHAR(36) REFERENCES auth.users(uuid) ON DELETE CASCADE,
    refresh_token VARCHAR(36) NOT NULL,
    user_agent VARCHAR(256) NOT NULL,
    fingerprint VARCHAR(256) NOT NULL,
    ip VARCHAR(15) NOT NULL,
    expires_in BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

ALTER TABLE auth.refresh_sessions OWNER TO postgres;
COMMENT ON TABLE auth.refresh_sessions IS 'Table with refresh tokens';
GRANT ALL ON TABLE auth.refresh_sessions TO postgres;