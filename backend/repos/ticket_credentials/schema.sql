CREATE TABLE ticket_credentials (
    id VARCHAR PRIMARY KEY,
    email VARCHAR NOT NULL,
    event_id VARCHAR NOT NULL,
    data VARCHAR NOT NULL,
    issued_at TIMESTAMPTZ NOT NULL,
    expire_at TIMESTAMPTZ NOT NULL
);

-- Create index on (event_id, email)
CREATE INDEX idx_event_id_email ON ticket_credentials(event_id, email);

-- Create index on email
CREATE INDEX idx_email ON ticket_credentials(email);