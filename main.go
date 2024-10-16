package main

import (
	"context"
	v1 "github.com/rosaekapratama/amartha-billing-engine-test/controllers/rest/v1"
	"github.com/rosaekapratama/amartha-billing-engine-test/repositories"
	"github.com/rosaekapratama/amartha-billing-engine-test/routers"
	"github.com/rosaekapratama/amartha-billing-engine-test/services"
	"github.com/rosaekapratama/go-starter/app"
)

func main() {
	ctx := context.Background()

	// Init repos
	billingRepository := repositories.NewBillingRepository(ctx)

	// Init services
	billingService := services.NewBillingService(ctx, billingRepository)

	// Init rest controller
	billingRestController := v1.NewBillingRestController(ctx, billingService)

	// Init rest router
	routers.Init(billingRestController)

	app.Run()
}
