-- name: CreateChannel :exec
INSERT INTO channels (title, description, link, last_updated)
VALUES (?, ?, ?, ?);

