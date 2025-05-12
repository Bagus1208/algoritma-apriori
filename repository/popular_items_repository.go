package repository

import (
	"algoritma-apriori/model/domain"
	"context"
)

type PopularItemsRepository interface {
	Save(ctx context.Context, popular []domain.PopularItem) error
	DeleteAll(ctx context.Context) error
}
