CREATE TABLE appointments (
    id UUID PRIMARY KEY NOT NULL,
    datetime DATE NOT NULL,
    slot INTEGER REFERENCES slots (id) NOT NULL,
    candidate UUID REFERENCES candidates (id) NOT NULL,
    status INTEGER REFERENCES "appointment_status" (id) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now() 
);

ALTER TABLE appointments
ADD CONSTRAINT unique_datetime UNIQUE (datetime);