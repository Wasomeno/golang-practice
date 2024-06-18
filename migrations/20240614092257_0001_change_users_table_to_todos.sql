-- +goose Up
-- +goose StatementBegin

CREATE TABLE todos (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE todos;
DROP TABLE users;

-- +goose StatementEnd
