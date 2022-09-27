// Package deals with the Trip (Delivery of selected consignments)
// Manages the Scheduling etc
package trip

import (
	"fmt"
	"sort"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/sudevkk/everest_kiki/internal/transport"
)

// Represents a Trip / delivery route
// the scheduling and other actions happens on a [trip.Trip]
type Trip struct {
	fleet     Fleet
	startTime time.Time // Only Used for the simulation time for schedule
}

// Returns a new [trip.Trip]
// Accepts a list of [trip.Vehicle]
func New(v []Vehicle) Trip {
	now := time.Now()
	f := NewFleet(v, now)

	t := Trip{
		fleet:     f,
		startTime: now,
	}

	return t
}

// Runs a scheduling (simulated) for a given list of consignements on the current [trip.Trip]
// The consignments are sorted on Weight and Distance.
// The scheduler utilises the RB tree of vehicle nodes to do the assignment in priority
// The complexity (n logn * n). This can be likely optimised further and not attempted.
func (t Trip) RunSchedule(consignments []transport.Consignment) {

	sort.Slice(consignments, func(i, j int) bool {
		// TODO: Ignoring the float comparison rounding erros etc for now
		if consignments[i].Box.Weight == consignments[j].Box.Weight {
			if consignments[i].Distance > consignments[j].Distance {
				return true
			}
		} else if consignments[i].Box.Weight > consignments[j].Box.Weight {
			return true
		}
		return false
	})

	fmt.Printf("%v \n", consignments)

	for _, c := range consignments {
		t.fleet.addConsignment(&c)
	}
}

// Prints the Cost and Simulated Delivery time
func (t Trip) PrintSchedule() {
	type scheduleOutput struct {
		packageID    string
		cost         *money.Money
		deliveryTime time.Duration
	}
	var packageSchedules []scheduleOutput
	it := t.fleet.Vehicles.Iterator()
	for it.Next() {
		_, value := it.Key(), it.Value()
		fleetVehicle := value.(FleetVehicle)

		if fleetVehicle.CurrentLoadWeight > 0 {
			for _, c := range fleetVehicle.consignments {
				deliveryTime := fleetVehicle.NextAvailableFrom.Sub(t.startTime)
				packageSchedules = append(packageSchedules, scheduleOutput{packageID: c.Box.ID, cost: c.Cost, deliveryTime: deliveryTime})
			}
		}
	}

	for _, v := range packageSchedules {
		fmt.Printf("%s %s %v \n", v.packageID, v.cost.Display(), v.deliveryTime)
	}
}
