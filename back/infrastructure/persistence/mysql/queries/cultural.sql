-- name: create-event
INSERT INTO events (title, description, location, start_date, end_date, duration,
price, is_accessible, organizer_id, image, created_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW())

-- name: create-tourist-attraction
INSERT INTO tourist_attractions (title, description, location, open_days, open_time, price, is_accessible, organizer_id, image, created_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW())


-- name: fetch-event-by-id
SELECT 
    e.id,  
    IFNULL(e.title, ''), 
    IFNULL(e.description, ''), 
    IFNULL(e.location, ''),
    IFNULL(e.start_date, ''), 
    IFNULL(e.end_date, ''), 
    IFNULL(e.duration, ''), 
    e.price, 
    e.is_accessible, 
    IFNULL(e.organizer_id, 0), 
    IFNULL(u.email, ''), 
    IFNULL(e.image, '')
FROM events e
LEFT JOIN users u ON e.organizer_id = u.id
WHERE e.id = ?

-- name: fetch-tourist-attraction-by-id
SELECT 
    ta.id, 
    IFNULL(ta.title, ''),
    IFNULL(ta.description, ''), 
    IFNULL(ta.location, ''), 
    IFNULL(ta.open_days, ''), 
    IFNULL(ta.open_time, ''), 
    ta.price, 
    ta.is_accessible, 
    IFNULL(ta.organizer_id, 0), 
    IFNULL(u.email, ''), 
    IFNULL(ta.image, '')
FROM tourist_attractions ta
LEFT JOIN users u ON ta.organizer_id = u.id
WHERE ta.id = ?

-- name: update-event
UPDATE events
SET title = ?, description = ?, location = ?, start_date = ?, end_date = ?, 
price = ?, is_accessible = ?, organizer_id = ?, image = ?,
updated_at = NOW()
WHERE id = ?

-- name: update-tourist-attraction
UPDATE tourist_attractions
SET title = ?, description = ?, location = ?, open_days = ?, open_time = ?,
price = ?, is_accessible = ?, organizer_id = ?, image = ?,
updated_at = NOW()
WHERE id = ?

-- name: delete-event
DELETE FROM events
WHERE id = ?

-- name: delete-tourist-attraction
DELETE FROM tourist_attractions
WHERE id = ?

-- name: fetch-events-ids-by-organizer
SELECT id FROM events WHERE organizer_id = ?;

-- name: fetch-tourist-attractions-ids-by-organizer
SELECT id FROM tourist_attractions WHERE organizer_id = ?;