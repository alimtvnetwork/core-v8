package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args")

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValueStatus_InvalidNoMessage_FromSeg1(t *testing.T) {
	safeTest(t, "Test_ValueStatus_InvalidNoMessage_FromSeg1", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{
			"notNil": vs != nil,
			"index": vs.Index,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"index": -1,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatusNoMessage returns valid struct -- default", actual)
	})
}

func Test_ValueStatus_Invalid_FromSeg1(t *testing.T) {
	safeTest(t, "Test_ValueStatus_Invalid_FromSeg1", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("err msg")

		// Act
		actual := args.Map{
			"notNil": vs != nil,
			"index": vs.Index,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"index": -1,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatus returns valid struct -- with message", actual)
	})
}

func Test_ValueStatus_Clone_FromSeg1(t *testing.T) {
	safeTest(t, "Test_ValueStatus_Clone_FromSeg1", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("clone me")
		cloned := vs.Clone()

		// Act
		actual := args.Map{
			"notNil": cloned != nil,
			"sameIdx": cloned.Index == vs.Index,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"sameIdx": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus Clone returns copy -- same index", actual)
	})
}

