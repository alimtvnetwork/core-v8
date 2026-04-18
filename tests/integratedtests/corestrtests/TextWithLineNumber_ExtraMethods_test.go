package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber — extra method coverage (line/emptiness predicates, from Seg8)
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextWithLineNumber_HasLineNumber_FromSeg8(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_HasLineNumber_FromSeg8", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		// Act
		actual := args.Map{
			"has": twln.HasLineNumber(),
			"invalid": twln.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"invalid": false,
		}
		expected.ShouldBeEqual(t, 0, "HasLineNumber -- valid", actual)
	})
}


func Test_TextWithLineNumber_InvalidLineNumber_FromSeg8(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_InvalidLineNumber_FromSeg8", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hello"}

		// Act
		actual := args.Map{
			"has": twln.HasLineNumber(),
			"invalid": twln.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": false,
			"invalid": true,
		}
		expected.ShouldBeEqual(t, 0, "InvalidLineNumber -- invalid", actual)
	})
}


func Test_TextWithLineNumber_Length_FromSeg8(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_Length_FromSeg8", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 1, Text: "abc"}

		// Act
		actual := args.Map{"len": twln.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Length -- 3", actual)
	})
}


func Test_TextWithLineNumber_Length_Nil_FromSeg8(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_Length_Nil_FromSeg8", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"len": twln.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length nil -- 0", actual)
	})
}


func Test_TextWithLineNumber_IsEmpty_FromSeg8(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmpty_FromSeg8", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}

		// Act
		actual := args.Map{"empty": twln.IsEmpty()}

		// Assert
		expected := args.Map{"empty": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- false", actual)
	})
}


func Test_TextWithLineNumber_IsEmpty_True_FromSeg8(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmpty_True_FromSeg8", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"empty": twln.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true (no text + invalid line)", actual)
	})
}


func Test_TextWithLineNumber_IsEmpty_Nil_FromSeg8(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmpty_Nil_FromSeg8", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"empty": twln.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty nil -- true", actual)
	})
}


func Test_TextWithLineNumber_IsEmptyText_FromSeg8(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmptyText_FromSeg8", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}

		// Act
		actual := args.Map{"emptyText": twln.IsEmptyText()}

		// Assert
		expected := args.Map{"emptyText": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyText -- true", actual)
	})
}


func Test_TextWithLineNumber_IsEmptyText_Nil_FromSeg8(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmptyText_Nil_FromSeg8", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"emptyText": twln.IsEmptyText()}

		// Assert
		expected := args.Map{"emptyText": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyText nil -- true", actual)
	})
}


func Test_TextWithLineNumber_IsEmptyTextLineBoth_FromSeg8(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmptyTextLineBoth_FromSeg8", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"both": twln.IsEmptyTextLineBoth()}

		// Assert
		expected := args.Map{"both": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyTextLineBoth -- delegates to IsEmpty", actual)
	})
}


func Test_TextWithLineNumber_HasLineNumber_Nil_FromSeg8(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_HasLineNumber_Nil_FromSeg8", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"has": twln.HasLineNumber(),
			"invalid": twln.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": false,
			"invalid": true,
		}
		expected.ShouldBeEqual(t, 0, "HasLineNumber nil -- false/true", actual)
	})
}
