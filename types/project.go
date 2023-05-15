package types

type decimal float64

type Project struct {
	EarnedValue  decimal `json:"earnedValue"`
	PlannedValue decimal `json:"plannedValue"`
	ActualCost   decimal `json:"actualCost"`
}

func (p *Project) CostVariance() decimal {
	return p.EarnedValue - p.ActualCost
}

func (p *Project) CostPerformanceIndex() decimal {
	return p.EarnedValue / p.ActualCost
}

func (p *Project) ScopeVariance() decimal {
	return p.EarnedValue - p.PlannedValue
}

func (p *Project) ScopePerformanceIndex() decimal {
	return p.EarnedValue / p.PlannedValue
}
