CREATE TABLE "appointment_status" (
    id SERIAL PRIMARY KEY,
    status VARCHAR(50) NOT NULL,
    value INTEGER NOT NULL
);

ALTER TABLE "appointment_status"
ADD CONSTRAINT unique_value UNIQUE (value);

INSERT INTO "appointment_status" (status, value) 
VALUES ('pending', 1),
       ('confirmed', 2),
       ('canceled', 3);