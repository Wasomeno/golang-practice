-- +goose Up

CREATE TABLE todos (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);

-- +goose Down

DROP TABLE todos;