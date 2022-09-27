package main

// Only the example Testcase
// Package level test cases added at package level
// Command line inputs and parsing cases can be added, ignored for the scope.
func ExampleCosterAndTimer() {
	/*
			PKG1 50 30 OFR001

		PKG2 75 125 OFFR0008

		PKG3 175 100 OFFR003
		PKG4 110 60 OFFR002

		PKG5 155 95 NA
	*/
	boxParams := []costInputs{
		{
			packageID: "PKG1",
			weight:    50,
			distance:  50,
			offercode: "OFR001",
		},
		{
			packageID: "PKG2",
			weight:    75,
			distance:  125,
			offercode: "OFR008",
		},
		{
			packageID: "PKG3",
			weight:    175,
			distance:  100,
			offercode: "OFR003",
		},
		{
			packageID: "PKG4",
			weight:    110,
			distance:  60,
			offercode: "OFR002",
		},
		{
			packageID: "PKG5",
			weight:    155,
			distance:  95,
			offercode: "NA",
		},
	}

	var vehicleParams timerInputs = timerInputs{noOfVehicles: 2, maxSpeed: 70, maxCarriableWeight: 200}

	costerAndTimer(100, boxParams, vehicleParams)
	// Output:
	// PKG1 $0.00 $750.00 3h98m0s
	// PKG2 $0.00 $1475.00 1h78m0s
	// PKG3 $0.00 $2350.00 1h42m0s
	// PKG4 $105.00 $1395.00 0h85m0s
	// PKG5 $0.00 $2125.00 4h19m0s
}
