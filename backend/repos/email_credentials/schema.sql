CREATE TABLE email_credentials (
    id VARCHAR PRIMARY KEY,
    identity_commitment VARCHAR NOT NULL,
    data VARCHAR NOT NULL,
    issued_at TIMESTAMPTZ NOT NULL,
    expire_at TIMESTAMPTZ NOT NULL
);

CREATE UNIQUE INDEX idx_identity_commitment ON email_credentials(identity_commitment);