-- name: CreateChannel :exec
INSERT INTO channels (title, description, link, last_updated)
VALUES (?, ?, ?, ?);

-- name: GetChannels :many
SELECT * FROM channels;

-- name: DeleteChannel :exec
DELETE FROM channels
WHERE title = ?;

-- name: TitleGetChannelID :one
SELECT id FROM channels
WHERE title = ?;

-- name: CreatePost :exec
INSERT INTO posts (channel_id, title, description, link, pub_date)
VALUES (?, ?, ?, ?, ?);

-- name: ChannelGetPosts :many
SELECT * FROM posts
WHERE channel_id = ?;
