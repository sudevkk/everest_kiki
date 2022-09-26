package trip

import (
	"math"
	"time"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/sudevkk/everest_kiki/internal/transport"
)

// A [trip.Vehicle] that is added to a [trip.Trip]
type FleetVehicle struct {
	Vehicle
	CurrentLoadWeight float64
	NextAvailableFrom time.Time
}

// The collection of assets used for a [trip.Trip]
// Vehicles are saved as an RBT tree for simplicity
type Fleet struct {
	Vehicles             *rbt.Tree // List of available Vehicles in the Fleet (Stored as a RB Tree)
	nextAvailableVehicle int
}

// Add a [trip.Vehicle] to a [trip.Fleet]
func (f Fleet) AddVehicle(v Vehicle) {

	fleetV := FleetVehicle{
		Vehicle:           v,
		CurrentLoadWeight: 0,
		NextAvailableFrom: time.Now(),
	}
	f.Vehicles.Put(v.ID, fleetV)
}

// Returns a new [trip.Fleet]
// Accepts a list of [trip.Vehicle] as a param, and adds those to the [trip.Fleet]'s [trip.Fleet]
func NewFleet(vehicles []Vehicle) Fleet {

	var f Fleet = Fleet{
		nextAvailableVehicle: 0,
		Vehicles:             rbt.NewWith(byAvailableTimeAndCapacity),
	}

	for _, v := range vehicles {
		f.AddVehicle(v)
	}
	return f
}

// Custom comparator (sort by Available time (Immediate) and Max Capacity (Small))
// Sorts based on the Immediate Available Vehicle and for vehicles available at the same one the one with the less
// MaxLoading capacity. The vehicle that already has a load, and not exhausted the capacity gets first priority.
func byAvailableTimeAndCapacity(a, b interface{}) int {

	c := a.(FleetVehicle)
	d := b.(FleetVehicle)

	// d  = 1c = -1
	if c.CurrentLoadWeight > 0 && (c.CurrentLoadWeight < c.MaxLoadWeight) && d.CurrentLoadWeight == 0 {
		return -1
	} else if d.CurrentLoadWeight > 0 && d.CurrentLoadWeight < d.MaxLoadWeight && c.CurrentLoadWeight == 0 {
		return 1
	}

	if c.NextAvailableFrom.Equal(d.NextAvailableFrom) {
		if c.MaxLoadWeight > d.MaxLoadWeight {
			return 1
		} else {
			return 0
		}
	} else if c.NextAvailableFrom.After(d.NextAvailableFrom) {
		return 1
	} else {
		return -1
	}
}

// Returns the next available [fleet.FleetVehicle] from the [trip.Trip] for the given weight (KG)
func (f Fleet) GetAvailableVehicleNode(weight float64) (*rbt.Node, bool) {

	leftNode := f.Vehicles.Left() // Get the first entry from the (sorted/balanced) tree

	// if vehicle {
	// Iterate and fetch from the from the tree until the current weight can be added to an available vehicle
	for vehicle := leftNode.Value.(FleetVehicle); vehicle.CurrentLoadWeight+weight > vehicle.MaxLoadWeight; {
		leftNode, vehicle = leftNode.Right, leftNode.Value.(FleetVehicle)
	}
	// }

	if vehicle := leftNode.Value.(FleetVehicle); vehicle.CurrentLoadWeight+weight <= vehicle.MaxLoadWeight {
		return leftNode, true
	}
	return leftNode, false
}

// Adds a delivery consignment to the [trip.Trip]
func (f Fleet) addConsignment(c *transport.Consignment) bool {
	vehicleNode, isVehicleAvailable := f.GetAvailableVehicleNode(c.Box.Weight)

	if isVehicleAvailable {
		vehicle := vehicleNode.Value.(FleetVehicle)
		vehicle.CurrentLoadWeight = vehicle.CurrentLoadWeight + c.Box.Weight
		timeNeededForDelivery := (math.Round((c.Distance/vehicle.Speed)*100) / 100) * 60 // Time needed for the delivery in Minutes
		vehicle.NextAvailableFrom.Add(time.Hour * time.Duration(timeNeededForDelivery))
		f.Vehicles.Put(vehicle.ID, vehicle)
	}

	return false
}
