package stringutiltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage Gaps Iteration 24
//
// Targets remaining reachable branches in stringutil (99.8% → 100%)
// ══════════════════════════════════════════════════════════════════════════════

// ── IsStartsWith: case-sensitive, equal length, non-matching ──
// Hits basePathLength <= startsWithLength guard (line ~36)

func Test_IsStartsWith_CaseSensitive_EqualLen_NoMatch(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsStartsWith("abc", "xyz", false),
	}

	// Assert
	expected := args.Map{
		"result": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns false -- case-sensitive equal-length no match", actual)
}

func Test_IsStartsWith_CaseSensitive_EqualLen_Match(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsStartsWith("abc", "abc", false),
	}

	// Assert
	expected := args.Map{
		"result": true,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns true -- case-sensitive equal-length match", actual)
}

func Test_IsStartsWith_Empty_Content(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsStartsWith("", "abc", false),
	}

	// Assert
	expected := args.Map{
		"result": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns false -- empty content non-empty starts", actual)
}

func Test_IsStartsWith_Empty_Both(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsStartsWith("", "", false),
	}

	// Assert
	expected := args.Map{
		"result": true,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns true -- both empty", actual)
}

func Test_IsStartsWith_Longer_StartsWith(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsStartsWith("ab", "abcdef", false),
	}

	// Assert
	expected := args.Map{
		"result": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns false -- startsWith longer than content", actual)
}

func Test_IsStartsWith_IgnoreCase_EqualLen(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsStartsWith("ABC", "abc", true),
	}

	// Assert
	expected := args.Map{
		"result": true,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns true -- ignore case equal len", actual)
}

// ── IsEndsWith: case-sensitive, equal length, non-matching ──

func Test_IsEndsWith_CaseSensitive_EqualLen_NoMatch(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsEndsWith("abc", "xyz", false),
	}

	// Assert
	expected := args.Map{
		"result": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns false -- case-sensitive equal-length no match", actual)
}

func Test_IsEndsWith_Empty_Content(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsEndsWith("", "abc", false),
	}

	// Assert
	expected := args.Map{
		"result": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns false -- empty content non-empty ends", actual)
}

func Test_IsEndsWith_Empty_Both(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsEndsWith("", "", false),
	}

	// Assert
	expected := args.Map{
		"result": true,
	}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns true -- both empty", actual)
}

func Test_IsEndsWith_Longer_EndsWith(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsEndsWith("ab", "abcdef", false),
	}

	// Assert
	expected := args.Map{
		"result": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns false -- endsWith longer than content", actual)
}

func Test_IsEndsWith_IgnoreCase_EqualLen(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsEndsWith("ABC", "abc", true),
	}

	// Assert
	expected := args.Map{
		"result": true,
	}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns true -- ignore case equal len", actual)
}

// ── SafeSubstring: both -1 params ──

func Test_SafeSubstring_BothMinusOne_ReturnsEmpty(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.SafeSubstring("hello", -1, -1),
	}

	// Assert
	expected := args.Map{
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns full -- both -1", actual)
}

func Test_SafeSubstring_EmptyContent(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.SafeSubstring("", 0, 0),
	}

	// Assert
	expected := args.Map{
		"result": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns empty -- empty content", actual)
}

func Test_SafeSubstring_StartMinusOne_UsesZero(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.SafeSubstring("hello", -1, 3),
	}

	// Assert
	expected := args.Map{
		"result": "hel",
	}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns substring -- startAt -1", actual)
}

func Test_SafeSubstring_EndMinusOne_UsesLength(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.SafeSubstring("hello", 2, -1),
	}

	// Assert
	expected := args.Map{
		"result": "llo",
	}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns substring -- endingLength -1", actual)
}

func Test_SafeSubstringEnds_EmptyContent(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.SafeSubstringEnds("", 3),
	}

	// Assert
	expected := args.Map{
		"result": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds returns empty -- empty content", actual)
}

func Test_SafeSubstringEnds_MinusOne_UsesLength(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.SafeSubstringEnds("hello", -1),
	}

	// Assert
	expected := args.Map{
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds returns full -- minus one", actual)
}

func Test_SafeSubstringStarts_MinusOne_UsesZero(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.SafeSubstringStarts("hello", -1),
	}

	// Assert
	expected := args.Map{
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts returns full -- minus one", actual)
}

func Test_SafeSubstringStarts_EmptyContent(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.SafeSubstringStarts("", 2),
	}

	// Assert
	expected := args.Map{
		"result": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts returns empty -- empty content", actual)
}

// ── MaskLine / MaskTrimLine edge cases ──

func Test_MaskLine_EmptyLine_ReturnsMask(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.MaskLine("XXXXXXXX", ""),
	}

	// Assert
	expected := args.Map{
		"result": "XXXXXXXX",
	}
	expected.ShouldBeEqual(t, 0, "MaskLine returns mask -- empty line", actual)
}

func Test_MaskLine_LongerLine(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.MaskLine("XX", "hello"),
	}

	// Assert
	expected := args.Map{
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "MaskLine returns line -- line longer than mask", actual)
}

func Test_MaskLine_EmptyMask(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.MaskLine("", "hello"),
	}

	// Assert
	expected := args.Map{
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "MaskLine returns line -- empty mask", actual)
}

func Test_MaskTrimLine_EmptyAfterTrim_ReturnsMask(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.MaskTrimLine("XXXXXXXX", "   "),
	}

	// Assert
	expected := args.Map{
		"result": "XXXXXXXX",
	}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns mask -- whitespace only line", actual)
}

