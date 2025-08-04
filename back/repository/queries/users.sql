-- name: register-user
INSERT INTO users (name, email, document, company_name, type, password) VALUES (?, ?, ?, ?, ?, ?);

-- name: fetch-user-id-by-email
SELECT id FROM users WHERE email = ?;
