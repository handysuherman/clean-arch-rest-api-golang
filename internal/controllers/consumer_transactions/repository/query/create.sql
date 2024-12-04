-- name: Create :execresult
INSERT INTO consumer_transactions (
    consumer_id,
    contract_number,
    admin_fee_amount,
    installment_amount,
    otr_amount,
    interest_rate,
    transaction_date,
    affiliated_dealer_id,
    created_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
);

