-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id  VARCHAR(150) PRIMARY KEY DEFAULT gen_random_uuid(),
    email CITEXT NOT NULL UNIQUE,
    email_verified TIMESTAMP WITH TIME ZONE,
    msisdn VARCHAR(15),
    first_name VARCHAR(150),
    last_name VARCHAR(150),
    full_name VARCHAR(150),
    image_url VARCHAR(150),
    password VARCHAR(250),
    account_id VARCHAR(150),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX ON "users" ("msisdn");
CREATE INDEX ON "users" ("account_id");


CREATE TABLE IF NOT EXISTS accounts (
    id  VARCHAR(150) PRIMARY KEY DEFAULT gen_random_uuid(),
    provider VARCHAR NOT NULL,
    provider_account_id VARCHAR NOT NULL,
    type VARCHAR(50) NOT NULL,
    expires_at INTEGER NOT NULL,
    token_type VARCHAR(50) NOT NULL,
    access_token VARCHAR,
    refresh_token VARCHAR,
    account_type VARCHAR(150),
    id_token VARCHAR,
    scope VARCHAR,
    user_id VARCHAR(150) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE UNIQUE INDEX ON "accounts" ("provider", "provider_account_id");
ALTER TABLE "accounts"
    ADD FOREIGN  KEY ("user_id") 
    REFERENCES "users" ("id")
    ON DELETE CASCADE;
-- +goose StatementEnd  

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
