CREATE TABLE weekdays (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    value INTEGER NOT NULL
);

INSERT INTO weekdays (name, value) 
VALUES ('segunda-feira', 1),
       ('terça-feira', 2),
       ('quarta-feira', 3),
       ('quinta-feira', 4),
       ('sexta-feira', 5);