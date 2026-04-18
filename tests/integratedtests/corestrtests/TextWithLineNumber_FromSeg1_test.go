package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args")

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextWithLineNumber_HasLineNumber_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_HasLineNumber_FromSeg1", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		// Act
		actual := args.Map{
			"has": tw.HasLineNumber(),
			"invalid": tw.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"invalid": false,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber HasLineNumber true -- valid line", actual)
	})
}

func Test_TextWithLineNumber_Invalid_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_Invalid_FromSeg1", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hello"}

		// Act
		actual := args.Map{
			"has": tw.HasLineNumber(),
			"invalid": tw.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": false,
			"invalid": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsInvalidLineNumber -- invalid line", actual)
	})
}

func Test_TextWithLineNumber_Length_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_Length_FromSeg1", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}

		// Act
		actual := args.Map{"len": tw.Length()}

		// Assert
		expected := args.Map{"len": 5}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber Length -- 5 chars", actual)
	})
}

func Test_TextWithLineNumber_NilLength_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_NilLength_FromSeg1", func() {
		// Arrange
		var tw *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"len": tw.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber Length -- nil", actual)
	})
}

func Test_TextWithLineNumber_IsEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmpty_FromSeg1", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hello"}

		// Act
		actual := args.Map{"empty": tw.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmpty -- invalid line number", actual)
	})
}

func Test_TextWithLineNumber_IsEmptyText_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmptyText_FromSeg1", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}

		// Act
		actual := args.Map{
			"emptyText": tw.IsEmptyText(),
			"emptyBoth": tw.IsEmptyTextLineBoth(),
		}

		// Assert
		expected := args.Map{
			"emptyText": true,
			"emptyBoth": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmptyText -- empty text", actual)
	})
}

func Test_TextWithLineNumber_NilEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_NilEmpty_FromSeg1", func() {
		// Arrange
		var tw *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"empty": tw.IsEmpty(),
			"emptyText": tw.IsEmptyText(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"emptyText": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmpty -- nil receiver", actual)
	})
}

