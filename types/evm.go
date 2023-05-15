package types

type decimal float64

type EarnedValueManagement struct {
	EarnedValue  decimal `json:"earnedValue"`
	PlannedValue decimal `json:"plannedValue"`
	ActualCost   decimal `json:"actualCost"`
}

func (evm *EarnedValueManagement) CostVariance() decimal {
	return evm.EarnedValue - evm.ActualCost
}

func (evm *EarnedValueManagement) CostPerformanceIndex() decimal {
	return evm.EarnedValue / evm.ActualCost
}

func (evm *EarnedValueManagement) ScopeVariance() decimal {
	return evm.EarnedValue - evm.PlannedValue
}

func (evm *EarnedValueManagement) ScopePerformanceIndex() decimal {
	return evm.EarnedValue / evm.PlannedValue
}
