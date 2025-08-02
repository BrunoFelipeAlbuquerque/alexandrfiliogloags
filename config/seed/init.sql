-- STATUS ENUM
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'spaceship_status') THEN
        CREATE TYPE spaceship_status AS ENUM ('Operational', 'Damaged', 'Retired');
    END IF;
END
$$;

-- SPACESHIP TABLE
CREATE TABLE IF NOT EXISTS spaceships (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    class TEXT NOT NULL,
    crew INTEGER NOT NULL CHECK (crew >= 0),
    image TEXT,
    value DECIMAL(10,2) NOT NULL CHECK (value >= 0),
    status spaceship_status NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- ARMAMENTS TABLE
CREATE TABLE IF NOT EXISTS armaments (
    id SERIAL PRIMARY KEY,
    spaceship_id INTEGER NOT NULL REFERENCES spaceships(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    qty INTEGER NOT NULL CHECK (qty > 0)
);

-- DUMMY DATA
INSERT INTO spaceships (name, class, crew, image, value, status) VALUES
('Devastator', 'Star Destroyer', 35000, 'https://placehold.co/600x400?text=Devastator', 1999.99, 'Operational'),
('Red Five', 'X-Wing', 1, 'https://placehold.co/600x400?text=Red+Five', 149.99, 'Damaged'),
('Outrider', 'YT-2400', 4, 'https://placehold.co/600x400?text=Outrider', 299.99, 'Retired');

INSERT INTO armaments (spaceship_id, title, qty) VALUES
(1, 'Turbo Laser', 60),
(1, 'Ion Cannons', 60),
(1, 'Tractor Beam', 10),
(2, 'Laser Cannons', 4),
(3, 'Concussion Missiles', 2);
