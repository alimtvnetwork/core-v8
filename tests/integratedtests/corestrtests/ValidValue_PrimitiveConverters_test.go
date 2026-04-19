// MIT License
// 
// Copyright (c) 2020–2026
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Bool/Int/Byte/Float converters + string predicates + regex/split/clone/json (from S01)
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidValue_ValueBool_ParsesTrueLiteral_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueBool_ParsesTrueLiteral_FromS01", func() {
		// Arrange
		v := corestr.NewValidValue("true")

		// Act
		actual := args.Map{"bool": v.ValueBool()}

		// Assert
		expected := args.Map{"bool": true}
		expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- true string", actual)
	})
}


func Test_ValidValue_ValueBool_EmptyReturnsFalse_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueBool_Empty", func() {
		// Arrange
		v := corestr.NewValidValue("")

		// Act
		actual := args.Map{"bool": v.ValueBool()}

		// Assert
		expected := args.Map{"bool": false}
		expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- empty string", actual)
	})
}


func Test_ValidValue_ValueBool_InvalidReturnsFalse_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueBool_InvalidReturnsFalse_FromS01", func() {
		// Arrange
		v := corestr.NewValidValue("notabool")

		// Act
		actual := args.Map{"bool": v.ValueBool()}

		// Assert
		expected := args.Map{"bool": false}
		expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- invalid string", actual)
	})
}


func Test_ValidValue_ValueInt_ParsesIntegerString_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueInt", func() {
		// Arrange
		v := corestr.NewValidValue("42")

		// Act
		actual := args.Map{
			"int": v.ValueInt(0),
			"defInt": v.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"int": 42,
			"defInt": 42,
		}
		expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- valid int", actual)
	})
}


func Test_ValidValue_ValueInt_InvalidUsesDefault_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueInt_Invalid", func() {
		// Arrange
		v := corestr.NewValidValue("abc")

		// Act
		actual := args.Map{
			"int": v.ValueInt(99),
			"defInt": v.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"int": 99,
			"defInt": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- invalid string", actual)
	})
}


func Test_ValidValue_ValueByte_ParsesByteValue_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueByte", func() {
		// Arrange
		v := corestr.NewValidValue("100")

		// Act
		actual := args.Map{
			"byte": v.ValueByte(0),
			"defByte": v.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"byte": byte(100),
			"defByte": byte(100),
		}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- valid byte", actual)
	})
}


func Test_ValidValue_ValueByte_OverflowClampsToMax_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueByte_OverflowClampsToMax_FromS01", func() {
		// Arrange
		v := corestr.NewValidValue("999")

		// Act
		actual := args.Map{"byte": v.ValueByte(5)}

		// Assert
		expected := args.Map{"byte": byte(255)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- overflow clamped to max", actual)
	})
}


func Test_ValidValue_ValueByte_NegativeClampsToZero_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueByte_Negative", func() {
		// Arrange
		v := corestr.NewValidValue("-1")

		// Act
		actual := args.Map{
			"byte": v.ValueByte(5),
			"defByte": v.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"byte": byte(0),
			"defByte": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- negative", actual)
	})
}


func Test_ValidValue_ValueByte_InvalidReturnsZero_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueByte_InvalidReturnsZero_FromS01", func() {
		// Arrange
		v := corestr.NewValidValue("abc")

		// Act
		actual := args.Map{
			"byte": v.ValueByte(5),
			"defByte": v.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"byte": byte(0),
			"defByte": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- invalid string", actual)
	})
}


func Test_ValidValue_ValueFloat64_ParsesDecimal_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueFloat64", func() {
		// Arrange
		v := corestr.NewValidValue("3.14")

		// Act
		actual := args.Map{
			"float": v.ValueFloat64(0),
			"defFloat": v.ValueDefFloat64(),
		}

		// Assert
		expected := args.Map{
			"float": 3.14,
			"defFloat": 3.14,
		}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 returns correct value -- valid float", actual)
	})
}


func Test_ValidValue_ValueFloat64_InvalidUsesDefault_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueFloat64_InvalidUsesDefault_FromS01", func() {
		// Arrange
		v := corestr.NewValidValue("xyz")

		// Act
		actual := args.Map{
			"float": v.ValueFloat64(1.5),
			"defFloat": v.ValueDefFloat64(),
		}

		// Assert
		expected := args.Map{
			"float": 1.5,
			"defFloat": float64(0),
		}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 returns correct value -- invalid string", actual)
	})
}

