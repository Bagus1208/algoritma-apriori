package controller

import (
	"algoritma-apriori/model/web/request"
	"context"
)

type RefreshController interface {
	Refresh(ctx context.Context, request request.AssociationRuleRefreshRequest)
}
