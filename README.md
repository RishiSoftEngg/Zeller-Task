
# Zeller Coding Challenge - Checkout System in Golang

According to the problem statement, I have designed a **Checkout System with Flexible Pricing Rules** using interfaces and structs.
This design allows different pricing behaviors to be implemented without modifying the core business logic of the engine.

The code is structured to be:

    • Easily extensible — new pricing rules can be added without changing existing logic

    • Highly testable — each rule can be unit-tested independently

	• Clean and maintainable — follows Go best practices and separation of concerns


#### To execute the main program:

```
  go run .
```

#### To run all test cases:

```
  go test -v
```