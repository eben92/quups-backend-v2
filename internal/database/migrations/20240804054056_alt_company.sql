-- +goose Up
-- +goose StatementBegin
ALTER TABLE "companies"
    ADD COLUMN IF NOT EXISTS "has_onboarded" BOOLEAN NOT NULL DEFAULT false,
    ADD COLUMN IF NOT EXISTS "is_deleted" BOOLEAN NOT NULL DEFAULT false;

ALTER TABLE "payment_accounts"
    ADD COLUMN IF NOT EXISTS "is_deleted" BOOLEAN NOT NULL DEFAULT false,
    ADD COLUMN IF NOT EXISTS "bank_id" VARCHAR(50) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "companies"
    DROP COLUMN IF EXISTS "has_onboarded",
    DROP COLUMN IF EXISTS "is_deleted";

ALTER TABLE "payment_accounts"
    DROP COLUMN IF EXISTS "is_deleted",
    DROP COLUMN IF EXISTS "bank_id";
-- +goose StatementEnd
