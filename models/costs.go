package models

type CostDto struct {
	Expenses []GraphCost `json:"expenses" bson:"expenses"`
	Incomes  []GraphCost `json:"incomes" bson:"incomes"`
}

type GraphCost struct {
	Day   int     `json:"day" bson:"day"`
	Value float64 `json:"value" bson:"value"`
}
