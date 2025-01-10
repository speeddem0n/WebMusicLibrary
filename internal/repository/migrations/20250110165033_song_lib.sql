-- +goose Up
-- +goose StatementBegin
CREATE TABLE song_lib 
(
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    song VARCHAR(255) NOT NULL,
    release_date TIMESTAMP NOT NULL,
    text TEXT NOT NULL,
    link TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS song_lib;
-- +goose StatementEnd
