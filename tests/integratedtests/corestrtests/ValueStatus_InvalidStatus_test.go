package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus — Invalid factory + Clone (extracted from S01)
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValueStatus_InvalidValueStatus_StoresMessageAndIndex(t *testing.T) {
	safeTest(t, "Test_InvalidValueStatus", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("bad")

		// Act
		actual := args.Map{
			"valid":   vs.ValueValid.IsValid,
			"msg":     vs.ValueValid.Message,
			"index":   vs.Index,
			"valEmpty": vs.ValueValid.Value == "",
		}

		// Assert
		expected := args.Map{
			"valid":   false,
			"msg":     "bad",
			"index":   -1,
			"valEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatus returns correct value -- with message", actual)
	})
}


func Test_ValueStatus_InvalidValueStatusNoMessage_StoresEmptyMessage(t *testing.T) {
	safeTest(t, "Test_InvalidValueStatusNoMessage", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{
			"valid": vs.ValueValid.IsValid,
			"msg": vs.ValueValid.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "",
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatusNoMessage returns correct value -- no message", actual)
	})
}


func Test_ValueStatus_Clone_ProducesIndependentCopy(t *testing.T) {
	safeTest(t, "Test_ValueStatus_Clone", func() {
		// Arrange
		vs := &corestr.ValueStatus{
			ValueValid: corestr.NewValidValue("hello"),
			Index:      5,
		}
		c := vs.Clone()

		// Act
		actual := args.Map{
			"val": c.ValueValid.Value,
			"index": c.Index,
			"samePtr": vs == c,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"index": 5,
			"samePtr": false,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus.Clone returns correct value -- deep copy", actual)
	})
}

// ── TextWithLineNumber ──

