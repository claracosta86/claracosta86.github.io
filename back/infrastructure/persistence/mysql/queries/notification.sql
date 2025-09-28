-- name: fetch-notifications-by-user-id
SELECT 
    n.id,
    n.title,
    n.cultural_type,
    n.cultural_id,
    n.type
FROM notifications n
WHERE n.user_id = ? 
AND n.cultural_id IN ( %s )
AND n.cultural_type = 'event'
AND n.seen = 0;


UNION ALL

SELECT 
    n.id,
    n.title,
    n.cultural_type,
    n.cultural_id,
    n.type
FROM notifications n
WHERE n.user_id = ? 
AND n.cultural_id IN ( %s )
AND n.cultural_type = 'tourist_attraction'
AND n.seen = 0;

-- name: mark-notifications-as-seen
UPDATE notifications
SET seen = 1
WHERE user_id = ?
AND id IN ( %s )