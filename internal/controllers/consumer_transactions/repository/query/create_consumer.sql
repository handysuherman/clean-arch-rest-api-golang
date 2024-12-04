-- name: CreateConsumers :execresult
INSERT INTO consumers (
    nik,
    full_name,
    legal_name,
    birth_place,
    birth_date,
    salary,
    ktp_photo,
    selfie_photo,
    created_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
);

