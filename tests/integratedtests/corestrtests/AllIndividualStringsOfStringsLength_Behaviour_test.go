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

func Test_AllIndividualStringsOfStringsLength_Nil_FromSeg1(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_Nil_FromSeg1", func() {
		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(nil)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- nil ptr", actual)
	})
}

func Test_AllIndividualStringsOfStringsLength_NilSlice_FromSeg1(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_NilSlice_FromSeg1", func() {
		// Arrange
		var s [][]string

		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- nil inner slice", actual)
	})
}

func Test_AllIndividualStringsOfStringsLength_WithItems_FromSeg1(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_WithItems_FromSeg1", func() {
		// Arrange
		s := [][]string{{"a", "b"}, {"c"}}

		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns total -- multi", actual)
	})
}

func Test_AllIndividualStringsOfStringsLength_Empty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_Empty_FromSeg1", func() {
		// Arrange
		s := [][]string{}

		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- empty outer", actual)
	})
}

