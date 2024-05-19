-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS companies (
    id VARCHAR(21) PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL UNIQUE,
    slug VARCHAR(50) NOT NULL,
    about VARCHAR(250),
    msisdn VARCHAR(15) NOT NULL,
    email CITEXT NOT NULL,
    tin VARCHAR(50),
    image_url VARCHAR(250),
    banner_url VARCHAR(250),
    brand_type VARCHAR(20) NOT NULL DEFAULT 'FOOD',
    owner_id VARCHAR NOT NULL,
    total_sales INTEGER NOT NULL DEFAULT 0, 
    is_active BOOLEAN NOT NULL DEFAULT false,
    currency_code VARCHAR(5) NOT NULL DEFAULT 'GHS',
    invitation_code VARCHAR(10),

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON "companies" ("msisdn");
CREATE INDEX ON "companies" ("invitation_code");

ALTER TABLE "companies"
    ADD FOREIGN KEY ("owner_id")
    REFERENCES "users" ("id");

CREATE TABLE IF NOT EXISTS members (
    id  VARCHAR(150) PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    msisdn VARCHAR(15) NOT NULL,
    email CITEXT,
    role VARCHAR(50) NOT NULL DEFAULT 'AGENT',
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    company_id VARCHAR(21) NOT NULL,
    user_id VARCHAR(150),

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE UNIQUE INDEX ON "members" ("msisdn", "company_id");

-- TODO: write test for this;
-- CREATE UNIQUE INDEX ON "members" ("user_id", "company_id");

ALTER TABLE "members"
    ADD FOREIGN KEY ("company_id")
    REFERENCES "companies" ("id");

ALTER TABLE "members"
    ADD FOREIGN KEY ("user_id")
    REFERENCES "users" ("id");

CREATE TABLE IF NOT EXISTS configurations (
    id  VARCHAR(150) PRIMARY KEY DEFAULT gen_random_uuid(),
    delivery BOOLEAN NOT NULL DEFAULT false,
    pickup BOOLEAN NOT NULL DEFAULT true,
    cash_on_delivery BOOLEAN NOT NULL DEFAULT false,
    digital_payments BOOLEAN NOT NULL DEFAULT true,

    company_id VARCHAR(21) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX ON "configurations" ("company_id");

ALTER TABLE "configurations"
    ADD FOREIGN KEY ("company_id")
    REFERENCES "companies" ("id");

CREATE TABLE IF NOT EXISTS working_hours (
    id  VARCHAR(150) PRIMARY KEY DEFAULT gen_random_uuid(),
    day CITEXT NOT NULL,
    opens_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    closes_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    
    company_id VARCHAR(21) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX ON "working_hours" ("company_id", "day");

ALTER TABLE "working_hours"
    ADD FOREIGN KEY ("company_id")
    REFERENCES "companies" ("id");


CREATE TABLE IF NOT EXISTS payment_accounts (
    id  VARCHAR(150) PRIMARY KEY DEFAULT gen_random_uuid(),
    account_type VARCHAR(50) NOT NULL DEFAULT 'mobile_money',
    account_number VARCHAR(50) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    bank_name VARCHAR(100) NOT NULL,
    bank_code VARCHAR(50) NOT NULL,
    bank_branch VARCHAR(100) NOT NULL,

    company_id VARCHAR(21) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX ON "payment_accounts" ("company_id", "account_number");

ALTER TABLE "payment_accounts"
    ADD FOREIGN KEY ("company_id")
    REFERENCES "companies" ("id");

CREATE TABLE IF NOT EXISTS payout_accounts (
    id  VARCHAR(150) PRIMARY KEY DEFAULT gen_random_uuid(),
    id_int INTEGER NOT NULL,
    currency VARCHAR(20) NOT NULL DEFAULT 'GHS',
    business_name VARCHAR(100) NOT NULL ,
    account_number VARCHAR(50) NOT NULL,
    primay_contact_name VARCHAR(150) NOT NULL,
    primay_contact_email VARCHAR(150) NOT NULL,
    primay_contact_phone VARCHAR(150) NOT NULL,
    description VARCHAR(250),
    subaccount_code VARCHAR(150) NOT NULL,
    settlement_bank VARCHAR(150) NOT NULL,
    percentage_charge DOUBLE PRECISION NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,

    bank_id VARCHAR(150) NOT NULL,
    payment_account_id VARCHAR(150) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX ON "payout_accounts" ("payment_account_id");
CREATE INDEX ON "payout_accounts" ("id_int");

ALTER TABLE "payout_accounts"
    ADD FOREIGN KEY ("payment_account_id")
    REFERENCES "payment_accounts" ("id")
    ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS payment_account_details (
    id VARCHAR(150) PRIMARY KEY DEFAULT gen_random_uuid(),
    id_int INTEGER NOT NULL,
    name VARCHAR NOT NULL,
    slug VARCHAR NOT NULL,
    code VARCHAR NOT NULL,
    longcode VARCHAR,
    gateway VARCHAR,
    pay_with_bank BOOLEAN NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    country VARCHAR NOT NULL,
    currency TEXT NOT NULL,
    type VARCHAR(20) NOT NULL DEFAULT 'mobile_money',

    payment_account_id VARCHAR(150) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX ON "payment_account_details" ("payment_account_id");
CREATE INDEX ON "payment_account_details" ("id_int");

ALTER TABLE "payment_account_details"
    ADD FOREIGN KEY ("payment_account_id")
    REFERENCES "payment_accounts" ("id")
    ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE members;
DROP TABLE configurations;
DROP TABLE working_hours;
DROP TABLE payout_accounts;
DROP TABLE payment_account_details;
DROP TABLE payment_accounts;
DROP TABLE companies;
-- +goose StatementEnd
