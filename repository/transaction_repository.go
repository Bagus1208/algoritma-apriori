package repository

import (
	"algoritma-apriori/model/domain"
	"context"
)

type TransactionRepository interface {
	FetchAll(ctx context.Context) ([]domain.Transaction, error)
}
