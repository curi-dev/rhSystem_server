CREATE TABLE slots (
    id SERIAL PRIMARY KEY,
    label VARCHAR(50),
    value INTEGER NOT NULL
);

-- ALTER TABLE slots
-- ADD CONSTRAINT unique_slots_value UNIQUE (value);

INSERT INTO slots (label, value) 
VALUES ('8:00 AM - 9:00 AM', 1),
       ('9:00 AM - 10:00 AM', 2),
       ('10:00 AM - 11:00 AM', 3),
       ('11:00 AM - 12:00 PM', 4),
       ('12:00 PM - 13:00 PM', 5),
       ('13:00 PM - 14:00 PM', 6),
       ('14:00 PM - 15:00 PM', 7),
       ('15:00 PM - 16:00 PM', 8),
       ('16:00 PM - 17:00 PM', 9),
       ('17:00 PM - 18:00 PM', 10);
       