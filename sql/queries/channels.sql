-- name: CreateChannel :exec
INSERT INTO channels (title, description, link, last_updated)
VALUES (?, ?, ?, ?);

-- name: ViewChannel :many
SELECT title FROM channels;

-- name: DeleteChannel :exec
DELETE FROM channels
WHERE title = ?;

