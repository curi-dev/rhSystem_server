CREATE TABLE "access_keys" (
    id UUID PRIMARY KEY NOT NULL,
    value VARCHAR(255) NOT NULL,
    candidate UUID NOT NULL REFERENCES candidates (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now()
);