package models

type CostDto struct {
	Expenses []GraphCost `json:"expenses" bson:"expenses"`
	Incomes  []GraphCost `json:"incomes" bson:"incomes"`
}

type GraphCost struct {
	Name  string  `json:"name" bson:"name"`
	Value float64 `json:"value" bson:"value"`
}
