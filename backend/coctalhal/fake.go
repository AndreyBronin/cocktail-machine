package coctalhal

const atomicValue = 1.0

type FakeCoctailHal struct {
	contains []float32 // how much mug filled

}

func (ch *FakeCoctailHal) GetBottleCount() int {
	return 8
}

func (ch *FakeCoctailHal) FillFrom(bottleNum int) error {
	ch.contains[bottleNum] += atomicValue
	return nil
}

func NewCoctailHal() FakeCoctailHal {
	self := FakeCoctailHal{}
	return self
}
