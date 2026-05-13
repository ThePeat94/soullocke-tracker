CREATE TABLE token_auth
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    token_hash BYTEA       NOT NULL UNIQUE,
    lobby_id   UUID REFERENCES lobbies (id) ON DELETE CASCADE NOT NULL,
    created_at TIMESTAMPTZ      DEFAULT now(),
    expires_at TIMESTAMPTZ NOT NULL
);