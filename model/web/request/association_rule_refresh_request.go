package request

type AssociationRuleRefreshRequest struct {
	MinSupport    float64 `json:"min_support"`
	MinConfidence float64 `json:"min_confidence"`
}
