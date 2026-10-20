

-- name: CreateFeedFollow :one

WITH inserted_feedfollows AS (
    INSERT INTO feedfollows (id, created_at, updated_at,user_id, feed_id)
    
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
)
    RETURNING *
)

SELECT
    inserted_feedfollows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feedfollows
INNER JOIN feeds ON inserted_feedfollows.feed_id = feeds.id
INNER JOIN users ON inserted_feedfollows.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT feedfollows.*,
        feeds.name AS feed_name,
        users.name AS user_name
FROM feedfollows
INNER JOIN feeds ON feeds.id = feedfollows.feed_id
INNER JOIN users ON users.id = feedfollows.user_id
WHERE  feedfollows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feedfollows
WHERE feedfollows.feed_id = $1  AND feedfollows.user_id = $2;

