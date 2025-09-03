CREATE DATABASE IF NOT EXISTS wallet;
USE wallet;

CREATE TABLE IF NOT EXISTS clients (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    created_at DATETIME
);

CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(255) PRIMARY KEY,
    client_id VARCHAR(255),
    balance INT,
    created_at DATETIME
);

CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(255) PRIMARY KEY,
    account_id_from VARCHAR(255),
    account_id_to VARCHAR(255),
    amount INT,
    created_at DATETIME
);

INSERT IGNORE INTO clients (id, name, email, created_at) VALUES
    ('7b1d43b6-1f22-4f5f-a72d-1d2e2d4695f9', 'Alice', 'alice@example.com', NOW()),
    ('a08cfb6a-6b5e-4d54-9c03-5c785d420fd7', 'Bob', 'bob@example.com', NOW()),
    ('d6242bb5-b42c-41ce-95b7-0c6d0ff6f229', 'Carol', 'carol@example.com', NOW());

INSERT IGNORE INTO accounts (id, client_id, balance, created_at) VALUES
    ('2a9d13a4-2f0b-49a8-bc47-23c14fb22c5e', '7b1d43b6-1f22-4f5f-a72d-1d2e2d4695f9', 1000, NOW()),
    ('fb4c6c91-2f45-4d91-8d4f-3e6b5d61c6b2', 'a08cfb6a-6b5e-4d54-9c03-5c785d420fd7', 500, NOW()),
    ('59e1fa4e-5c7b-4c52-94a9-8f39c6d0f4e1', 'd6242bb5-b42c-41ce-95b7-0c6d0ff6f229', 200, NOW());
