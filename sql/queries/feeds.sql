-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetFeedByUser :many
SELECT *
FROM feeds
WHERE user_id = $1;

-- name: GetFeeds :many
SELECT *
FROM feeds LIMIT 10;

-- name: GetNextFeedToFetch :many
SELECT *
FROM feeds
order by last_time_fetched_at asc nulls first LIMIT $1;

-- name: MarkFeedAtFetched :one
UPDATE feeds
SET last_time_fetched_at = NOW(),
    updated_at           = NOW()
WHERE id = $1 RETURNING *;
