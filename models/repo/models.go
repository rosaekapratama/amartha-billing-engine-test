package repo

import (
	"github.com/google/uuid"
	"time"
)

// Billing represents the billings table in the database.
type Billing struct {
	ID                         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	LoanID                     uuid.UUID `gorm:"type:uuid;not null"`
	InstallmentNo              int       `gorm:"type:int;not null"`
	InstallmentPrincipalAmount float64   `gorm:"type:numeric(15,3);not null"`
	InstallmentPaymentAmount   float64   `gorm:"type:numeric(15,3);not null"`
	InstallmentInterestAmount  float64   `gorm:"type:numeric(15,3);"`
	InterestRate               float64   `gorm:"type:numeric(7,5);not null"`
	DueDt                      time.Time `gorm:"type:timestamptz;not null"`
}

// Payment represents the payments table in the database.
type Payment struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	LoanID        uuid.UUID `gorm:"type:uuid;not null"`
	PaymentAmount float64   `gorm:"type:numeric(15,3);not null"`
	PaymentDt     time.Time `gorm:"type:timestamptz;not null"`
}
