-- name: CreatePaymentAccount :one
INSERT INTO payment_accounts (
        company_id,
        account_number,
        bank_type,
        first_name,
        last_name,
        bank_code,
        bank_name,
        bank_id,
        bank_currency
    )
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
    )
RETURNING *;

-- name: UpdatePaymentAccountByCompanyID :one
UPDATE payment_accounts SET 
        account_number = $2,
        bank_type = $3,
        first_name = $4,
        last_name = $5,
        bank_code = $6,
        bank_name = $7,
        bank_id = $8,
        bank_currency = $9
    WHERE company_id = $1
RETURNING *;


-- name: GetPaymentAccountByCompanyID :one
SELECT * FROM payment_accounts
 WHERE company_id = $1
 LIMIT 1;

 -- name: RemoveEmployee :exec
DELETE FROM payment_accounts WHERE id = $1;


-- PAYOUT ACCOUNT

-- name: CreatePayoutAccount :one
INSERT INTO payout_accounts (
        payment_account_id,
        id_int,
        currency,
        business_name,
        account_number,
        primay_contact_name,
        primay_contact_email,
        primay_contact_phone,
        description,
        subaccount_code,
        settlement_bank,
        percentage_charge,
        active,
        bank_id
    )
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
    )
RETURNING *;

-- name: UpdatePayoutAccountByPaymentAccountID :one
UPDATE payout_accounts SET 
        percentage_charge = $2,
        active = $3,
        account_number = $4,
        subaccount_code = $5,
        settlement_bank = $6

    WHERE payment_account_id = $1
RETURNING *;


-- name: GetPayoutAccountByPaymentAccountID :one
SELECT * FROM payout_accounts
 WHERE payment_account_id = $1
 LIMIT 1;


-- PAYMENT ACCOUNT DETAILS

-- name: CreatePaymentAccountDetails :one
INSERT INTO payment_account_details (
        payment_account_id,
        id_int,
        currency,
        name,
        slug,
        code,
        longcode,
        gateway,
        pay_with_bank,
        is_deleted,
        country,
        type,
        active
    )
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
    )
RETURNING *;

-- name: UpdatePaymentAccountDetailsByPaymentAccountID :one
UPDATE payment_account_details SET 
        name = $2,
        active = $3,
        currency = $4,
        type = $5,
        slug = $6,
        code = $7
        
    WHERE payment_account_id = $1
RETURNING *;

-- name: GetPaymentAccountDetailsByPaymentAccountID :one
SELECT * FROM payment_account_details
 WHERE payment_account_id = $1
 LIMIT 1;
