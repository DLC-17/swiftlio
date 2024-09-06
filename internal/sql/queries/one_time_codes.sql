-- name: CreateOneTimeCode :one
INSERT INTO one_time_codes (id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetOneTimeCode: one
SELECT code FROM one_time_codes
WHERE code = $1 and user_id = $2;
