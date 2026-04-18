package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args")

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftRight_Creators_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_Creators_FromSeg1", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewLeftRight -- valid pair", actual)
	})
}

func Test_LeftRight_Invalid_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_Invalid_FromSeg1", func() {
		// Arrange
		lr := corestr.InvalidLeftRight("err")

		// Act
		actual := args.Map{
			"valid": lr.IsValid,
			"msg": lr.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "err",
		}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRight -- invalid with message", actual)
	})
}

func Test_LeftRight_InvalidNoMessage_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_InvalidNoMessage_FromSeg1", func() {
		// Arrange
		lr := corestr.InvalidLeftRightNoMessage()

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRightNoMessage -- invalid", actual)
	})
}

func Test_LeftRight_UsingSlice_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_UsingSlice_FromSeg1", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice -- 2 items valid", actual)
	})
}

func Test_LeftRight_UsingSliceEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_UsingSliceEmpty_FromSeg1", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice -- empty slice invalid", actual)
	})
}

func Test_LeftRight_UsingSliceSingle_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_UsingSliceSingle_FromSeg1", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"only"})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "only",
			"right": "",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice -- single item", actual)
	})
}

func Test_LeftRight_UsingSlicePtr_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_UsingSlicePtr_FromSeg1", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr -- delegates to UsingSlice", actual)
	})
}

func Test_LeftRight_UsingSlicePtrEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_UsingSlicePtrEmpty_FromSeg1", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlicePtr([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr -- empty", actual)
	})
}

func Test_LeftRight_TrimmedUsingSlice_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_TrimmedUsingSlice_FromSeg1", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- trimmed", actual)
	})
}

func Test_LeftRight_TrimmedUsingSliceNil_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_TrimmedUsingSliceNil_FromSeg1", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice(nil)

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- nil", actual)
	})
}

func Test_LeftRight_TrimmedUsingSliceEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_TrimmedUsingSliceEmpty_FromSeg1", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- empty", actual)
	})
}

func Test_LeftRight_TrimmedUsingSliceSingle_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_TrimmedUsingSliceSingle_FromSeg1", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" only "})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "only",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- single item not trimmed", actual)
	})
}

func Test_LeftRight_StringMethods_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_StringMethods_FromSeg1", func() {
		// Arrange
		lr := corestr.NewLeftRight(" hello ", " world ")

		// Act
		actual := args.Map{
			"leftBytes":  string(lr.LeftBytes()),
			"rightBytes": string(lr.RightBytes()),
			"leftTrim":   lr.LeftTrim(),
			"rightTrim":  lr.RightTrim(),
		}

		// Assert
		expected := args.Map{
			"leftBytes":  " hello ",
			"rightBytes": " world ",
			"leftTrim":   "hello",
			"rightTrim":  "world",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight string methods -- bytes and trim", actual)
	})
}

func Test_LeftRight_EmptyChecks_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_EmptyChecks_FromSeg1", func() {
		// Arrange
		lr := corestr.NewLeftRight("", "")

		// Act
		actual := args.Map{
			"leftEmpty":  lr.IsLeftEmpty(),
			"rightEmpty": lr.IsRightEmpty(),
			"leftWS":     lr.IsLeftWhitespace(),
			"rightWS":    lr.IsRightWhitespace(),
		}

		// Assert
		expected := args.Map{
			"leftEmpty":  true,
			"rightEmpty": true,
			"leftWS":     true,
			"rightWS":    true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight empty checks -- empty strings", actual)
	})
}

func Test_LeftRight_ValidNonEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_ValidNonEmpty_FromSeg1", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"validLeft":    lr.HasValidNonEmptyLeft(),
			"validRight":   lr.HasValidNonEmptyRight(),
			"validWSLeft":  lr.HasValidNonWhitespaceLeft(),
			"validWSRight": lr.HasValidNonWhitespaceRight(),
			"safeNonEmpty": lr.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"validLeft":    true,
			"validRight":   true,
			"validWSLeft":  true,
			"validWSRight": true,
			"safeNonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight HasValidNonEmpty -- all valid", actual)
	})
}

func Test_LeftRight_NonPtrPtr_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_NonPtrPtr_FromSeg1", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		np := lr.NonPtr()
		p := lr.Ptr()

		// Act
		actual := args.Map{
			"nonPtrLeft": np.Left,
			"ptrLeft": p.Left,
		}

		// Assert
		expected := args.Map{
			"nonPtrLeft": "a",
			"ptrLeft": "a",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight NonPtr/Ptr -- same values", actual)
	})
}

func Test_LeftRight_RegexMatch_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_RegexMatch_FromSeg1", func() {
		// Arrange
		lr := corestr.NewLeftRight("hello123", "world456")
		re := regexp.MustCompile(`[0-9]+`)

		// Act
		actual := args.Map{
			"leftMatch":  lr.IsLeftRegexMatch(re),
			"rightMatch": lr.IsRightRegexMatch(re),
			"nilRegex":   lr.IsLeftRegexMatch(nil),
		}

		// Assert
		expected := args.Map{
			"leftMatch":  true,
			"rightMatch": true,
			"nilRegex":   false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight regex match -- valid and nil regex", actual)
	})
}

func Test_LeftRight_IsComparisons_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_IsComparisons_FromSeg1", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"isLeft":  lr.IsLeft("a"),
			"isRight": lr.IsRight("b"),
			"is":      lr.Is("a", "b"),
		}

		// Assert
		expected := args.Map{
			"isLeft":  true,
			"isRight": true,
			"is":      true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight Is comparisons -- match", actual)
	})
}

func Test_LeftRight_IsEqual_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_IsEqual_FromSeg1", func() {
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("x", "y")

		// Act
		actual := args.Map{
			"eq":      lr1.IsEqual(lr2),
			"neq":     lr1.IsEqual(lr3),
			"nilBoth": (*corestr.LeftRight)(nil).IsEqual(nil),
			"nilOne":  lr1.IsEqual(nil),
		}

		// Assert
		expected := args.Map{
			"eq":      true,
			"neq":     false,
			"nilBoth": true,
			"nilOne":  false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight IsEqual -- equal, not equal, nil cases", actual)
	})
}

func Test_LeftRight_Clone_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_Clone_FromSeg1", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()

		// Act
		actual := args.Map{
			"left": c.Left,
			"right": c.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight Clone -- same values", actual)
	})
}

func Test_LeftRight_ClearDispose_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_ClearDispose_FromSeg1", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		lr.Clear()

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight Clear -- emptied", actual)
	})
}

func Test_LeftRight_DisposeNil_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LeftRight_DisposeNil_FromSeg1", func() {
		var lr *corestr.LeftRight
		lr.Dispose() // should not panic
		lr.Clear()   // should not panic
	})
}

