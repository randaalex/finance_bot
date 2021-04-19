
-- +migrate Up
CREATE TABLE mappings
(
    id                  serial NOT NULL PRIMARY KEY,
    transaction_details varchar(255) NOT NULL,
    category_id         integer NOT NULL,
    created_at          timestamp WITH TIME ZONE DEFAULT NOW(),
    updated_at          timestamp WITH TIME ZONE
);

CREATE UNIQUE INDEX transaction_details_idx ON mappings(transaction_details);

-- +migrate Down
DROP TABLE mappings;
