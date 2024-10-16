package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/rosaekapratama/amartha-billing-engine-test/repositories"
	"github.com/rosaekapratama/go-starter/log"
	"github.com/rosaekapratama/go-starter/response"
	"time"
)

func NewBillingService(_ context.Context, billingRepository repositories.BillingRepository) BillingService {
	return &billingServiceImpl{
		billingRepository: billingRepository,
	}
}

// GetOutstanding returns the current outstanding on a loan, or 0 if there is none.
func (s *billingServiceImpl) GetOutstanding(ctx context.Context, loanId uuid.UUID) (outstandingAmount float64, err error) {
	outstandingAmount, err = s.billingRepository.GetOutstanding(ctx, loanId)
	if err != nil {
		log.Errorf(ctx, err, "error on s.billingRepository.GetOutstanding(ctx, loanId), loanId=%s", loanId)
	}
	return
}

// IsDelinquent checks if there are more than 2 weeks of non-payment on the loan.
func (s *billingServiceImpl) IsDelinquent(ctx context.Context, loanId uuid.UUID) (isDelinquent bool, err error) {
	currentDt := time.Now()
	isDelinquent, err = s.billingRepository.IsDelinquent(ctx, loanId, currentDt)
	if err != nil {
		log.Errorf(ctx, err, "error on s.billingRepository.IsDelinquent(ctx, loanId), loanId=%s", loanId)
	}
	return
}

// MakePayment records a payment for a loan.
func (s *billingServiceImpl) MakePayment(ctx context.Context, loanId uuid.UUID, paymentAmount float64, paymentDt *time.Time) (err error) {
	if paymentAmount <= 0 || paymentDt == nil {
		return response.InvalidArgument
	}

	err = s.billingRepository.MakePayment(ctx, loanId, paymentAmount, *paymentDt)
	if err != nil {
		log.Errorf(ctx, err, "error on s.billingRepository.MakePayment(ctx, loanId, paymentAmount), loanId=%s, paymentAmount=%f", loanId, paymentAmount)
	}
	return
}
