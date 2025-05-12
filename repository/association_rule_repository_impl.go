package repository

import (
	"algoritma-apriori/model/domain"
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type AssociationRuleRepositoryImpl struct {
	DB *firestore.Client
}

func NewAssociationRuleRepository(DB *firestore.Client) AssociationRuleRepository {
	return &AssociationRuleRepositoryImpl{
		DB: DB,
	}
}

func (repository AssociationRuleRepositoryImpl) FetchAll(ctx context.Context) ([]domain.Rule, error) {
	var rules []domain.Rule

	iter := repository.DB.Collection("association-rule").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var rule domain.Rule
		if err = doc.DataTo(&rule); err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	return rules, nil
}

func (repository AssociationRuleRepositoryImpl) Save(ctx context.Context, rules []domain.Rule) error {
	batch := repository.DB.Batch()
	rulesCol := repository.DB.Collection("association_rules")
	for i, r := range rules {
		batch.Set(rulesCol.Doc(fmt.Sprintf("rule_%d", i)), r)
	}
	_, err := batch.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repository AssociationRuleRepositoryImpl) DeleteAll(ctx context.Context) error {
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
