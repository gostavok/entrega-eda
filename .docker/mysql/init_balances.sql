CREATE DATABASE IF NOT EXISTS balances;
USE balances;

CREATE TABLE IF NOT EXISTS balances (
    account_id VARCHAR(255) PRIMARY KEY,
    balance INT
);

INSERT IGNORE INTO balances (account_id, balance) VALUES
    ('2a9d13a4-2f0b-49a8-bc47-23c14fb22c5e', 1000),
    ('fb4c6c91-2f45-4d91-8d4f-3e6b5d61c6b2', 500),
    ('59e1fa4e-5c7b-4c52-94a9-8f39c6d0f4e1', 200);
