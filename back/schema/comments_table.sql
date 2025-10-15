CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    cultural_id INTEGER REFERENCES cultural(id),
    cultural_type ENUM('event', 'attraction') NOT NULL DEFAULT 'event',
    user_id INTEGER REFERENCES users(id),
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE comments ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;