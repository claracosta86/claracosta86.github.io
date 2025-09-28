-- name: register-user
INSERT INTO users (name, email, document, company_name, type, password, document_type) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: fetch-user-by-id
SELECT 
    id, 
    name, 
    email, 
    document, 
    company_name, 
    type 
FROM users WHERE id = ?;

-- name: fetch-user-by-email
SELECT 
    id,
    type
FROM users 
WHERE email = ?;

-- name: fetch-user-password-by-id
SELECT 
    password 
FROM users WHERE id = ?;

-- name: update-user-profile
UPDATE users
SET 
    name = ?, 
    email = ?, 
    company_name = ? 
WHERE id = ?;

-- name: update-user-password
UPDATE users
SET 
    password = ?
WHERE id = ?;

-- name: delete-user-favorites-by-id
DELETE FROM user_favorites WHERE user_id = ? AND user_id != NULL;

-- name: delete-user-by-id
DELETE FROM users WHERE id = ?;

-- name: delete-users-favorites-by-event-id
DELETE FROM user_favorites WHERE favorite_type = 'event' AND favorite_id IN (%s);

-- name: delete-users-favorites-by-tourist-attraction-id
DELETE FROM user_favorites WHERE favorite_type = 'tourist_attraction' AND favorite_id IN (%s);

-- name: add-user-favorite
INSERT INTO user_favorites (user_id, favorite_type, favorite_id, favorited_at, last_seen_at) VALUES (?, ?, ?, NOW(), NOW());
ON DUPLICATE KEY UPDATE last_seen_at = NOW();

-- name: remove-user-favorite
DELETE FROM user_favorites WHERE user_id = ? AND favorite_type = ? AND favorite_id = ?;

-- name: fetch-user-favorites-by-id
SELECT 
    uf.favorite_id,
    e.title,
    uf.favorite_type
FROM user_favorites uf
INNER JOIN events e ON uf.favorite_id = e.id
WHERE uf.user_id = ? AND uf.favorite_type = "event"

UNION ALL

SELECT 
    uf.favorite_id,
    ta.title,
    uf.favorite_type
FROM user_favorites uf
INNER JOIN tourist_attractions ta ON uf.favorite_id = ta.id
WHERE uf.user_id = ? AND uf.favorite_type = "tourist_attraction";
