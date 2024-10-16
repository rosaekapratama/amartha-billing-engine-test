package routers

import (
	v1 "github.com/rosaekapratama/amartha-billing-engine-test/controllers/rest/v1"
	"github.com/rosaekapratama/go-starter/transport/restserver"
)

func Init(billingController v1.BillingRestController) {
	rv1 := restserver.Router.Group("/v1")
	rv1.GET("/billing/:loanId/outstanding", billingController.GetOutstanding)
	rv1.GET("/billing/:loanId/status", billingController.IsDelinquent)
	rv1.POST("/billing/payment", billingController.MakePayment)
}
