-- name: register-user
INSERT INTO users (name, email, document, company_name, type, password, document_type) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: fetch-user-id-by-email
SELECT id FROM users WHERE email = ?;

-- name: fetch-user-password-by-id
SELECT password FROM users WHERE id = ?;