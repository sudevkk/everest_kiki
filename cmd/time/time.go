/*
	 CMD Time
	   dommand line tool to generate delivery cost and times for the given number of packages

	   usage:
		time -no XX -basecost XX

		Parameters:-
			no : The number of packages to be considered
			basecost: The basecost of delivery, same for all packages

		The tool will next prompt and accepts further inputs as formatted below,

		PackageID Weight Distance Offercode

		The above will need to be entered, repeteadly based on the original input
		Post that the inputs for the Delivery time calculation as below,

		NoOfVehicles MaxWeightPerVehicle MaxWeightPerVehicle
*/
package main

import (
	"flag"
	"fmt"

	"github.com/Rhymond/go-money"
	"github.com/sudevkk/everest_kiki/internal/cargo"
	"github.com/sudevkk/everest_kiki/internal/transport"
	"github.com/sudevkk/everest_kiki/internal/trip"
)

type costInputs struct {
	packageID string
	weight    float64
	distance  float64
	offercode string
}

type timerInputs struct {
	noOfVehicles       int
	maxSpeed           float64
	maxCarriableWeight float64
}

func main() {
	boxParams := make([]costInputs, 0)
	var vehicleParams timerInputs
	paramNoOfPackages := flag.Int("no", 0, "No Of Packages")
	paramBaseCost := flag.Float64("basecost", 0, "Base cost for the delivery")

	flag.Parse()

	// TODO: Keep only CMD Parsing here and avoid the Switch and move rest to internal, to specific modules
	// Have some kind of Injection of commands from the Modules and make this dynamic than 'Switching' here.
	// TODO: Ignored model level validations for now
	for i := 0; i < *paramNoOfPackages; i++ {
		var input = costInputs{}
		fmt.Scanf("%s %f %f %s", &input.packageID, &input.weight, &input.distance, &input.offercode)
		boxParams = append(boxParams, input)
	}

	fmt.Scanf("%d %f %f", &vehicleParams.noOfVehicles, &vehicleParams.maxSpeed, &vehicleParams.maxCarriableWeight)

	costerAndTimer(*paramBaseCost, boxParams, vehicleParams)
}

func costerAndTimer(paramBaseCost float64, boxParams []costInputs, vehicleParams timerInputs) {

	var consignments []transport.Consignment
	for _, input := range boxParams {
		box := &cargo.Box{ID: input.packageID, Weight: input.weight}
		consignment := transport.NewConsignment(box)
		consignment.SetDistance(input.distance)

		baseCost := money.NewFromFloat(paramBaseCost, money.USD)
		consignment.SetCostingStratergy(transport.StandardCosting{
			Baseprice: baseCost,
			Offercode: input.offercode,
		})
		consignment.CalcCost()
		consignments = append(consignments, *consignment)
		// fmt.Printf("%s %s %s \n", consignment.Box.ID, consignment.Discount.Display(), consignment.Cost.Display())
	}

	var allVehicles []trip.Vehicle

	for i := 0; i < vehicleParams.noOfVehicles; i++ {
		allVehicles = append(allVehicles, trip.Vehicle{ID: i, MaxLoadWeight: vehicleParams.maxCarriableWeight, Speed: vehicleParams.maxSpeed})
	}

	deliveryRoute := trip.New(allVehicles)
	deliveryRoute.RunSchedule(consignments)
	deliveryRoute.PrintSchedule()
}
