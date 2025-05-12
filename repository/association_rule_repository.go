package repository

import (
	"algoritma-apriori/model/domain"
	"context"
)

type AssociationRuleRepository interface {
	FetchAll(ctx context.Context) ([]domain.Rule, error)
	Save(ctx context.Context, rules []domain.Rule) error
	DeleteAll(ctx context.Context) error
}
