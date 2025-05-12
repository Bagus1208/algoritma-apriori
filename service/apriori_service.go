package service

import (
	"algoritma-apriori/model/web/request"
	"context"
)

type AprioriService interface {
	Apriori(ctx context.Context, request request.AssociationRuleRefreshRequest) error
}
