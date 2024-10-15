BEGIN;
ALTER TABLE deliveries DROP COLUMN medicine_id;
COMMIT;