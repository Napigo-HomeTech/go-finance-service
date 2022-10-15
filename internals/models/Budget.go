package models

type Budget struct {
	BudgetId string `json:"budget_id" validation:"required"`
	Revision string `json:"revision"  validation:"required"`
}
