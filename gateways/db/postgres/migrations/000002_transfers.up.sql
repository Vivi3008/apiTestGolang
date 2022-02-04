CREATE TABLE IF NOT EXISTS transfers
(
	id UUID PRIMARY KEY,
	account_origin_id UUID NOT NULL REFERENCES accounts (id),
	account_destination_id UUID NOT NULL REFERENCES accounts (id) CHECK (account_origin_id != account_destination_id),
	amount int default 0 NOT NULL CHECK (amount > 0 ),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);