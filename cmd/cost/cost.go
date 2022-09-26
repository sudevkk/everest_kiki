package main

import (
	"flag"
	"fmt"

	"github.com/Rhymond/go-money"
	"github.com/sudevkk/everest_kiki/internal/cargo"
	"github.com/sudevkk/everest_kiki/internal/transport"
)

type costInputs struct {
	packageID string
	weight    float64
	distance  float64
	offercode string
}

func main() {

	boxParams := make([]costInputs, 0)
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

	for _, input := range boxParams {
		box := &cargo.Box{ID: input.packageID, Weight: input.weight}
		consignment := transport.NewConsignment(box)
		consignment.SetDistance(input.distance)

		baseCost := money.NewFromFloat(*paramBaseCost, money.USD)
		consignment.SetCostingStratergy(transport.StandardCosting{
			Baseprice: baseCost,
			Offercode: input.offercode,
		})
		consignment.CalcCost()
		fmt.Printf("%s %s %s \n", consignment.Box.ID, consignment.Discount.Display(), consignment.Cost.Display())
	}
}
