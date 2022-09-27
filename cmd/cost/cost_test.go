package main

// Only the example Testcase
// Package level test cases added at package level
// Command line inputs and parsing cases can be added, ignored for the scope.
func ExampleCoster() {

	boxParams := []costInputs{
		{
			packageID: "PKG1",
			weight:    5,
			distance:  5,
			offercode: "OFR001",
		},
		{
			packageID: "PKG2",
			weight:    15,
			distance:  5,
			offercode: "OFR002",
		},
		{
			packageID: "PKG3",
			weight:    10,
			distance:  100,
			offercode: "OFR003",
		},
	}

	coster(1, 100, boxParams)

	// Output:
	// PKG1 $0.00 $175.00
	// PKG2 $0.00 $275.00
	// PKG3 $35.00 $665.00
}
