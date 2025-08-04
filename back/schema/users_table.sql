CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    document VARCHAR(50) NOT NULL,
    company_name VARCHAR(100) NULL,
    type VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL
);
