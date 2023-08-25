CREATE TABLE "valid_slots" (
    id UUID PRIMARY KEY NOT NULL,
    weekday INTEGER REFERENCES weekdays (id) NOT NULL,
    slot INTEGER REFERENCES slots (id) NOT NULL
);

INSERT INTO "valid_slots" (id, weekday, slot) 
VALUES (uuid_generate_v4(), 2, 1), 
       (uuid_generate_v4(), 2, 2), 
       (uuid_generate_v4(), 2, 3), 
       (uuid_generate_v4(), 2, 6), 
       (uuid_generate_v4(), 3, 4), 
       (uuid_generate_v4(), 3, 6), 
       (uuid_generate_v4(), 4, 1), 
       (uuid_generate_v4(), 4, 7), 
       (uuid_generate_v4(), 5, 4), 
       (uuid_generate_v4(), 5, 6), 
       (uuid_generate_v4(), 5, 10), 
       (uuid_generate_v4(), 6, 2); 

-- VALUES (uuid_generate_v4(), 2, 1), -- terça-feira, 8:00 AM - 9:00 AM
--        (uuid_generate_v4(), 2, 2), -- terça-feira, 9:00 AM - 10:00 AM
--        (uuid_generate_v4(), 2, 3), -- terça-feira, 10:00 AM - 11:00 AM
--        (uuid_generate_v4(), 2, 6), -- terça-feira, 13:00 PM - 14:00 PM
--        (uuid_generate_v4(), 3, 4), -- quarta-feira, 11:00 AM - 12:00 PM
--        (uuid_generate_v4(), 3, 6), -- quarta-feira, 13:00 PM - 14:00 PM
--        (uuid_generate_v4(), 4, 1), -- quinta-feira, 8:00 AM - 9:00 AM
--        (uuid_generate_v4(), 4, 7), -- quarta-feira, 14:00 PM - 15:00 PM
--        (uuid_generate_v4(), 5, 4), -- quinta-feira, 11:00 AM - 12:00 AM
--        (uuid_generate_v4(), 5, 6), -- quinta-feira, 13:00 PM - 14:00 PM
--        (uuid_generate_v4(), 5, 10), -- quinta-feira, 17:00 PM - 18:00 PM
--        (uuid_generate_v4(), 6, 2); -- sexta-feira, 9:00 AM - 10:00 AM
       