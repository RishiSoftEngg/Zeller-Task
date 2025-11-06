package main

type Item struct {
	SKU   string
	Name  string
	Price float64
}

type Checkout struct {
	pricingRules []PricingRuleIface
	items        []Item
}

// @description
// Scans an item into the checkout.
//
// @param
// item Item - The item to scan
func (c *Checkout) Scan(item Item) {
	c.items = append(c.items, item)
}

// @description
// Calculates the total price after applying all pricing rules.
//
// @return
// float64 - The total price after applying all pricing rules
func (c *Checkout) Total() float64 {
	total := 0.0
	itemGroups := make(map[string][]Item)
	for _, item := range c.items {
		itemGroups[item.SKU] = append(itemGroups[item.SKU], item)
	}

	for sku, items := range itemGroups {
		appliedRule := false

		for _, rule := range c.pricingRules {
			if rule.AppliesTo(sku) {
				total += rule.Apply(items)
				appliedRule = true
				break
			}
		}
		if !appliedRule {
			for _, item := range items {
				total += item.Price
			}

		}
	}
	return total
}
