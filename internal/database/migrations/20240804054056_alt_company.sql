-- +goose Up
-- +goose StatementBegin
ALTER TABLE "companies"
    ADD COLUMN IF NOT EXISTS "has_onboarded" BOOLEAN NOT NULL DEFAULT false,
    ADD COLUMN IF NOT EXISTS "is_deleted" BOOLEAN NOT NULL DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "companies"
    DROP COLUMN IF EXISTS "has_onboarded",
    DROP COLUMN IF EXISTS "is_deleted";
-- +goose StatementEnd
