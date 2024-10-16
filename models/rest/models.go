package rest

import "github.com/google/uuid"

type GetOutstandingAmountResponse struct {
	LoanId            uuid.UUID `json:"loanId"`
	OutstandingAmount float64   `json:"outstandingAmount"`
}

type IsDelinquentResponse struct {
	LoanId       uuid.UUID `json:"loanId"`
	IsDelinquent bool      `json:"isDelinquent"`
}

type MakePaymentRequest struct {
	LoanId        string  `json:"loanId"`
	PaymentAmount float64 `json:"paymentAmount"`
	PaymentDt     *string `json:"paymentDt"`
}
