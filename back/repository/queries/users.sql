-- name: register-user
INSERT INTO users (name, email, document, company_name, type, password, document_type) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: fetch-user-id-by-email
SELECT 
    id 
FROM users 
WHERE email = ?;

-- name: fetch-user-password-by-id
SELECT 
    password 
FROM users WHERE id = ?;

-- name: fetch-user-by-id
SELECT 
    id, 
    name, 
    email, 
    document, 
    company_name, 
    type 
FROM users WHERE id = ?;

-- name: update-user-profile
UPDATE users
SET 
    name = ?, 
    email = ?, 
    company_name = ?, 
WHERE id = ?;

-- name: delete-user-favorites-by-id
DELETE FROM user_favorites WHERE user_id = ?;

-- name: delete-user-by-id
DELETE FROM users WHERE id = ?;