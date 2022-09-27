package trip

import (
	"math"
	"strconv"
	"strings"
	"time"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/sudevkk/everest_kiki/internal/transport"
)

// A [trip.Vehicle] that is added to a [trip.Trip]
type FleetVehicle struct {
	Vehicle
	CurrentLoadWeight float64
	NextAvailableFrom time.Time
	consignments      []transport.Consignment
}

// The collection of assets used for a [trip.Trip]
// Vehicles are saved as an RBT tree for simplicity
type Fleet struct {
	Vehicles *rbt.Tree // List of available Vehicles in the Fleet (Stored as a RB Tree)
}

// Add a [trip.Vehicle] to a [trip.Fleet]
func (f Fleet) AddVehicle(v Vehicle, nextAvailableTime time.Time) {

	fleetV := FleetVehicle{
		Vehicle:           v,
		CurrentLoadWeight: 0,
		NextAvailableFrom: nextAvailableTime,
	}
	f.Vehicles.Put(strconv.Itoa(v.ID)+"-"+strconv.Itoa(int(fleetV.NextAvailableFrom.Unix()))+"-"+strconv.Itoa(int(v.MaxLoadWeight)), fleetV)
}

// Returns a new [trip.Fleet]
// Accepts a list of [trip.Vehicle] as a param, and adds those to the [trip.Fleet]'s [trip.Fleet]
func NewFleet(vehicles []Vehicle, startTime time.Time) Fleet {

	var f Fleet = Fleet{
		Vehicles: rbt.NewWith(byAvailableTimeAndCapacity),
	}

	for _, v := range vehicles {
		f.AddVehicle(v, startTime)
	}
	return f
}

// Custom comparator (sort by Available time (Immediate) and Max Capacity (Small))
// Sorts based on the Immediate Available Vehicle and for vehicles available at the same one the one with the less
// MaxLoading capacity. The vehicle that already has a load, and not exhausted the capacity gets first priority.
// TODO: Use Value for comparison and avaoid splitting etc. Ignored errors for now. Not Thread safe
func byAvailableTimeAndCapacity(a, b interface{}) int {

	c := a.(string)
	d := b.(string)

	if c == d {
		return 0
	}
	// Commented intentionally. Skipping for a simpler comparison logic for now
	/*
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
	*/

	keySplitOne := strings.Split(c, "-")
	keySplitTwo := strings.Split(d, "-")
	var availableWeightTwo = 0.0
	availableWeightOne, _ := strconv.ParseFloat(keySplitOne[2], 64)
	availableWeightTwo, _ = strconv.ParseFloat(keySplitTwo[2], 64)

	unixOne, error := strconv.Atoi(keySplitOne[1])

	if error != nil {
		unixTwo, error := strconv.Atoi(keySplitTwo[1])
		if error != nil {
			if unixOne < unixTwo && availableWeightOne > 0 {
				return -1
			} else {
				if unixOne > unixTwo && availableWeightTwo > 0 {
					return 1
				}
			}
		}
	}

	if availableWeightOne > 0 && availableWeightOne < availableWeightTwo {
		return -1
	}

	if availableWeightTwo > 0 && availableWeightTwo < availableWeightOne {
		return 1
	}

	return 1
}

// Returns the next available [fleet.FleetVehicle] from the [trip.Trip] for the given weight (KG)
func (f Fleet) GetAvailableVehicleNode(weight float64) (*rbt.Node, bool) {

	leftNode := f.Vehicles.Left() // Get the first entry from the (sorted/balanced) tree

	// if vehicle {
	// Iterate and fetch from the from the tree until the current weight can be added to an available vehicle
	for vehicle := leftNode.Value.(FleetVehicle); leftNode != nil && vehicle.CurrentLoadWeight+weight > vehicle.MaxLoadWeight; {
		if leftNode != nil {
			leftNode, vehicle = leftNode.Right, leftNode.Value.(FleetVehicle)
		}
	}
	// }

	if leftNode != nil {
		if vehicle := leftNode.Value.(FleetVehicle); vehicle.CurrentLoadWeight+weight <= vehicle.MaxLoadWeight {
			return leftNode, true
		}
	}
	return leftNode, false
}

// Adds a delivery consignment to the [trip.Trip]
func (f Fleet) addConsignment(c *transport.Consignment) bool {
	vehicleNode, isVehicleAvailable := f.GetAvailableVehicleNode(c.Box.Weight)

	if isVehicleAvailable {
		vehicle := vehicleNode.Value.(FleetVehicle)
		f.Vehicles.Remove(vehicleNode.Key)
		vehicle.CurrentLoadWeight = vehicle.CurrentLoadWeight + c.Box.Weight
		timeNeededForDelivery := (math.Round((c.Distance/vehicle.Speed)*100) / 100) * 60 // Time needed for the delivery in Minutes
		vehicle.NextAvailableFrom = vehicle.NextAvailableFrom.Add(time.Minute * time.Duration(timeNeededForDelivery))
		vehicle.consignments = append(vehicle.consignments, *c)
		availableLoad := vehicle.MaxLoadWeight - vehicle.CurrentLoadWeight
		f.Vehicles.Put(strconv.Itoa(vehicle.ID)+"-"+strconv.Itoa(int(vehicle.NextAvailableFrom.Unix()))+"-"+strconv.Itoa(int(availableLoad)), vehicle)
	}

	return true
}
