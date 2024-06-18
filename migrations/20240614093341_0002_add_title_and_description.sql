-- +goose Up
-- +goose StatementBegin
DROP TABLE todos;

CREATE TABLE todos (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title Varchar(255) NOT NULL,
    description TEXT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todos;
-- +goose StatementEnd
