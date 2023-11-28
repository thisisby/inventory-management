CREATE TABLE IF NOT EXISTS stores(
    id uuid PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    address VARCHAR(50) NOT NULL,
    contacts VARCHAR(50) NULL
);