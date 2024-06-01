-- name: AddMember :one
INSERT INTO members (
        company_id,
        name,
        email,
        msisdn,
        role,
        user_id,
        status
    )
VALUES (
    $1, $2, $3, $4, $5, $6, $7
    )
RETURNING *;

-- name: GetMembersByCompanyID :many
SELECT * FROM members
 WHERE company_id = $1
 ORDER BY created_at desc
 LIMIT $2 OFFSET $3;

-- name: GetUserTeams :many
SELECT members.*,  
    companies.email as company_email,
    companies.name as company_name,
    companies.slug as company_slug,
    companies.banner_url as company_banner_url,
    companies.image_url as company_image_url,
    companies.about as company_about,
    companies.is_active as company_is_active
FROM members
JOIN companies ON members.company_id = companies.id
WHERE members.user_id = sqlc.arg(user_id);


--SELECT members.*, companies.*
--FROM members
--JOIN companies ON members.company_id = companies.id
--WHERE members.user_id = sqlc.arg(user_id);

-- name: UpdateMember :one
UPDATE members SET 
        name = $2, 
        email = $3,
        role = $4,
        status = $5
    WHERE id = $1
RETURNING *;

-- name: UpdateMemberInvitationStatus :one
UPDATE members SET 
        status = $2
    WHERE id = $1
RETURNING *;

 -- name: RemoveMember :exec
DELETE FROM members WHERE id = $1;
