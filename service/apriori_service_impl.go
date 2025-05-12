package service

import (
	"algoritma-apriori/helpers"
	"algoritma-apriori/model/domain"
	"algoritma-apriori/model/web/request"
	"algoritma-apriori/repository"
	"context"
	"sort"
)

type AprioriServiceImpl struct {
	TransactionRepository     repository.TransactionRepository
	PopularItemsRepository    repository.PopularItemsRepository
	AssociationRuleRepository repository.AssociationRuleRepository
}

func NewAprioriService(TransactionRepository repository.TransactionRepository, PopularItemsRepository repository.PopularItemsRepository, AssociationRuleRepository repository.AssociationRuleRepository) AprioriService {
	return &AprioriServiceImpl{
		TransactionRepository:     TransactionRepository,
		PopularItemsRepository:    PopularItemsRepository,
		AssociationRuleRepository: AssociationRuleRepository,
	}
}

func (service AprioriServiceImpl) Apriori(ctx context.Context, request request.AssociationRuleRefreshRequest) error {
	transactions, err := service.TransactionRepository.FetchAll(ctx)
	if err != nil {
		return err
	}

	itemsets := helpers.ExtractItemsets(transactions)
	freq := helpers.Apriori(itemsets, request.MinSupport)
	rules := helpers.GenerateRules(freq, request.MinConfidence)

	var popular []domain.PopularItem
	for _, is := range freq {
		if len(is.Itemset) == 1 {

			popular = append(popular, domain.PopularItem{MenuId: is.Itemset[0], Support: is.Support})
		}
	}
	sort.Slice(popular, func(i, j int) bool {
		return popular[i].Support > popular[j].Support
	})

	err = service.PopularItemsRepository.DeleteAll(ctx)
	if err != nil {
		return err
	}

	err = service.AssociationRuleRepository.DeleteAll(ctx)
	if err != nil {
		return err
	}

	err = service.PopularItemsRepository.Save(ctx, popular)
	if err != nil {
		return err
	}

	err = service.AssociationRuleRepository.Save(ctx, rules)
	if err != nil {
		return err
	}

	return nil
}