// ── String checks ──


func Test_ValidValue_IsEmpty_TrueOnEmptyString_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_IsEmpty", func() {
		// Arrange
		v := corestr.NewValidValue("")
		v2 := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"emptyIsEmpty": v.IsEmpty(),
			"helloIsEmpty": v2.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"emptyIsEmpty": true,
			"helloIsEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty returns correct value -- empty vs non-empty", actual)
	})
}


func Test_ValidValue_IsWhitespace_TrueForOnlyWhitespace_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_IsWhitespace", func() {
		// Arrange
		v := corestr.NewValidValue("   ")
		v2 := corestr.NewValidValue("hi")

		// Act
		actual := args.Map{
			"wsIsWs": v.IsWhitespace(),
			"hiIsWs": v2.IsWhitespace(),
		}

		// Assert
		expected := args.Map{
			"wsIsWs": true,
			"hiIsWs": false,
		}
		expected.ShouldBeEqual(t, 0, "IsWhitespace returns correct value -- whitespace vs text", actual)
	})
}


func Test_ValidValue_Trim_RemovesLeadingTrailingSpaces_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_Trim", func() {
		// Arrange
		v := corestr.NewValidValue("  hello  ")

		// Act
		actual := args.Map{"trimmed": v.Trim()}

		// Assert
		expected := args.Map{"trimmed": "hello"}
		expected.ShouldBeEqual(t, 0, "Trim returns correct value -- leading/trailing spaces", actual)
	})
}


func Test_ValidValue_HasValidNonEmpty_TrueOnlyForValidNonEmpty_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_HasValidNonEmpty", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		v2 := corestr.NewValidValue("")
		v3 := corestr.InvalidValidValue("x")

		// Act
		actual := args.Map{
			"valid": v.HasValidNonEmpty(),
			"empty": v2.HasValidNonEmpty(),
			"invalid": v3.HasValidNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"empty": false,
			"invalid": false,
		}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmpty returns correct value -- various cases", actual)
	})
}


func Test_ValidValue_HasValidNonWhitespace_TrueForVisibleText_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_HasValidNonWhitespace", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		v2 := corestr.NewValidValue("   ")

		// Act
		actual := args.Map{
			"valid": v.HasValidNonWhitespace(),
			"ws": v2.HasValidNonWhitespace(),
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"ws": false,
		}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespace returns correct value -- text vs whitespace", actual)
	})
}


func Test_ValidValue_HasSafeNonEmpty_TrueForValidNonEmpty_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_HasSafeNonEmpty", func() {
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"safe": v.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"safe": true}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty returns correct value -- valid non-empty", actual)
	})
}

// ── Is / IsAnyOf / IsContains / IsAnyContains ──


func Test_ValidValue_Is_ExactStringMatch_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_Is", func() {
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"yes": v.Is("hello"),
			"no": v.Is("world"),
		}

		// Assert
		expected := args.Map{
			"yes": true,
			"no": false,
		}
		expected.ShouldBeEqual(t, 0, "Is returns correct value -- match vs no match", actual)
	})
}


func Test_ValidValue_IsAnyOf_AcceptsAnyMatchingCandidate_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_IsAnyOf", func() {
		// Arrange
		v := corestr.NewValidValue("b")

		// Act
		actual := args.Map{
			"found": v.IsAnyOf("a", "b", "c"),
			"notFound": v.IsAnyOf("x", "y"),
			"empty": v.IsAnyOf(),
		}

		// Assert
		expected := args.Map{
			"found": true,
			"notFound": false,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "IsAnyOf returns correct value -- found, not found, empty", actual)
	})
}


func Test_ValidValue_IsContains_TrueWhenSubstringPresent_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_IsContains", func() {
		// Arrange
		v := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{
			"yes": v.IsContains("world"),
			"no": v.IsContains("xyz"),
		}

		// Assert
		expected := args.Map{
			"yes": true,
			"no": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- substring match", actual)
	})
}


func Test_ValidValue_IsAnyContains_TrueWhenAnyCandidatePresent_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_IsAnyContains", func() {
		// Arrange
		v := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{
			"found": v.IsAnyContains("xyz", "world"),
			"notFound": v.IsAnyContains("abc"),
			"empty": v.IsAnyContains(),
		}

		// Assert
		expected := args.Map{
			"found": true,
			"notFound": false,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "IsAnyContains returns correct value -- found, not found, empty", actual)
	})
}


