
-- +migrate Up
CREATE TABLE transactions_logs
(
    id          serial       NOT NULL PRIMARY KEY,
    description varchar(255) NOT NULL,
    category_id integer      NOT NULL,
    created_at  timestamp WITH TIME ZONE DEFAULT NOW(),
    updated_at  timestamp WITH TIME ZONE
);

CREATE UNIQUE INDEX description_idx ON transactions_logs(description);

-- +migrate Down
DROP TABLE transactions_logs;
