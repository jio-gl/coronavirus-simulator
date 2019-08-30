package world

import (
	"github.com/nbio/st"
	"testing"
)

func TestSmallWorld(t *testing.T) {
	worldAttacked := LoadWorld("../maps/small.map")

	st.Expect(t, worldAttacked.NumberOfCities(), 6)
	st.Expect(t, worldAttacked.CityName(0), "Foo")
}

func TestBigWorld(t *testing.T) {
	worldAttacked := LoadWorld("../maps/big.map")

	st.Expect(t, worldAttacked.NumberOfCities(), 17)
	st.Expect(t, worldAttacked.CityName(0), "Roma")
}


