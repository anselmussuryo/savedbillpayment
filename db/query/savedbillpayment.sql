-- name: CreateSavedBillPayment :execresult
INSERT INTO saved_bill_payment (
  merchant_id,
  customer_id,
  subsciber_no,
  description,
  isshowib,
  isshowmobile,
  amount,
  isfavorite
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetSavedBillPayment :one
SELECT * FROM saved_bill_payment
WHERE id = ? LIMIT 1;

-- name: ListSavedBillPayment :many
SELECT * FROM saved_bill_payment
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: UpdateSavedBillPayment :exec
UPDATE saved_bill_payment 
SET merchant_id = ?, customer_id = ?, subsciber_no = ?, description = ?, isshowib = ?, isshowmobile = ?, amount = ?, isfavorite = ?
WHERE id = ?;

-- name: DeleteSavedBillPayment :exec
DELETE FROM saved_bill_payment
WHERE id = ?;