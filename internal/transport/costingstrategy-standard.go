package transport

import (
	"github.com/Rhymond/go-money"
	"github.com/sudevkk/everest_kiki/internal/offer"
)

// The standard/default consignment cost calculation strategy
// Calculates the Cost based on the formula Delivery Cost = Base Delivery Cost + (Package Total Weight * 10) + (Distance to Destination * 5)
// Applies Offer/Coupon code if  a valid code is supplied
type StandardCosting struct {
	Baseprice *money.Money
	Offercode string
}

func (s StandardCosting) Calc(weight float64, distance float64) (cost *money.Money, discount *money.Money) {

	var err error
	var isValidCode bool

	discount = money.New(0, money.USD)

	cost = money.NewFromFloat(float64((weight*10)+(distance*5)), money.USD)
	cost, err = s.Baseprice.Add(cost)

	if err == nil {
		if s.Offercode != "" {
			discount, isValidCode = offer.DiscountByCode(s.Offercode, distance, weight, cost)
			if isValidCode {
				cost, _ = cost.Subtract(discount)
			}
		}
	}

	return cost, discount
}

var DefaultCostingStratergy = StandardCosting{}
