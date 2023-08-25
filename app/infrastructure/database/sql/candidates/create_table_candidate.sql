CREATE TABLE candidates (
    id UUID PRIMARY KEY NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(12) NOT NULL,
    "resumeUrl" VARCHAR(255) -- NOT NULL (?)
);

ALTER TABLE candidates
ADD CONSTRAINT unique_email UNIQUE (email);

ALTER TABLE candidates
ADD CONSTRAINT unique_phone UNIQUE (phone);