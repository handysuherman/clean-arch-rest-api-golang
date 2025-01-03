// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"database/sql"
)

type AffiliatedDealer struct {
	ID                   int64  `json:"id"`
	AffiliatedDealerName string `json:"affiliated_dealer_name"`
	// format should be like 0001-01-01 00:00:00Z
	CreatedAt string `json:"created_at"`
	// format should be like 0001-01-01 00:00:00Z
	UpdatedAt   string         `json:"updated_at"`
	UpdatedBy   sql.NullString `json:"updated_by"`
	IsActivated bool           `json:"is_activated"`
	// format should be like 0001-01-01 00:00:00Z
	IsActivatedAt string `json:"is_activated_at"`
	// format should be like 0001-01-01 00:00:00Z
	IsActivatedUpdatedAt string `json:"is_activated_updated_at"`
}
