package main

import "testing"

func TestCheckout_MultiBuyRule(t *testing.T) {
	rule := MultiBuyRule{SKU: "atv", Required: 3, PayFor: 2, UnitPrice: 50.0}
	items := []Item{
		{SKU: "atv", Price: 50.0},
		{SKU: "atv", Price: 50.0},
		{SKU: "atv", Price: 50.0},
		{SKU: "atv", Price: 50.0},
	}

	total := rule.Apply(items)
	expected := 150.0
	if total != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, total)
	}
}

func TestCheckout_BulkDiscountRule(t *testing.T) {
	rule := BulkDiscountRule{SKU: "ipd", Threshold: 4, RegularPrice: 300, Discounted: 200}
	items := []Item{
		{SKU: "ipd", Price: 300},
		{SKU: "ipd", Price: 300},
		{SKU: "ipd", Price: 300},
		{SKU: "ipd", Price: 300},
		{SKU: "ipd", Price: 300},
	}

	total := rule.Apply(items)
	expected := 1000.0
	if total != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, total)
	}
}
func TestCheckout_MultiBuyAndBulkDiscountCombined(t *testing.T) {
	items := []Item{
		{SKU: "atv", Price: 60.0},
		{SKU: "atv", Price: 60.0},
		{SKU: "atv", Price: 60.0},
		{SKU: "ipd", Price: 300.0},
		{SKU: "ipd", Price: 300.0},
		{SKU: "ipd", Price: 300.0},
		{SKU: "ipd", Price: 300.0},
		{SKU: "ipd", Price: 300.0},
		{SKU: "vga", Price: 30},
		{SKU: "mbp", Price: 1400},
	}
	rules := []PricingRuleIface{
		MultiBuyRule{SKU: "atv", Required: 3, PayFor: 2, UnitPrice: 60.0},
		BulkDiscountRule{SKU: "ipd", Threshold: 4, RegularPrice: 300, Discounted: 200},
	}
	checkout := Checkout{pricingRules: rules}
	for _, item := range items {
		checkout.Scan(item)
	}
	total := checkout.Total()
	expected := 2550.0
	if total != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, total)
	}
}

func TestCheckout_NoRulesApplied(t *testing.T) {
	items := []Item{
		{SKU: "mbp", Price: 1400},
		{SKU: "vga", Price: 30},
	}
	checkout := Checkout{}
	for _, item := range items {
		checkout.Scan(item)
	}
	total := checkout.Total()
	expected := 1430.0 // No discounts
	if total != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, total)
	}
}

func TestCheckout_EmptyCart(t *testing.T) {
	checkout := Checkout{}
	total := checkout.Total()
	expected := 0.0
	if total != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, total)
	}
}