func Test_ValidValue_IsEqualNonSensitive_CaseInsensitiveMatch_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_IsEqualNonSensitive", func() {
		// Arrange
		v := corestr.NewValidValue("Hello")

		// Act
		actual := args.Map{
			"yes": v.IsEqualNonSensitive("hello"),
			"no": v.IsEqualNonSensitive("world"),
		}

		// Assert
		expected := args.Map{
			"yes": true,
			"no": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualNonSensitive returns correct value -- case insensitive", actual)
	})
}

// ── Regex ──


func Test_ValidValue_IsRegexMatches_HandlesNilPattern_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_IsRegexMatches", func() {
		// Arrange
		v := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"matches": v.IsRegexMatches(re),
			"nilRegex": v.IsRegexMatches(nil),
		}

		// Assert
		expected := args.Map{
			"matches": true,
			"nilRegex": false,
		}
		expected.ShouldBeEqual(t, 0, "IsRegexMatches returns correct value -- match and nil", actual)
	})
}


func Test_ValidValue_RegexFindString_ReturnsFirstMatchOrEmpty_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_RegexFindString", func() {
		// Arrange
		v := corestr.NewValidValue("abc123def")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"found": v.RegexFindString(re),
			"nil": v.RegexFindString(nil),
		}

		// Assert
		expected := args.Map{
			"found": "123",
			"nil": "",
		}
		expected.ShouldBeEqual(t, 0, "RegexFindString returns correct value -- match and nil regex", actual)
	})
}


func Test_ValidValue_RegexFindAllStrings_CountsAllMatches_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_RegexFindAllStrings", func() {
		// Arrange
		v := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items := v.RegexFindAllStrings(re, -1)
		nilItems := v.RegexFindAllStrings(nil, -1)

		// Act
		actual := args.Map{
			"count": len(items),
			"nilCount": len(nilItems),
		}

		// Assert
		expected := args.Map{
			"count": 3,
			"nilCount": 0,
		}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStrings returns correct value -- matches and nil", actual)
	})
}


func Test_ValidValue_RegexFindAllStringsWithFlag_ReportsHasAnyFlag_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_RegexFindAllStringsWithFlag", func() {
		// Arrange
		v := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items, hasAny := v.RegexFindAllStringsWithFlag(re, -1)
		nilItems, nilHas := v.RegexFindAllStringsWithFlag(nil, -1)

		// Act
		actual := args.Map{
			"count": len(items),
			"hasAny": hasAny,
			"nilCount": len(nilItems),
			"nilHas": nilHas,
		}

		// Assert
		expected := args.Map{
			"count": 3,
			"hasAny": true,
			"nilCount": 0,
			"nilHas": false,
		}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStringsWithFlag returns correct value -- matches and nil", actual)
	})
}

// ── Split ──


func Test_ValidValue_Split_ByDelimiter_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_Split", func() {
		// Arrange
		v := corestr.NewValidValue("a,b,c")
		parts := v.Split(",")

		// Act
		actual := args.Map{
			"count": len(parts),
			"first": parts[0],
			"last": parts[2],
		}

		// Assert
		expected := args.Map{
			"count": 3,
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "Split returns correct value -- comma separated", actual)
	})
}


func Test_ValidValue_SplitTrimNonWhitespace_TrimsParts_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_SplitTrimNonWhitespace", func() {
		// Arrange
		v := corestr.NewValidValue("a, ,b")
		parts := v.SplitTrimNonWhitespace(",")

		// Act
		actual := args.Map{"count": len(parts)}

		// Assert
		expected := args.Map{"count": 3}
		expected.ShouldBeEqual(t, 0, "SplitTrimNonWhitespace returns correct value -- trims whitespace", actual)
	})
}

// ── ValueBytesOnce ──


func Test_ValidValue_ValueBytesOnce_CachesByteSlice_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ValueBytesOnce", func() {
		// Arrange
		v := corestr.NewValidValue("hi")
		b := v.ValueBytesOnce()
		b2 := v.ValueBytesOncePtr()

		// Act
		actual := args.Map{
			"len": len(b),
			"lenPtr": len(b2),
			"eq": string(b) == string(b2),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"lenPtr": 2,
			"eq": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueBytesOnce returns correct value -- caches bytes", actual)
	})
}

// ── Clone / Clear / Dispose / String / FullString ──


