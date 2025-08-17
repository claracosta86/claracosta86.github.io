CREATE TABLE user_favorites (
    id SERIAL PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    event_id BIGINT UNSIGNED NOT NULL,
    attraction_id BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (event_id) REFERENCES events(id),
    FOREIGN KEY (attraction_id) REFERENCES tourist_attractions(id)
);

ALTER TABLE user_favorites DROP FOREIGN KEY user_favorites_ibfk_1;
ALTER TABLE user_favorites DROP FOREIGN KEY user_favorites_ibfk_2;
ALTER TABLE user_favorites DROP FOREIGN KEY user_favorites_ibfk_3;

ALTER TABLE user_favorites DROP COLUMN event_id;
ALTER TABLE user_favorites DROP COLUMN attraction_id;
ALTER TABLE user_favorites ADD COLUMN favorite_type ENUM('event', 'attraction') NOT NULL DEFAULT 'event';
ALTER TABLE user_favorites ADD COLUMN favorite_id BIGINT UNSIGNED NOT NULL;
ALTER TABLE user_favorites ADD COLUMN cultural_id BIGINT UNSIGNED NOT NULL;

ALTER TABLE user_favorites DROP COLUMN cultural_id;