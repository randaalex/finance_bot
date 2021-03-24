
-- +migrate Up
CREATE TABLE mappings
(
    id               serial NOT NULL PRIMARY KEY,
    transaction_hash varchar(255) NOT NULL,
    transaction_text varchar(255) NOT NULL,
    category_id      integer NOT NULL,
    created_at       timestamp WITH TIME ZONE DEFAULT NOW(),
    updated_at       timestamp WITH TIME ZONE
);

CREATE UNIQUE INDEX transaction_hash_idx ON mappings(transaction_hash);

-- +migrate Down
DROP TABLE mappings;
