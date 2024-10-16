package repositories

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/rosaekapratama/amartha-billing-engine-test/models/repo"
	"github.com/rosaekapratama/go-starter/database"
	"github.com/rosaekapratama/go-starter/log"
	"gorm.io/gorm"
	"time"
)

func NewBillingRepository(ctx context.Context) BillingRepository {
	gormDB, sqlDB, err := database.Manager.DB(ctx, "playground")
	if err != nil {
		log.Fatal(ctx, err, "error on database.Manager.DB(ctx, \"playground\")")
	}

	return &billingRepositoryImpl{
		gormDB: gormDB,
		sqlDB:  sqlDB,
	}
}

// GetOutstanding returns the current outstanding on a loan, or 0 if there is none.
func (r *billingRepositoryImpl) GetOutstanding(ctx context.Context, loanId uuid.UUID) (outstandingAmount float64, err error) {
	err = r.gormDB.WithContext(ctx).Raw(queryGetOutstanding, loanId, loanId).Scan(&outstandingAmount).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorf(ctx, err, "error on queryGetOutstanding, loanId=%s", loanId)
	}
	return
}

// IsDelinquent checks if there are more than 2 weeks of non-payment on the loan.
func (r *billingRepositoryImpl) IsDelinquent(ctx context.Context, loanId uuid.UUID, currentDate time.Time) (isDelinquent bool, err error) {
	err = r.gormDB.WithContext(ctx).Raw(queryIsDelinquent, loanId, currentDate, currentDate, currentDate).Scan(&isDelinquent).Error
	if err != nil {
		log.Errorf(ctx, err, "error on queryIsDelinquent, loanId=%s", loanId)
	}
	return
}

// MakePayment records a payment for a loan.
func (r *billingRepositoryImpl) MakePayment(ctx context.Context, loanId uuid.UUID, paymentAmount float64, paymentDt time.Time) (err error) {
	payment := repo.Payment{
		LoanID:        loanId,
		PaymentAmount: paymentAmount,
		PaymentDt:     paymentDt, // Set payment date to now
	}

	// Insert the payment into the payments table
	err = r.gormDB.WithContext(ctx).Create(&payment).Error
	return
}
