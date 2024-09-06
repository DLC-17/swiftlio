-- +goose Up
CREATE TABLE IF NOT EXISTS one_time_codes (
    id UUID PRIMARY KEY,
    code INTEGER UNIQUE NOT NULL DEFAULT (
        FLOOR(RANDOM() * 900000 + 100000)::INTEGER
    ), 
    expires_at TIMESTAMP NOT NULL DEFAULT (NOW() + INTERVAL '5 minutes'),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE one_time_codes;
