CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    cultural_id INTEGER REFERENCES cultural(id),
    cultural_type ENUM('event', 'attraction') NOT NULL DEFAULT 'event',
    user_id INTEGER REFERENCES users(id),
    type ENUM('updated', 'canceled', 'commented') NOT NULL DEFAULT 'updated',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE notifications ADD COLUMN title VARCHAR(255);
ALTER TABLE notifications ADD COLUMN seen TINYINT(1) DEFAULT 0;

ALTER TABLE notifications MODIFY COLUMN type ENUM('updated', 'canceled', 'commented', 'closed') NOT NULL DEFAULT 'updated';

ALTER TABLE notifications DROP COLUMN title;

ALTER TABLE notifications MODIFY COLUMN cultural_type ENUM('event', 'tourist_attraction') NOT NULL DEFAULT 'event';
