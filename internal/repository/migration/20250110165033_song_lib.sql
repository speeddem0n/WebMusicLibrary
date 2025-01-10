-- +goose Up
-- +goose StatementBegin
CREATE TABLE song_lib (
    id SERIAL PRIMARY KEY
    group VARCHAR(255) NOT NULL
    song VARCHAR(255) NOT NULL
    releaseDate TIMESTAMP NOT NULL
    text TEXT NOT NULL
    link text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE song_lib;
-- +goose StatementEnd
