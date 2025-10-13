-- name: fetch-notifications-by-user-id
SELECT
    n.id,
    COALESCE(e.title, ta.title) AS title,
    n.cultural_type,
    n.cultural_id,
    n.type
FROM
    notifications n
LEFT JOIN
    events e ON n.cultural_id = e.id AND n.cultural_type = 'event'
LEFT JOIN
    tourist_attractions ta ON n.cultural_id = ta.id AND n.cultural_type = 'tourist_attraction'
WHERE
    n.user_id = ?
    AND n.seen = 0
    AND n.type IN ('commented', 'canceled', 'updated') 

UNION ALL

SELECT 
    0 AS id, 
    COALESCE(e.title, ta.title) AS title,
    uf.favorite_type AS cultural_type,
    uf.favorite_id AS cultural_id,
    'updated' AS type
FROM 
    user_favorites uf
LEFT JOIN 
    events e ON uf.favorite_id = e.id AND uf.favorite_type = 'event'
LEFT JOIN 
    tourist_attractions ta ON uf.favorite_id = ta.id AND uf.favorite_type = 'tourist_attraction'
LEFT JOIN 
    notifications n ON uf.favorite_id = n.cultural_id 
                    AND uf.favorite_type = n.cultural_type 
                    AND n.user_id = uf.user_id 
                    AND n.seen = 0 
                    AND n.type = 'updated'
WHERE 
    uf.user_id = ?
    AND (uf.last_seen_at IS NULL OR COALESCE(e.updated_at, ta.updated_at) > uf.last_seen_at)
    AND n.id IS NULL;

-- name: mark-notifications-as-seen
UPDATE notifications
SET seen = 1
WHERE user_id = ?
AND id IN ( %s )

-- name: create-notification
INSERT INTO notifications (user_id, cultural_id, cultural_type, type, seen) VALUES (?, ?, ?, ?, ?);