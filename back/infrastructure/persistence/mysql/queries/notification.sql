-- name: fetch-notifications-by-user-id
SELECT 
    uf.favorite_id,
    e.title
FROM user_favorites uf
INNER JOIN events e ON uf.favorite_id = e.id
WHERE uf.user_id = ? AND uf.favorite_type = "event" AND e.updated_at > uf.last_seen_at

UNION ALL

SELECT 
    uf.favorite_id,
    ta.title
FROM user_favorites uf
INNER JOIN tourist_attractions ta ON uf.favorite_id = ta.id
WHERE uf.user_id = ? AND uf.favorite_type = "tourist_attraction" AND ta.updated_at > uf.last_seen_at;