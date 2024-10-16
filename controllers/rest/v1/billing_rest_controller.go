package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rosaekapratama/amartha-billing-engine-test/models/rest"
	thisResponse "github.com/rosaekapratama/amartha-billing-engine-test/responses"
	"github.com/rosaekapratama/amartha-billing-engine-test/services"
	"github.com/rosaekapratama/go-starter/constant/str"
	"github.com/rosaekapratama/go-starter/log"
	"github.com/rosaekapratama/go-starter/response"
	"github.com/rosaekapratama/go-starter/transport/restserver"
	"time"
)

func NewBillingRestController(_ context.Context, billingService services.BillingService) BillingRestController {
	return &billingRestControllerImpl{
		billingService: billingService,
	}
}

func validateLoanId(c *gin.Context, loanIdStr string) (loanId uuid.UUID, isValid bool) {
	ctx := c.Request.Context()
	w := c.Writer

	// Return invalid argument if loanId is empty
	if loanIdStr == str.Empty {
		log.Tracef(ctx, "loan ID is not valid, loanId=%s", loanIdStr)
		restserver.SetResponse(w, response.InvalidArgument)
		return
	}

	// Return invalid argument if loanId is not valid UUID
	loanId, err := uuid.Parse(loanIdStr)
	if err != nil {
		log.Tracef(ctx, "loan ID is not valid, loanId=%s", loanIdStr)
		restserver.SetResponse(w, response.InvalidArgument)
		return
	}
	isValid = true
	return
}

func (ctrl *billingRestControllerImpl) GetOutstanding(c *gin.Context) {
	ctx := c.Request.Context()
	w := c.Writer

	loanIdStr := c.Param(pathParamLoadId)
	if loanId, isValid := validateLoanId(c, loanIdStr); isValid {
		outstandingAmount, err := ctrl.billingService.GetOutstanding(ctx, loanId)
		if err != nil {
			log.Error(ctx, err, "error on ctrl.billingService.GetOutstanding(ctx, loanId), loanId=%s", loanId)
			restserver.SetResponse(w, response.GeneralError)
			return
		}

		res := &rest.GetOutstandingAmountResponse{
			LoanId:            loanId,
			OutstandingAmount: outstandingAmount,
		}

		restserver.SetResponse(w, response.Success)
		c.JSON(response.Success.HttpStatusCode(), res)
	}
}

func (ctrl *billingRestControllerImpl) IsDelinquent(c *gin.Context) {
	ctx := c.Request.Context()
	w := c.Writer

	loanIdStr := c.Param(pathParamLoadId)
	if loanId, isValid := validateLoanId(c, loanIdStr); isValid {
		isDelinquent, err := ctrl.billingService.IsDelinquent(ctx, loanId)
		if err != nil {
			log.Error(ctx, err, "error on ctrl.billingService.IsDelinquent(ctx, loanId), loanId=%s", loanId)
			restserver.SetResponse(w, response.GeneralError)
			return
		}

		res := &rest.IsDelinquentResponse{
			LoanId:       loanId,
			IsDelinquent: isDelinquent,
		}

		restserver.SetResponse(w, response.Success)
		c.JSON(response.Success.HttpStatusCode(), res)
	}
}
func (ctrl *billingRestControllerImpl) MakePayment(c *gin.Context) {
	ctx := c.Request.Context()
	w := c.Writer

	req := &rest.MakePaymentRequest{}
	err := c.BindJSON(req)
	if err != nil {
		log.Error(ctx, err, "error on c.BindJSON(req)")
		restserver.SetResponse(w, response.InvalidArgument)
		return
	}

	if req.PaymentAmount <= 0 {
		log.Tracef(ctx, "amount must be greater than 0, loanId=%s, paymentAmount=%f", req.LoanId, req.PaymentAmount)
		restserver.SetResponse(w, thisResponse.PaymentAmountMustBeGreaterThanZero)
		return
	}

	paymentDt := time.Now()
	if req.PaymentDt != nil {
		paymentDt, err = time.Parse(time.RFC3339, *req.PaymentDt)
		if err != nil {
			log.Warnf(ctx, "invalid paymentDt value, paymentDt=%s, error=%s", *req.PaymentDt, err.Error())
			restserver.SetResponse(w, response.InvalidArgument)
			return
		}
	}

	if loanId, isValid := validateLoanId(c, req.LoanId); isValid {
		err := ctrl.billingService.MakePayment(ctx, loanId, req.PaymentAmount, &paymentDt)
		if err != nil {
			log.Error(ctx, err, "error on ctrl.billingService.MakePayment(ctx, loanId, req.PaymentAmount), loanId=%s, paymentAmount=%f", loanId, req.PaymentAmount)
			restserver.SetResponse(w, response.GeneralError)
			return
		}

		restserver.SetResponse(w, response.Success)
	}
}
