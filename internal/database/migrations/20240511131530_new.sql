-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id  VARCHAR(150) PRIMARY KEY DEFAULT gen_random_uuid(),
    email CITEXT NOT NULL UNIQUE,
    msisdn VARCHAR(15),
    email_verified TIMESTAMP WITH TIME ZONE,
    name VARCHAR(50),
    image_url VARCHAR(150),
    tin_number VARCHAR(30),
    gender VARCHAR(10),
    dob TIMESTAMP WITHOUT TIME ZONE,
    otp VARCHAR(150),
    app_push_token VARCHAR(150),
    web_push_token VARCHAR(150),

    password VARCHAR(250),

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX ON "users" ("msisdn");

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


CREATE TABLE IF NOT EXISTS addresses (
    id  VARCHAR(150) PRIMARY KEY DEFAULT gen_random_uuid(),
    street VARCHAR(100) NOT NULL,
    city VARCHAR(70) NOT NULL,
    region VARCHAR(70),
    country VARCHAR(70) NOT NULL,
    country_code VARCHAR(5) NOT NULL,
    formatted_address VARCHAR,
    description VARCHAR,
    postal_code VARCHAR(20),
    latitude FLOAT,
    longitude FLOAT,
    msisdn VARCHAR,
    is_default BOOLEAN DEFAULT FALSE,
    
    user_id VARCHAR(150),
    company_id VARCHAR(150),

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE "addresses"
    ADD FOREIGN  KEY ("user_id") 
    REFERENCES "users" ("id");

-- ALTER TABLE "addresses"
--     ADD FOREIGN  KEY ("company_id") 
--     REFERENCES "company" ("id");



-- +goose StatementEnd  

-- +goose Down
-- +goose StatementBegin
DROP TABLE accounts;
DROP TABLE addresses;
DROP TABLE users;
-- +goose StatementEnd
