-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (created_at, updated_at, email, username, password)
VALUES (NOW(), NOW(), 'statuspage@api.com', 'StatusPageAPI', 'b63729771dbc8166067f73c735c7a78dde61ccbe25d05d0cefec4ba2acbcfa23');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd