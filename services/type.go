package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/rosaekapratama/amartha-billing-engine-test/repositories"
	"time"
)

type BillingService interface {
	GetOutstanding(ctx context.Context, loanId uuid.UUID) (outstandingAmount float64, err error)
	IsDelinquent(ctx context.Context, loanId uuid.UUID) (isDelinquent bool, err error)
	MakePayment(ctx context.Context, loanId uuid.UUID, paymentAmount float64, paymentDt *time.Time) (err error)
}

type billingServiceImpl struct {
	billingRepository repositories.BillingRepository
}
