BEGIN;

CREATE TABLE games
(
    id           UUID PRIMARY KEY,
    title        VARCHAR(255) NOT NULL,
    description  TEXT         NOT NULL,
    genre        VARCHAR(30)  NOT NULL,
    release_year INTEGER      NOT NULL CHECK (release_year >= 1950 AND release_year <= 2100),
    created_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
    UNIQUE (title)
);

CREATE INDEX idx_games_genre ON games(genre);
CREATE INDEX idx_games_release_year ON games(release_year);

COMMIT;