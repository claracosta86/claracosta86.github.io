-- name: save-commentary
INSERT INTO commentaries (cultural_id, cultural_type, user_id, content, created_at)
VALUES (?, ?, ?, ?, NOW())

-- name: fetch-commentaries-by-cultural
SELECT
    c.id,
    c.cultural_id,
    c.cultural_type,
    c.user_id,
    c.content AS commentary,
    c.created_at,
    c.updated_at
FROM commentaries c
LEFT JOIN users u ON c.user_id = u.id