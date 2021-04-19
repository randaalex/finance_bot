
-- +migrate Up
CREATE TABLE processed_transactions
(
    hash       varchar(64) NOT NULL PRIMARY KEY UNIQUE,
    firefly_id integer,
    created_at timestamp WITH TIME ZONE DEFAULT NOW()
);

CREATE UNIQUE INDEX firefly_id_idx ON processed_transactions(firefly_id);

-- +migrate Down
DROP TABLE processed_transactions;
