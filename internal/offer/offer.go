// Package deals with the Offers/Discount Coupons etc
package offer

import (
	"github.com/Rhymond/go-money"
)

// Represents an Offer
type offer struct {
	Code               string  // The actual offer/coupon code
	DiscountPercentage float64 // The Discount Percentage that should be applied when the coupon is added (Assumes whole value perc for now)
	MinDistance        float64 // The minimum distance, needed for the code to be eligible in a Consignment
	MaxDistance        float64 // The maximum distance, needed for the code to be eligible in a Consignment
	MinBoxWeight       float64 // The minimum weight, needed for the code to be eligible in a Consignment
	MaxBoxWeight       float64 // The maximum weight, needed for the code to be eligible in a Consignment
}

// A Global for holding the available list of Offercodes and details as a Key-Value pair
// Assuming only initialized once during init and only reads post that, (Thus, though [offers] is not threadsafe,
// is not a problem here).
// # TODO::

// The actual implementation will involve persisting the Offer/Coupon data in DB etc.
// And also since this information will be something that will get fetched frequently (On each consignment addition),
// to reduce the actual DB reads this can be cached.
// Possible possible caching stratergy: Use a Redis or some in-memory store to save the Coupon/Offer data.
// On application bootstrap the DB can be read and all the info can be Written to Redis (Offercode as Key)
// Will need a Mechanism/Service to check and invalidate codes from Redis as needed based on diff criterias (Code Expiry, Max allowed usage
// exhausted for a code etc)
// For simplicity the persistance/caching is not done here, instead a simple in memory list of Maps holds the list of Codes that will be used directly.
var offers map[string]*offer

func init() {
	// TODO: Bootstrapping: Read offers from DB and setup in-memory lookup cache
	// Setting up inline Offers map for now
	offers = map[string]*offer{
		"OFR001": {Code: "OFR001", DiscountPercentage: 10, MinDistance: 0, MaxDistance: 200, MinBoxWeight: 70, MaxBoxWeight: 200},
		"OFR002": {Code: "OFR002", DiscountPercentage: 7, MinDistance: 50, MaxDistance: 150, MinBoxWeight: 100, MaxBoxWeight: 250},
		"OFR003": {Code: "OFR003", DiscountPercentage: 5, MinDistance: 50, MaxDistance: 250, MinBoxWeight: 10, MaxBoxWeight: 150},
	}
}

// Returns a Discount Percentage, for the supplied offer code
// Returns 0, false if no qualifying offer code was found
func DiscountPercByCode(offerCode string, d float64, weight float64) (discountPerc float64, isValidOffercode bool) {
	if offers != nil {
		offer, ok := offers[offerCode]
		if ok && (d >= offer.MinDistance && d <= offer.MaxDistance && weight >= offer.MinBoxWeight && weight <= offer.MaxBoxWeight) {
			return offer.DiscountPercentage, ok
		}
	}

	return 0, false
}

// Returns a applicable Discount amount (Assuming only Dollar as currency for now), for the supplied offer code
// Returns 0, false if no qualifying offer code was found
// TODO: Make this currency flexible.
func DiscountByCode(offerCode string, d float64, weight float64, amount *money.Money) (*money.Money, bool) {
	var discount = money.New(0, money.USD)
	discountPerc, isValidOffercode := DiscountPercByCode(offerCode, d, weight)
	if isValidOffercode {
		discount = money.NewFromFloat(amount.AsMajorUnits()*discountPerc/100, money.USD)
	}
	return discount, isValidOffercode
}
