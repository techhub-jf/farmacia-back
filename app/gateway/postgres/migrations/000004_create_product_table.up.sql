BEGIN;
CREATE TABLE IF NOT EXISTS Product (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    reference TEXT NOT NULL,
    brand TEXT NOT NULL,
    description TEXT NOT NULL,
    type_id INT NOT NULL,
    stock INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
COMMIT;