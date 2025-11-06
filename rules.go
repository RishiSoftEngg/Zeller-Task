package main

// PricingRuleIface defines an interface for flexible pricing strategies
type PricingRuleIface interface {
	AppliesTo(sku string) bool
	Apply(items []Item) float64
}

type MultiBuyRule struct {
	SKU       string
	Required  int
	PayFor    int
	UnitPrice float64
}

func (r MultiBuyRule) AppliesTo(sku string) bool {
	return r.SKU == sku
}

// @description MultiBuyRule implements a "buy X, pay for Y" deal (e.g., 3 for 2)
//
// @param
// items []Item - The list of items to apply the rule to
//
// @return
// float64 - The total price after applying the rule
func (r MultiBuyRule) Apply(items []Item) float64 {
	count := 0
	for _, item := range items {
		if item.SKU == r.SKU {
			count++
		}
	}

	if count < r.Required {
		return float64(count) * r.UnitPrice
	}

	totalSets := count / r.Required

	remainingItems := count % r.Required

	paidItems := (totalSets * r.PayFor) + remainingItems

	return float64(paidItems) * r.UnitPrice
}

type BulkDiscountRule struct {
	SKU          string
	Threshold    int
	RegularPrice float64
	Discounted   float64
}

func (r BulkDiscountRule) AppliesTo(sku string) bool {
	return r.SKU == sku
}

// @description BulkDiscountRule implements a bulk discount for a specific SKU
//
// @param
// items []Item - The list of items to apply the rule to
//
// @return
// float64 - The total price after applying the rule
func (r BulkDiscountRule) Apply(items []Item) float64 {
	count := 0
	for _, item := range items {
		if item.SKU == r.SKU {
			count++
		}
	}

	if count == 0 {
		return 0
	}

	price := r.RegularPrice
	if count > r.Threshold {
		price = r.Discounted
	}
	return float64(count) * price
}
