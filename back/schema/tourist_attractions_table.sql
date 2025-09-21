CREATE TABLE tourist_attractions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    location VARCHAR(255),
    description TEXT,
    open_days VARCHAR(255) NOT NULL,
    open_time VARCHAR(255) NOT NULL,
    is_accessible BOOLEAN DEFAULT TRUE,
    price DECIMAL(10, 2) DEFAULT 0.00,
    image VARCHAR(255),
    organizer_id INT,
    organizer_contact VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE tourist_attractions MODIFY COLUMN price VARCHAR(255);

ALTER TABLE tourist_attractions MODIFY COLUMN id SERIAL;

ALTER TABLE tourist_attractions ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE tourist_attractions ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;

ALTER TABLE tourist_attractions MODIFY COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL;
ALTER TABLE tourist_attractions MODIFY COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL;

ALTER TABLE tourist_attractions DROP COLUMN organizer_contact;