func Test_MaskTrimLine_LongerAfterTrim(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.MaskTrimLine("XX", " hello world "),
	}

	// Assert
	expected := args.Map{
		"result": "hello world",
	}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns trimmed -- line longer than mask", actual)
}

func Test_MaskTrimLine_EmptyMask(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.MaskTrimLine("", "hello"),
	}

	// Assert
	expected := args.Map{
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns line -- empty mask", actual)
}

// ── MaskLines / MaskTrimLines: empty slices ──

func Test_MaskLines_Empty_ReturnsNil(t *testing.T) {
	// Arrange & Act
	result := stringutil.MaskLines("XXXX")

	// Assert
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MaskLines returns empty -- no lines", actual)
}

func Test_MaskTrimLines_Empty_ReturnsNil(t *testing.T) {
	// Arrange & Act
	result := stringutil.MaskTrimLines("XXXX")

	// Assert
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines returns empty -- no lines", actual)
}

func Test_MaskTrimLines_LongerLine(t *testing.T) {
	// Arrange & Act
	result := stringutil.MaskTrimLines("XX", " hello world ")

	// Assert
	actual := args.Map{"v": result[0]}
	expected := args.Map{"v": "hello world"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines returns trimmed -- line longer than mask", actual)
}

func Test_MaskTrimLines_EmptyMask(t *testing.T) {
	// Arrange & Act
	result := stringutil.MaskTrimLines("", "hello")

	// Assert
	actual := args.Map{"v": result[0]}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines returns line -- empty mask", actual)
}

func Test_MaskLines_LongerLine(t *testing.T) {
	// Arrange & Act
	result := stringutil.MaskLines("XX", "hello")

	// Assert
	actual := args.Map{"v": result[0]}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "MaskLines returns line -- line longer than mask", actual)
}

func Test_MaskLines_EmptyMask(t *testing.T) {
	// Arrange & Act
	result := stringutil.MaskLines("", "hello")

	// Assert
	actual := args.Map{"v": result[0]}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "MaskLines returns line -- empty mask", actual)
}

// ── RemoveMany: empty content ──

func Test_RemoveMany_EmptyContent(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.RemoveMany("", "a", "b"),
	}

	// Assert
	expected := args.Map{
		"result": "",
	}
	expected.ShouldBeEqual(t, 0, "RemoveMany returns empty -- empty content", actual)
}

// ── ReplaceWhiteSpacesToSingle: empty input ──

func Test_ReplaceWhiteSpacesToSingle_Empty_ReturnsEmpty(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.ReplaceWhiteSpacesToSingle(""),
	}

	// Assert
	expected := args.Map{
		"result": "",
	}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns empty -- empty", actual)
}

func Test_ReplaceWhiteSpacesToSingle_WhitespaceOnly(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.ReplaceWhiteSpacesToSingle("   \t\n  "),
	}

	// Assert
	expected := args.Map{
		"result": "",
	}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns empty -- whitespace only", actual)
}

func Test_ReplaceWhiteSpacesToSingle_TabsNewlines(t *testing.T) {
	// Arrange & Act — \t \n \f \r are skipped entirely (continue branch)
	actual := args.Map{
		"result": stringutil.ReplaceWhiteSpacesToSingle("a\tb\nc\fd\re"),
	}

	// Assert
	expected := args.Map{
		"result": "abcde",
	}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle strips tabs and newlines -- special chars", actual)
}

// ── IsStartsAndEnds (case-sensitive variant) ──

func Test_IsStartsAndEnds_Match(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsStartsAndEnds("hello world", "hello", "world"),
	}

	// Assert
	expected := args.Map{
		"result": true,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEnds returns true -- match", actual)
}

func Test_IsStartsAndEnds_NoMatch(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.IsStartsAndEnds("hello world", "hello", "xyz"),
	}

	// Assert
	expected := args.Map{
		"result": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEnds returns false -- no match", actual)
}

// ── ToIntUsingRegexMatch: nil regex ──

func Test_ToIntUsingRegexMatch_NilRegex_ReturnsDefault(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.ToIntUsingRegexMatch(nil, "42"),
	}

	// Assert
	expected := args.Map{
		"result": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch returns 0 -- nil regex", actual)
}

// ── ToByte: negative value ──

func Test_ToByte_Negative(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.ToByte("-1", 99),
	}

	// Assert
	expected := args.Map{
		"result": byte(99),
	}
	expected.ShouldBeEqual(t, 0, "ToByte returns default -- negative", actual)
}

func Test_ToByteDefault_Negative(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.ToByteDefault("-1"),
	}

	// Assert
	expected := args.Map{
		"result": byte(0),
	}
	expected.ShouldBeEqual(t, 0, "ToByteDefault returns 0 -- negative", actual)
}

func Test_ToByteDefault_OutOfRange_ReturnsFallback(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.ToByteDefault("999"),
	}

	// Assert
	expected := args.Map{
		"result": byte(0),
	}
	expected.ShouldBeEqual(t, 0, "ToByteDefault returns 0 -- out of range", actual)
}

func Test_ToByteDefault_Invalid_ReturnsFallback(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"result": stringutil.ToByteDefault("abc"),
	}

	// Assert
	expected := args.Map{
		"result": byte(0),
	}
	expected.ShouldBeEqual(t, 0, "ToByteDefault returns 0 -- invalid", actual)
}
