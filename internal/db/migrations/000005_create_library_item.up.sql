BEGIN;

CREATE TABLE library_items
(
    id       UUID PRIMARY KEY,
    user_id  UUID         NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    game_id  UUID         NOT NULL REFERENCES games (id) ON DELETE CASCADE,
    status   VARCHAR(50)  NOT NULL,
    added_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, game_id)
);

CREATE INDEX idx_library_items_user_id ON library_items (user_id);
CREATE INDEX idx_library_items_game_id ON library_items (game_id);

COMMIT;
