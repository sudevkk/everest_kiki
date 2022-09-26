package trip

// The Vehicle entity
type Vehicle struct {
	ID            int
	MaxLoadWeight float64 // The maximum weight in KGs the vehicle can carry
	Speed         float64 // Speed KM/HR (Assuming constant speed)
}
