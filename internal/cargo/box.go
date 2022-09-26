// The package deals with the Freight collection and the Freight content itself
// This is the Raw package before it converts to the Actual Consignment
package cargo

// A Physical Package that needs to be delivered.
// Package seems to be ideal name, not using to avoid confusion
// Can contain Package's Physical data and the data collected from the Sender (Addresses, Recieved Date)
type Box struct {
	ID     string  // Package ID
	Weight float64 // Package Weight in KGs
	// From Address
	// To Address
	// Booked DateTime
	// Content Info
	// etc
}

// Returns a New [cargo.Box]
func NewBox(id string, weight float64) *Box {
	return &Box{ID: id, Weight: weight}
}