func Test_ValidValue_Clone_ProducesIndependentCopy_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_Clone", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		c := v.Clone()

		// Act
		actual := args.Map{
			"val": c.Value,
			"valid": c.IsValid,
			"samePtr": v == c,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": true,
			"samePtr": false,
		}
		expected.ShouldBeEqual(t, 0, "Clone returns correct value -- deep copy", actual)
	})
}


func Test_ValidValue_Clone_NilReceiverReturnsNil_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_Clone_Nil", func() {
		// Arrange
		var v *corestr.ValidValue
		c := v.Clone()

		// Act
		actual := args.Map{"isNil": c == nil}

		// Assert
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "Clone returns correct value -- nil receiver", actual)
	})
}


func Test_ValidValue_Clear_ResetsAllFields_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_Clear", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		v.Clear()

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
			"msg": v.Message,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": false,
			"msg": "",
		}
		expected.ShouldBeEqual(t, 0, "Clear returns correct value -- resets all fields", actual)
	})
}


func Test_ValidValue_Dispose_ClearsAllFields_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_Dispose", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		v.Dispose()

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- clears", actual)
	})
}


func Test_ValidValue_Dispose_NilReceiverNoPanic_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_Dispose_Nil", func() {
		// Arrange
		var v *corestr.ValidValue
		// Should not panic
		v.Dispose()

		// Act
		actual := args.Map{"isNil": v == nil}

		// Assert
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- nil receiver no panic", actual)
	})
}


func Test_ValidValue_String_ReturnsValueText_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_String", func() {
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"str": v.String()}

		// Assert
		expected := args.Map{"str": "hello"}
		expected.ShouldBeEqual(t, 0, "String returns correct value -- value", actual)
	})
}


func Test_ValidValue_String_NilReceiverReturnsEmpty_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_String_Nil", func() {
		// Arrange
		var v *corestr.ValidValue

		// Act
		actual := args.Map{"str": v.String()}

		// Assert
		expected := args.Map{"str": ""}
		expected.ShouldBeEqual(t, 0, "String returns correct value -- nil receiver", actual)
	})
}


func Test_ValidValue_FullString_NonEmptyForValid_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_FullString", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		s := v.FullString()

		// Act
		actual := args.Map{"notEmpty": s != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "FullString returns correct value -- non-empty", actual)
	})
}


func Test_ValidValue_FullString_NilReceiverReturnsEmpty_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_FullString_Nil", func() {
		// Arrange
		var v *corestr.ValidValue

		// Act
		actual := args.Map{"str": v.FullString()}

		// Assert
		expected := args.Map{"str": ""}
		expected.ShouldBeEqual(t, 0, "FullString returns correct value -- nil receiver", actual)
	})
}

// ── Json / Serialize ──


func Test_ValidValue_Json_ProducesNonEmptyBytes_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_Json", func() {
		// Arrange
		v := corestr.ValidValue{Value: "test", IsValid: true}
		j := v.Json()

		// Act
		actual := args.Map{"hasBytes": j.HasBytes()}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Json returns correct value -- has bytes", actual)
	})
}


func Test_ValidValue_JsonPtr_ReturnsNonNilResult_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_JsonPtr_ReturnsNonNilResult_FromS01", func() {
		// Arrange
		v := corestr.ValidValue{Value: "test", IsValid: true}
		j := v.JsonPtr()

		// Act
		actual := args.Map{"notNil": j != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns correct value -- not nil", actual)
	})
}


func Test_ValidValue_Serialize_ReturnsBytesNoError_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_Serialize", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		b, err := v.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize returns correct value -- no error", actual)
	})
}


func Test_ValidValue_ParseInjectUsingJson_RoundTripsSuccessfully_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_ParseInjectUsingJson", func() {
		// Arrange
		v := corestr.ValidValue{Value: "test", IsValid: true}
		j := v.JsonPtr()
		v2 := &corestr.ValidValue{}
		result, err := v2.ParseInjectUsingJson(j)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"val": result.Value,
			"valid": result.IsValid,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"val": "test",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns correct value -- round trip", actual)
	})
}


func Test_ValidValue_Deserialize_RoundTripsSuccessfully_FromS01(t *testing.T) {
	safeTest(t, "Test_ValidValue_Deserialize", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		var v2 corestr.ValidValue
		err := v.Deserialize(&v2)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"val": v2.Value,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"val": "hello",
		}
		expected.ShouldBeEqual(t, 0, "Deserialize returns correct value -- round trip", actual)
	})
}
