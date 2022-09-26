// Deals with [transport.Consignment] Delivery cost Calculations
// Could be part of Box package, separated to manage different Costing structures that might be possible over time
// eg. Festive Rate, Corporate/Partner plans etc. better to Isolate Box from the actual Costing scheme
// Uses Strtegy pattern to accomadate different costing stratergies
package transport

import "github.com/Rhymond/go-money"

// Interface that should be implemented by all CostingStatergies (ie aCoster)
type CostingStrategy interface {
	Calc(weight float64, distance float64) (cost *money.Money, discount *money.Money) // Calculates the Consignment transport cost based on the strategy
}
