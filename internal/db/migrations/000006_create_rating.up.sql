BEGIN;

CREATE TABLE ratings
(
    id         UUID PRIMARY KEY,
    user_id    UUID         NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    game_id    UUID         NOT NULL REFERENCES games (id) ON DELETE CASCADE,
    rating     INT          NOT NULL CHECK (rating >= 1 AND rating <= 5),
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, game_id)
);

CREATE INDEX idx_ratings_user_id ON ratings (user_id);
CREATE INDEX idx_ratings_game_id ON ratings (game_id);

COMMIT;
