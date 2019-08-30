package invasion

import (
	"github.com/nbio/st"
	"testing"
)

func TestBigWorldSync(t *testing.T) {

	mainInvasion := New("../maps/big.map", 3)
	mainInvasion.RunInvasionSync( 16 )

	st.Expect(t, mainInvasion.GetWorld().NumberOfCities(), 16)
	st.Expect(t, mainInvasion.GetAliens().NumberOfAliensAlive(), 1)
}

func TestSmallWorldSync(t *testing.T) {

	mainInvasion := New("../maps/small.map", 2)
	mainInvasion.RunInvasionSync( 3 )

	st.Expect(t, mainInvasion.GetWorld().NumberOfCities(), 5)
	st.Expect(t, mainInvasion.GetAliens().NumberOfAliensAlive(), 0)
}

func TestBigWorldAsync(t *testing.T) {

	mainInvasion := New("../maps/big.map", 3)
	mainInvasion.RunInvasionAsync( 16 )

	st.Expect(t, mainInvasion.GetWorld().NumberOfCities(), 16)
	st.Expect(t, mainInvasion.GetAliens().NumberOfAliensAlive(), 1)
}

func TestSmallWorldAsync(t *testing.T) {

	mainInvasion := New("../maps/small.map", 2)
	mainInvasion.RunInvasionAsync( 3 )

	st.Expect(t, mainInvasion.GetWorld().NumberOfCities(), 5)
	st.Expect(t, mainInvasion.GetAliens().NumberOfAliensAlive(), 0)
}

