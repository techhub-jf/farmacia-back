BEGIN;

CREATE TABLE IF NOT EXISTS Medicine (
    id BIGSERIAL PRIMARY KEY,
    reference VARCHAR(255) NOT NULL,
    client_id BIGINT NOT NULL,
    medicine_id BIGINT NOT NULL,
    qty INT NOT NULL,
    unit_id INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

COMMIT;
