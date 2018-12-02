package coctalhal

type Contains []float32

// CoctailHal is hardware abstraction layer for cocktail machine hardware
type CoctailHal interface {
	// GetBottleCount returns bottles count
	GetBottleCount() int
	// FillFrom get liquid from bottle №
	FillFrom(int) error
	// Deliver takes jug to a customer
	Deliver() (Contains, error)
}
