-- name: CreateFeedFollow :one
WITH insert_feed_follows AS (
    INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5)

    RETURNING *
)

SELECT     
    insert_feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM insert_feed_follows
JOIN users ON users.id = insert_feed_follows.user_id
JOIN feeds ON feeds.id = insert_feed_follows.feed_id;




-- name: GetFeedsThatUserFollows :many
SELECT 
    feed_follows.*,
    users.name AS user_name,
    feeds.name AS feed_name
FROM feed_follows
JOIN users ON users.id = feed_follows.user_id
JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;