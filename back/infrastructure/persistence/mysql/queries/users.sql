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
    type,
    name
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
    company_name = ?, 
WHERE id = ?;

-- name: update-user-password
UPDATE users
SET 
    password = ?
WHERE id = ?;

-- name: remove-from-favorites
DELETE FROM user_favorites WHERE user_id = ? AND favorite_id = ?;

-- name: delete-user-favorites-by-id
DELETE FROM user_favorites WHERE user_id = ?;

-- name: delete-user-by-id
DELETE FROM users WHERE id = ?;