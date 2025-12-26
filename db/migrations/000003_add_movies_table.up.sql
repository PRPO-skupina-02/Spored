CREATE TABLE IF NOT EXISTS movies(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    title varchar NOT NULL,
    description varchar NOT NULL,
    image_url varchar NOT NULL,
    rating real NOT NULL
);