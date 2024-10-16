package repositories

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BillingRepository interface {
	GetOutstanding(ctx context.Context, loanId uuid.UUID) (outstandingAmount float64, err error)
	IsDelinquent(ctx context.Context, loanId uuid.UUID, currentDt time.Time) (isDelinquent bool, err error)
	MakePayment(ctx context.Context, loanId uuid.UUID, paymentAmount float64, paymentDt time.Time) (err error)
}

type billingRepositoryImpl struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}
