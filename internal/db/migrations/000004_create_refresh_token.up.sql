BEGIN;
CREATE TABLE refresh_tokens
(
    id           UUID PRIMARY KEY,
    user_id      UUID                NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    hashed_token VARCHAR(255) UNIQUE NOT NULL,
    is_active    BOOLEAN             NOT NULL DEFAULT true,
    expires_at   TIMESTAMP           NOT NULL,
    created_at   TIMESTAMP           NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens (user_id);
COMMIT;