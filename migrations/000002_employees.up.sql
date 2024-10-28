CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT CHECK (age > 0) NOT NULL,
    salary DECIMAL(10, 2) CHECK (salary >= 0) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);