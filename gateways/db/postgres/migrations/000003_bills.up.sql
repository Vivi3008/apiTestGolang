CREATE TABLE IF NOT EXISTS "bills" 
(
    "id" UUID PRIMARY KEY,
    "account_id" UUID NOT NULL REFERENCES accounts (id),
    "description" TEXT NOT NULL,
    "value" INT NOT NULL default 0 CHECK (value > 0),
    "due_date" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "scheduled_date" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "status" varchar(50) NOT NULL
);