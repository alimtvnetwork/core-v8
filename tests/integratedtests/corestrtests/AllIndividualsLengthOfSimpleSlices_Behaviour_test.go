package corestrtests

import (
	"errors"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ══════════════════════════════════════════════════════════════════════════════

func Test_AllIndividualsLengthOfSimpleSlices_Nil_FromSeg1(t *testing.T) {
	safeTest(t, "Test_AllIndividualsLengthOfSimpleSlices_Nil_FromSeg1", func() {
		// Act
		actual := args.Map{"len": corestr.AllIndividualsLengthOfSimpleSlices()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns 0 -- no args", actual)
	})
}

func Test_AllIndividualsLengthOfSimpleSlices_WithItems_FromSeg1(t *testing.T) {
	safeTest(t, "Test_AllIndividualsLengthOfSimpleSlices_WithItems_FromSeg1", func() {
		// Arrange
		s1 := corestr.SimpleSlice{"a", "b"}
		s2 := corestr.SimpleSlice{"c"}

		// Act
		actual := args.Map{"len": corestr.AllIndividualsLengthOfSimpleSlices(&s1, &s2)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns total -- multi slices", actual)
	})
}

