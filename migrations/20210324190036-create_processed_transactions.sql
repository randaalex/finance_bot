
-- +migrate Up
CREATE TABLE processed_transactions
(
    id         varchar(255) NOT NULL PRIMARY KEY UNIQUE ,
    created_at timestamp WITH TIME ZONE DEFAULT NOW()
);

-- +migrate Down
DROP TABLE processed_transactions;
