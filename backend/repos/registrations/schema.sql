CREATE TABLE registrations (
    id SERIAL PRIMARY KEY,
    event_id INTEGER REFERENCES events(id) ON DELETE CASCADE,
    email VARCHAR NOT NULL,
    UNIQUE(event_id, email)
);