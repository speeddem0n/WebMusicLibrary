-- +goose Up
-- +goose StatementBegin
CREATE TABLE song_lib (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    song_name VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    text TEXT NOT NULL,
    link VARCHAR(512)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table is exists song_lib;
-- +goose StatementEnd
