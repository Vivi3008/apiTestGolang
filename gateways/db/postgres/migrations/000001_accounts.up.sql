
CREATE TABLE IF NOT EXISTS accounts
(
	id uuid PRIMARY KEY,
	name varchar(255) NOT NULL,
	cpf TEXT NOT NULL UNIQUE,
	secret varchar(100) NOT NULL,
	balance int default 0 NOT NULL CHECK (balance >=0 ),
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
