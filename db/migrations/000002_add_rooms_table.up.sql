CREATE TABLE IF NOT EXISTS rooms(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    theater_id uuid,
    name varchar NOT NULL,
    rows int NOT NULL,
    columns int NOT NULL,
    CONSTRAINT "THEATRE_ID_FKEY" FOREIGN KEY (theater_id) REFERENCES theaters(id)
);