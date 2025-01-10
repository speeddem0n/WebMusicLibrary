-- +goose Up
-- +goose StatementBegin
CREATE TABLE song_verse{
    id SERIAL PRIMARY KEY
    song_id int REFERENCES song_lib (id) ON DELETE CASCADE NOT NULL
    verse TEXT NOT NULL
}
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE song_verse
-- +goose StatementEnd
