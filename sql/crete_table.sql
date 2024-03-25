CREATE TABLE IF NOT EXISTS vehicles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    status VARCHAR(50),
    color VARCHAR(50),
    mileage INT,
    body_type VARCHAR(50),
    transmission VARCHAR(50),
    fuel_type VARCHAR(50),
    doors INT,
    review BOOLEAN,
    price DECIMAL,
    description TEXT
);