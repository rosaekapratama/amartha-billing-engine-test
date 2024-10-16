package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rosaekapratama/amartha-billing-engine-test/services"
)

type BillingRestController interface {
	GetOutstanding(c *gin.Context)
	IsDelinquent(c *gin.Context)
	MakePayment(c *gin.Context)
}

type billingRestControllerImpl struct {
	billingService services.BillingService
}
