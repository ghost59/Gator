-- name: CreateFeed :one
INSERT INTO feeds(id,created_at,updated_at,name, url, user_id,last_fetched_at)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: Getfeed :many
SELECT name, url, user_id
FROM feeds;

-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1;  
-- name: MarkFeedFetched :one
UPDATE feeds
  set updated_at = NOW(),
    last_fetched_at = NOW()
WHERE feeds.id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY feeds.last_fetched_at ASC NULLS FIRST, feeds.updated_at ASC
LIMIT 1;
