package main

import (
	"algoritma-apriori/app"
	"algoritma-apriori/controller"
	"algoritma-apriori/model/web/request"
	"algoritma-apriori/repository"
	"algoritma-apriori/service"
	"context"
	"fmt"
)

func main() {
	db := app.NewDB()
	ctx := context.Background()

	var minSupport float64
	var minConfidence float64

	fmt.Print("Masukkan nilai minimum support: ")
	fmt.Scanln(&minSupport)
	fmt.Print("Masukkan nilai minimum confidence: ")
	fmt.Scanln(&minConfidence)

	transactionRepository := repository.NewTransactionRepositoryImpl(db)
	popularItemsRepository := repository.NewPopularItemsRepository(db)
	associationRuleRepository := repository.NewAssociationRuleRepository(db)
	aprioriService := service.NewAprioriService(transactionRepository, popularItemsRepository, associationRuleRepository)
	calculate := controller.NewRefreshController(aprioriService)

	calculate.Refresh(ctx, request.AssociationRuleRefreshRequest{MinSupport: minSupport, MinConfidence: minConfidence})
}
