package coctalhal

type CoctailHal interface {
	GetBottleCount() int
	// FillFrom get liquid from bottle №
	FillFrom(int) error
}
