-- name: CreateAffiliatedDealer :execresult
INSERT INTO affiliated_dealers (
    affiliated_dealer_name,
    created_at
) VALUES (
    ?, ?
);

