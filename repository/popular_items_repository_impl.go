package repository

import (
	"algoritma-apriori/model/domain"
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type PopularItemsRepositoryImpl struct {
	DB *firestore.Client
}

func NewPopularItemsRepository(DB *firestore.Client) PopularItemsRepository {
	return &PopularItemsRepositoryImpl{
		DB: DB,
	}
}

func (repository PopularItemsRepositoryImpl) Save(ctx context.Context, popular []domain.PopularItem) error {
	batch := repository.DB.Batch()
	rulesCol := repository.DB.Collection("popular_items")
	for i, p := range popular {
		batch.Set(rulesCol.Doc(fmt.Sprintf("item_%d", i)), p)
	}
	_, err := batch.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repository PopularItemsRepositoryImpl) DeleteAll(ctx context.Context) error {
	col := repository.DB.Collection("association_rules")
	bulkWriter := repository.DB.BulkWriter(ctx)

	for {
		iter := col.Limit(100).Documents(ctx)
		numDeleted := 0

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			bulkWriter.Delete(doc.Ref)
			numDeleted++
		}

		if numDeleted == 0 {
			bulkWriter.End()
			break
		}

		bulkWriter.Flush()
	}
	return nil
}
