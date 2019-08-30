package aliens

import (
	"github.com/nbio/st"
	"testing"
)

func TestFewAliens(t *testing.T) {

	aliensInvading := New(3, 5 )

	st.Expect(t, aliensInvading.GetLocation(2), 2)
}


func TestLotsOfAliens(t *testing.T) {

	aliensInvading := New(50, 100 )

	st.Expect(t, aliensInvading.GetLocation(2), 18)
}
