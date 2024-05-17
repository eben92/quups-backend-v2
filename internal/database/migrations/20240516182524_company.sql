-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS companies (
    id VARCHAR(21) PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL UNIQUE,
    slug VARCHAR(50) NOT NULL,
    about VARCHAR(250),
    msisdn VARCHAR(15) NOT NULL,
    email VARCHAR(100),
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

CREATE TABLE IF NOT EXISTS company_employees (
    id  VARCHAR(150) PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    msisdn VARCHAR(15) NOT NULL,
    email VARCHAR(100),
    role VARCHAR(50) NOT NULL DEFAULT 'AGENT',
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    company_id VARCHAR(21) NOT NULL,
    user_id VARCHAR(150),

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE UNIQUE INDEX ON "company_employees" ("msisdn", "company_id");

-- TODO: write test for this;
-- CREATE UNIQUE INDEX ON "company_employees" ("user_id", "company_id");

ALTER TABLE "company_employees"
    ADD FOREIGN KEY ("company_id")
    REFERENCES "companies" ("id");

ALTER TABLE "company_employees"
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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE company_employees;
DROP TABLE configurations;
DROP TABLE companies;
-- +goose StatementEnd
