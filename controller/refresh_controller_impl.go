package controller

import (
	"algoritma-apriori/model/web/request"
	"algoritma-apriori/service"
	"context"
	"fmt"
)

type RefreshControllerImpl struct {
	AprioriService service.AprioriService
}

func NewRefreshController(aprioriService service.AprioriService) RefreshController {
	return &RefreshControllerImpl{
		AprioriService: aprioriService,
	}
}

func (controller RefreshControllerImpl) Refresh(ctx context.Context, request request.AssociationRuleRefreshRequest) {

	err := controller.AprioriService.Apriori(ctx, request)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses Analisis Data Menggunakan Algoritma Apriori")
}
