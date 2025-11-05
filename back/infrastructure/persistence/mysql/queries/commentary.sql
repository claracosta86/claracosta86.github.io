-- name: save-commentary
INSERT INTO comments (cultural_id, cultural_type, user_id, content, created_at)
VALUES (?, ?, ?, ?, NOW())

-- name: fetch-commentaries-by-cultural
SELECT
    c.id,
    c.cultural_id,
    c.cultural_type,
    u.name AS user_name,
    c.content AS commentary,
    c.created_at,
    c.updated_at
FROM comments c
LEFT JOIN users u ON c.user_id = u.id
WHERE c.cultural_id = ? AND c.cultural_type = ?
ORDER BY c.created_at DESC;