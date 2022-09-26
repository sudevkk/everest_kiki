package transport

import (
	"github.com/Rhymond/go-money"
	"github.com/sudevkk/everest_kiki/internal/cargo"
)

// A Consignment that gets Transported
// One to One Mapping with a [cargo.Box]
// [cargo.Box] gets converted to a Consignment with Distance, Coster, Offer etc applied on it.
// ([cargo.Box] could have raw info like From and To addresses etc, which on converting to a [transport.Consignement]
//  will be mapped to the Actual Distance,  or the Weight gets converted to the Volumetric Weight etc)
type Consignment struct {
	Box       *cargo.Box   // The [carg.Box] that is part of this consignment
	Distance  float64      // The actual distance for the shipment
	Offercode string       // Offercode, if any
	Cost      *money.Money // The calculated cost for the shipment based on the selected [transport.Coster]
	Discount  *money.Money // The calculated discount based on the supplied Offercode, if any. 0 by default
	Coster    CostingStrategy
	Mode      TransportMode
	// OrderID
	// ShipmentID
	// InvoiceID
	// etc
}

// Sets the Actual Distance on the Consignment.
// Ideally could be auto calculated
func (c *Consignment) SetDistance(distance float64) {
	c.Distance = distance
}

// Sets the supplied Coupon/Offer code on the Consignment.
// TODO: Should check if the code is valid here and ignore if not?
func (c *Consignment) SetOffercode(offercode string) {
	c.Offercode = offercode
}

// Sets the supplied costing strategy on the consignment
func (c *Consignment) SetCostingStratergy(coster CostingStrategy) {
	c.Coster = coster
}

// Sets the supplied costing strategy on the consignment
func (c *Consignment) CalcCost() {
	c.Cost, c.Discount = c.Coster.Calc(c.Box.Weight, c.Distance)
}

// Creates and returns a [transport.Consignment] from [cargo.Box]
// Sets [transport.DefaultCostingStratergy] as the costing strategy
// Sets [transport.DefaultTransportMode] as the Transport Mode
func NewConsignment(b *cargo.Box) *Consignment {
	c := &Consignment{
		Box:      b,
		Coster:   DefaultCostingStratergy, // Default Costing Stratergy
		Mode:     DefaultTransportMode,    // Default Transport Mode
		Cost:     money.New(0, money.USD),
		Discount: money.New(0, money.USD),
	}

	return c
}
