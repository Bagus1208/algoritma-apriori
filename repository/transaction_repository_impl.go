package repository

import (
	"algoritma-apriori/model/domain"
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type TransactionRepositoryImpl struct {
	DB *firestore.Client
}

func NewTransactionRepositoryImpl(DB *firestore.Client) TransactionRepository {
	return &TransactionRepositoryImpl{
		DB: DB,
	}
}

func (repository TransactionRepositoryImpl) FetchAll(ctx context.Context) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	iter := repository.DB.Collection("transactions").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var transaction domain.Transaction
		err = doc.DataTo(&transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
