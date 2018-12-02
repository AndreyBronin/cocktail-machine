package coctailhal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeCoctailHal_Deliver(t *testing.T) {
	fch := NewFakeCoctailHal()
	for i := 1; i < 5; i++ {
		err := fch.FillFrom(i)
		assert.NoError(t, err)
	}
	c, err := fch.Deliver()
	assert.NoError(t, err)
	t.Logf(" c = %#v", c)
	t.Logf(" fch.contains = %#v", fch.contains)
}
