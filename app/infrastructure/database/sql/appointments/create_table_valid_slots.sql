CREATE TABLE "valid_slots" (
    id UUID PRIMARY KEY NOT NULL,
    weekday INTEGER REFERENCES weekdays (id) NOT NULL,
    slot INTEGER REFERENCES slots (id) NOT NULL
);

INSERT INTO "valid_slots" (id, weekday, slot) 
VALUES ('43a21f87-5477-41ef-b473-58814c6acffd', 2, 1), 
       ('b54e6710-e0f6-43bf-9847-aeb246ba4829', 2, 2), 
       ('375ef686-7daf-42f7-b04a-d29717b4c0e1', 2, 3), 
       ('a3e6978f-f5cd-42ff-bd7e-35050e0d7cb0', 2, 6), 
       ('66cb57cd-f76f-402a-862a-6043cf1598fd', 3, 4), 
       ('002c6e43-156e-4905-9dfe-3213d39bfb59', 3, 6), 
       ('c42b7ad2-bfac-4115-b7b0-1f594fd7eb75', 4, 1), 
       ('2a9cbbc1-ebd3-45e6-8725-9ab23635f744', 4, 7), 
       ('a3bd9749-03aa-4ab0-bec6-e4c76dd8ebff', 5, 4), 
       ('a03e33dc-aafe-4cb2-9747-598dc46e19b8', 5, 6), 
       ('9fa18e04-b548-4ca3-9579-09b24906fd0a', 5, 10), 
       ('a328159a-677a-41d3-8acd-90f652da1e45', 6, 2); 

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
       