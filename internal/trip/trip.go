// Package deals with the Trip (Delivery of selected consignments)
// Manages the Scheduling etc
package trip

import (
	"fmt"
	"sort"

	"github.com/sudevkk/everest_kiki/internal/transport"
)

// Represents a Trip / delivery route
// the scheduling and other actions happens on a [trip.Trip]
type Trip struct {
	fleet Fleet
}

// Returns a new [trip.Trip]
// Accepts a list of [trip.Vehicle]
func New(v []Vehicle) Trip {
	f := NewFleet(v)

	t := Trip{
		fleet: f,
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

	fmt.Printf("%v \n", t.fleet)
}
