package main

import "fmt"

func main() {

	products := map[string]Item{
		"ipd": {"ipd", "Super iPad", 549.99},
		"mbp": {"mbp", "MacBook Pro", 1399.99},
		"atv": {"atv", "Apple TV", 109.50},
		"vga": {"vga", "VGA adapter", 30.00},
	}

	pricingRules := []PricingRuleIface{
		MultiBuyRule{"atv", 3, 2, 109.50},          // 3-for-2 Apple TVs
		BulkDiscountRule{"ipd", 4, 549.99, 499.99}, // Bulk discount for iPads if greater then 4 price will be 499.99 for each item
	}

	//I took this scenario from github example https://github.com/zeller-public/code-challenge-bff/tree/golang?tab=readme-ov-file#example-scenarios
	githubExampleScenario1 := Checkout{pricingRules: pricingRules}
	githubExampleScenario1.Scan(products["atv"])
	githubExampleScenario1.Scan(products["atv"])
	githubExampleScenario1.Scan(products["atv"])
	githubExampleScenario1.Scan(products["vga"])
	fmt.Printf("Total Price with VGA and Apple TVs: $%.2f\n", githubExampleScenario1.Total()) //$249.00 Output

	githubExampleScenario2 := Checkout{pricingRules: pricingRules}
	githubExampleScenario2.Scan(products["atv"])
	githubExampleScenario2.Scan(products["ipd"])
	githubExampleScenario2.Scan(products["ipd"])
	githubExampleScenario2.Scan(products["atv"])
	githubExampleScenario2.Scan(products["ipd"])
	githubExampleScenario2.Scan(products["ipd"])
	githubExampleScenario2.Scan(products["ipd"])
	fmt.Printf("Total Price with 5 iPads and Apple TVs: $%.2f\n", githubExampleScenario2.Total()) //$2718.95 Output

	//I created my own test cases
	checkout1 := Checkout{pricingRules: pricingRules}
	checkout2 := Checkout{pricingRules: pricingRules}
	checkout3 := Checkout{pricingRules: pricingRules}

	// TEST 1
	checkout1.Scan(products["atv"])
	checkout1.Scan(products["atv"])
	checkout1.Scan(products["atv"])
	checkout1.Scan(products["atv"])
	checkout1.Scan(products["atv"])
	checkout1.Scan(products["atv"])
	fmt.Printf("Total Price with Apple TVs: $%.2f\n", checkout1.Total()) //$438.00 Output
	// TEST 2
	checkout2.Scan(products["vga"])
	checkout2.Scan(products["ipd"])
	checkout2.Scan(products["ipd"])
	checkout2.Scan(products["ipd"])
	checkout2.Scan(products["ipd"])
	checkout2.Scan(products["ipd"])
	fmt.Printf("Total Price with iPads: $%.2f\n", checkout2.Total()) //$2529.95 Output

	// TEST 3
	checkout3.Scan(products["vga"])
	checkout3.Scan(products["vga"])
	checkout3.Scan(products["atv"])
	checkout3.Scan(products["atv"])
	checkout3.Scan(products["atv"])
	checkout3.Scan(products["ipd"])
	checkout3.Scan(products["ipd"])
	checkout3.Scan(products["ipd"])
	checkout3.Scan(products["ipd"])
	checkout3.Scan(products["ipd"])
	checkout3.Scan(products["mbp"])
	fmt.Printf("Total Price with Apple TVs and iPads: $%.2f\n", checkout3.Total()) //$4178.94 Output

}
