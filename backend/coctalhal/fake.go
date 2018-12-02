package coctalhal

const bottleCount = 8
const atomicValue = 1.0

type FakeCoctailHal struct {
	contains [bottleCount]float32 // how much mug filled

}

func (ch *FakeCoctailHal) GetBottleCount() int {
	return bottleCount
}

func (ch *FakeCoctailHal) FillFrom(bottleNum int) error {
	ch.contains[bottleNum] += atomicValue
	return nil
}

func NewCoctailHal() FakeCoctailHal {
	self := FakeCoctailHal{}
	return self
}
