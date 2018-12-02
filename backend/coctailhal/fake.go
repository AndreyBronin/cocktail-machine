package coctailhal

import (
	"errors"
)

const bottleCount = 8
const atomicValue = 1.0

type FakeCoctailHal struct {
	// liquid count in glass
	contains Contains // how much mug filled

}

func (ch *FakeCoctailHal) Deliver() (Contains, error) {
	defer func() {
		ch.contains = make(Contains, bottleCount)
	}()
	return ch.contains, nil
}

func (ch *FakeCoctailHal) GetBottleCount() int {
	return bottleCount
}

func (ch *FakeCoctailHal) FillFrom(bottleNum int) error {
	if bottleNum >= bottleCount {
		return errors.New("no such bottle")
	}
	ch.contains[bottleNum] += atomicValue
	return nil
}

func NewFakeCoctailHal() FakeCoctailHal {
	self := FakeCoctailHal{
		contains: make(Contains, bottleCount),
	}
	return self
}
