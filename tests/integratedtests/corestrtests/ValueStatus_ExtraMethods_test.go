package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus — extra method coverage (Invalid factories + Clone, from Seg8)
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValueStatus_InvalidValueStatus_FromSeg8(t *testing.T) {
	safeTest(t, "Test_ValueStatus_InvalidValueStatus_FromSeg8", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("err")

		// Act
		actual := args.Map{
			"valid": vs.ValueValid.IsValid,
			"msg": vs.ValueValid.Message,
			"idx": vs.Index,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "err",
			"idx": -1,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatus -- invalid with message", actual)
	})
}


func Test_ValueStatus_InvalidValueStatusNoMessage_FromSeg8(t *testing.T) {
	safeTest(t, "Test_ValueStatus_InvalidValueStatusNoMessage_FromSeg8", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{
			"valid": vs.ValueValid.IsValid,
			"idx": vs.Index,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"idx": -1,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatusNoMessage -- invalid", actual)
	})
}


func Test_ValueStatus_Clone_FromSeg8(t *testing.T) {
	safeTest(t, "Test_ValueStatus_Clone_FromSeg8", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("err")
		c := vs.Clone()

		// Act
		actual := args.Map{
			"msg": c.ValueValid.Message,
			"idx": c.Index,
			"diff": c != vs,
		}

		// Assert
		expected := args.Map{
			"msg": "err",
			"idx": -1,
			"diff": true,
		}
		expected.ShouldBeEqual(t, 0, "Clone -- copy", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber — Segment 8e
// ══════════════════════════════════════════════════════════════════════════════

