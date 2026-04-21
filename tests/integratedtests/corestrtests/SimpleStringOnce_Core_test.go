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

	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/issetter"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ========================================
// S17: SimpleStringOnce core methods
//   Value, Set, Get, numeric conversions,
//   Boolean, IsSetter, comparison, state
// ========================================

func Test_SimpleStringOnce_Value_IsInitialized(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Value_IsInitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act & Assert
		actual := args.Map{"result": sso.Value() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'hello', got ''", actual)
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected initialized", actual)
	})
}

func Test_SimpleStringOnce_IsDefined_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsDefined", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		actual := args.Map{"result": sso.IsDefined()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected defined", actual)
	})
}

func Test_SimpleStringOnce_IsUninitialized_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsUninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act & Assert
		actual := args.Map{"result": sso.IsUninitialized()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

func Test_SimpleStringOnce_Invalidate_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Invalidate", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("val")

		// Act
		sso.Invalidate()

		// Assert
		actual := args.Map{"result": sso.IsInitialized() || sso.Value() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalidated", actual)
	})
}

func Test_SimpleStringOnce_Reset_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Reset", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("val")

		// Act
		sso.Reset()

		// Assert
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected reset", actual)
	})
}

func Test_SimpleStringOnce_IsInvalid_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsInvalid", func() {
		// Arrange
		uninit := corestr.New.SimpleStringOnce.Empty()
		initEmpty := corestr.New.SimpleStringOnce.Init("")
		valid := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		actual := args.Map{"result": uninit.IsInvalid()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected invalid for uninitialized", actual)
		actual = args.Map{"result": initEmpty.IsInvalid()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected invalid for empty value", actual)
		actual = args.Map{"result": valid.IsInvalid()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid", actual)
	})
}

func Test_SimpleStringOnce_IsInvalid_Nil(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsInvalid_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act & Assert
		actual := args.Map{"result": sso.IsInvalid()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected invalid for nil", actual)
	})
}

func Test_SimpleStringOnce_ValueBytes_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ValueBytes", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		result := sso.ValueBytes()

		// Assert
		actual := args.Map{"result": string(result) != "abc"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "bytes mismatch", actual)
	})
}

func Test_SimpleStringOnce_ValueBytesPtr(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ValueBytesPtr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("xyz")

		// Act
		result := sso.ValueBytesPtr()

		// Assert
		actual := args.Map{"result": string(result) != "xyz"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "bytes mismatch", actual)
	})
}

func Test_SimpleStringOnce_SetOnUninitialized_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_SetOnUninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		err := sso.SetOnUninitialized("val")

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
		actual = args.Map{"result": sso.Value() != "val"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "value not set", actual)
	})
}

func Test_SimpleStringOnce_SetOnUninitialized_AlreadyInit_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_SetOnUninitialized_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("existing")

		// Act
		err := sso.SetOnUninitialized("new")

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for already initialized", actual)
		actual = args.Map{"result": sso.Value() != "existing"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "value should not change", actual)
	})
}

func Test_SimpleStringOnce_GetSetOnce_Uninitialized(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_GetSetOnce_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		result := sso.GetSetOnce("first")

		// Assert
		actual := args.Map{"result": result != "first"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'first', got ''", actual)
	})
}

func Test_SimpleStringOnce_GetSetOnce_AlreadyInit_Core(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_GetSetOnce_AlreadyInit_Core", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("existing")

		// Act
		result := sso.GetSetOnce("new")

		// Assert
		actual := args.Map{"result": result != "existing"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'existing', got ''", actual)
	})
}

func Test_SimpleStringOnce_GetOnce_Uninitialized(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_GetOnce_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		result := sso.GetOnce()

		// Assert
		actual := args.Map{"result": result != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty string", actual)
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be initialized after GetOnce", actual)
	})
}

func Test_SimpleStringOnce_GetOnce_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_GetOnce_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("val")

		// Act
		result := sso.GetOnce()

		// Assert
		actual := args.Map{"result": result != "val"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'val'", actual)
	})
}

func Test_SimpleStringOnce_GetOnceFunc_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_GetOnceFunc", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		result := sso.GetOnceFunc(func() string { return "computed" })

		// Assert
		actual := args.Map{"result": result != "computed"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'computed', got ''", actual)
	})
}

func Test_SimpleStringOnce_GetOnceFunc_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_GetOnceFunc_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("existing")

		// Act
		result := sso.GetOnceFunc(func() string { return "new" })

		// Assert
		actual := args.Map{"result": result != "existing"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'existing'", actual)
	})
}

func Test_SimpleStringOnce_SetOnceIfUninitialized_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_SetOnceIfUninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		isSet := sso.SetOnceIfUninitialized("val")

		// Assert
		actual := args.Map{"result": isSet}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SimpleStringOnce_SetOnceIfUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_SetOnceIfUninitialized_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act
		isSet := sso.SetOnceIfUninitialized("new")

		// Assert
		actual := args.Map{"result": isSet}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleStringOnce_SetInitialize_SetUnInit(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_SetInitialize_SetUnInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		sso.SetInitialize()

		// Assert
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected initialized", actual)

		// Act
		sso.SetUnInit()

		// Assert
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

func Test_SimpleStringOnce_ConcatNew_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ConcatNew", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		result := sso.ConcatNew(" world")

		// Assert
		actual := args.Map{"result": result.Value() != "hello world"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'hello world', got ''", actual)
	})
}

func Test_SimpleStringOnce_ConcatNewUsingStrings_Core(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ConcatNewUsingStrings_Core", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a")

		// Act
		result := sso.ConcatNewUsingStrings("-", "b", "c")

		// Assert
		actual := args.Map{"result": result.Value() != "a-b-c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a-b-c', got ''", actual)
	})
}

func Test_SimpleStringOnce_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsEmpty_IsWhitespace", func() {
		// Arrange
		empty := corestr.New.SimpleStringOnce.Init("")
		ws := corestr.New.SimpleStringOnce.Init("  ")
		val := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		actual := args.Map{"result": empty.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ws.IsWhitespace()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected whitespace", actual)
		actual = args.Map{"result": val.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_SimpleStringOnce_Trim(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Trim", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init(" hello ")

		// Act & Assert
		actual := args.Map{"result": sso.Trim() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "trim mismatch", actual)
	})
}

func Test_SimpleStringOnce_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_HasValidNonEmpty", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("x")
		empty := corestr.New.SimpleStringOnce.Init("")

		// Act & Assert
		actual := args.Map{"result": valid.HasValidNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": empty.HasValidNonEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleStringOnce_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_HasValidNonWhitespace", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("x")
		ws := corestr.New.SimpleStringOnce.Init("  ")

		// Act & Assert
		actual := args.Map{"result": valid.HasValidNonWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ws.HasValidNonWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleStringOnce_IsValueBool(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsValueBool", func() {
		// Arrange
		ssoFalse := corestr.New.SimpleStringOnce.Init("false")
		ssoTrue := corestr.New.SimpleStringOnce.Init("true")

		// Act & Assert
		actual := args.Map{"result": ssoFalse.IsValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": ssoTrue.IsValueBool()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SimpleStringOnce_SafeValue_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_SafeValue", func() {
		// Arrange
		init := corestr.New.SimpleStringOnce.Init("val")
		uninit := corestr.New.SimpleStringOnce.Empty()

		// Act & Assert
		actual := args.Map{"result": init.SafeValue() != "val"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'val'", actual)
		actual = args.Map{"result": uninit.SafeValue() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for uninitialized", actual)
	})
}

func Test_SimpleStringOnce_Int_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Int", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("42")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": sso.Int() != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		actual = args.Map{"result": invalid.Int() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleStringOnce_Byte_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Byte", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("200")
		overflow := corestr.New.SimpleStringOnce.Init("300")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": valid.Byte() != 200}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 200", actual)
		actual = args.Map{"result": overflow.Byte() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for overflow", actual)
		actual = args.Map{"result": invalid.Byte() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for invalid", actual)
	})
}

func Test_SimpleStringOnce_Int16_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Int16", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("1000")
		overflow := corestr.New.SimpleStringOnce.Init("40000")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": valid.Int16() != 1000}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1000", actual)
		actual = args.Map{"result": overflow.Int16() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for overflow", actual)
		actual = args.Map{"result": invalid.Int16() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for invalid", actual)
	})
}

func Test_SimpleStringOnce_Int32_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Int32", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("100000")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": valid.Int32() != 100000}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100000", actual)
		actual = args.Map{"result": invalid.Int32() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleStringOnce_Uint16_Core(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Uint16_Core", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("500")

		// Act
		val, inRange := valid.Uint16()

		// Assert
		actual := args.Map{"result": inRange || val != 500}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 500 in range, got", actual)
	})
}

func Test_SimpleStringOnce_Uint32_Core(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Uint32_Core", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("100000")

		// Act
		val, inRange := valid.Uint32()

		// Assert
		actual := args.Map{"result": inRange || val != 100000}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 100000", actual)
	})
}

func Test_SimpleStringOnce_WithinRange_InRange_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_WithinRange_InRange", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("50")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		actual := args.Map{"result": inRange || val != 50}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 50 in range", actual)
	})
}

func Test_SimpleStringOnce_WithinRange_BelowMin_WithBoundary(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_WithinRange_BelowMin_WithBoundary", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("-5")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		actual := args.Map{"result": inRange}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected out of range", actual)
		actual = args.Map{"result": val != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected boundary min 0", actual)
	})
}

func Test_SimpleStringOnce_WithinRange_AboveMax_WithBoundary(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_WithinRange_AboveMax_WithBoundary", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("200")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		actual := args.Map{"result": inRange}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected out of range", actual)
		actual = args.Map{"result": val != 100}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected boundary max 100", actual)
	})
}

func Test_SimpleStringOnce_WithinRange_NoBoundary_Core(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_WithinRange_NoBoundary_Core", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("200")

		// Act
		val, inRange := sso.WithinRange(false, 0, 100)

		// Assert
		actual := args.Map{"result": inRange}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected out of range", actual)
		actual = args.Map{"result": val != 200}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected raw value 200", actual)
	})
}

func Test_SimpleStringOnce_WithinRange_Invalid(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_WithinRange_Invalid", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		actual := args.Map{"result": inRange}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for invalid", actual)
		actual = args.Map{"result": val != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleStringOnce_WithinRangeDefault(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_WithinRangeDefault", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("50")

		// Act
		val, inRange := sso.WithinRangeDefault(0, 100)

		// Assert
		actual := args.Map{"result": inRange || val != 50}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 50 in range", actual)
	})
}

func Test_SimpleStringOnce_Boolean_True_Values(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Boolean_True_Values", func() {
		// Arrange
		tests := []string{"true", "yes", "y", "1", "YES", "Y"}

		for _, v := range tests {
			sso := corestr.New.SimpleStringOnce.Init(v)

			// Act & Assert
			actual := args.Map{"result": sso.Boolean(false)}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected true for ''", actual)
		}
	})
}

func Test_SimpleStringOnce_Boolean_False_Values(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Boolean_False_Values", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("false")

		// Act & Assert
		actual := args.Map{"result": sso.Boolean(false)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleStringOnce_Boolean_Invalid(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Boolean_Invalid", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("xyz")

		// Act & Assert
		actual := args.Map{"result": sso.Boolean(false)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for invalid", actual)
	})
}

func Test_SimpleStringOnce_Boolean_ConsiderInit_Uninitialized(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Boolean_ConsiderInit_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Uninitialized("true")

		// Act & Assert
		actual := args.Map{"result": sso.Boolean(true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for uninitialized with considerInit", actual)
	})
}

func Test_SimpleStringOnce_BooleanDefault_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_BooleanDefault", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("yes")

		// Act & Assert
		actual := args.Map{"result": sso.BooleanDefault()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SimpleStringOnce_IsSetter_True(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsSetter_True", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("yes")

		// Act
		result := sso.IsSetter(false)

		// Assert
		actual := args.Map{"result": result != issetter.True}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected True", actual)
	})
}

func Test_SimpleStringOnce_IsSetter_False(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsSetter_False", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("false")

		// Act
		result := sso.IsSetter(false)

		// Assert
		actual := args.Map{"result": result != issetter.False}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected False", actual)
	})
}

func Test_SimpleStringOnce_IsSetter_Invalid(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsSetter_Invalid", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("xyz")

		// Act
		result := sso.IsSetter(false)

		// Assert
		actual := args.Map{"result": result != issetter.Uninitialized}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected Uninitialized", actual)
	})
}

func Test_SimpleStringOnce_IsSetter_ConsiderInit_Uninitialized(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsSetter_ConsiderInit_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Uninitialized("true")

		// Act
		result := sso.IsSetter(true)

		// Assert
		actual := args.Map{"result": result != issetter.False}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected False", actual)
	})
}

func Test_SimpleStringOnce_ValueInt_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ValueInt", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("42")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": sso.ValueInt(0) != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		actual = args.Map{"result": invalid.ValueInt(99) != 99}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 99", actual)
	})
}

func Test_SimpleStringOnce_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ValueDefInt", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("10")
		invalid := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		actual := args.Map{"result": sso.ValueDefInt() != 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
		actual = args.Map{"result": invalid.ValueDefInt() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleStringOnce_ValueByte(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ValueByte", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("100")
		overflow := corestr.New.SimpleStringOnce.Init("300")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": sso.ValueByte(0) != 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		actual = args.Map{"result": overflow.ValueByte(5) != 5}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5 for overflow", actual)
		actual = args.Map{"result": invalid.ValueByte(7) != 7}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 7 for invalid", actual)
	})
}

func Test_SimpleStringOnce_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ValueDefByte", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("50")
		overflow := corestr.New.SimpleStringOnce.Init("999")

		// Act & Assert
		actual := args.Map{"result": sso.ValueDefByte() != 50}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 50", actual)
		actual = args.Map{"result": overflow.ValueDefByte() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleStringOnce_ValueFloat64_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ValueFloat64", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("3.14")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": sso.ValueFloat64(0) != 3.14}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		actual = args.Map{"result": invalid.ValueFloat64(1.5) != 1.5}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1.5", actual)
	})
}

func Test_SimpleStringOnce_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ValueDefFloat64", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("2.5")

		// Act & Assert
		actual := args.Map{"result": sso.ValueDefFloat64() != 2.5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	})
}

func Test_SimpleStringOnce_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_NonPtr_Ptr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act
		nonPtr := sso.NonPtr()
		ptr := sso.Ptr()

		// Assert
		actual := args.Map{"result": nonPtr.Value() != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nonPtr mismatch", actual)
		actual = args.Map{"result": ptr == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ptr nil", actual)
	})
}

func Test_SimpleStringOnce_HasSafeNonEmpty_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_HasSafeNonEmpty", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("x")
		empty := corestr.New.SimpleStringOnce.Init("")

		// Act & Assert
		actual := args.Map{"result": valid.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": empty.HasSafeNonEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleStringOnce_Is_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Is", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act & Assert
		actual := args.Map{"result": sso.Is("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.Is("world")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleStringOnce_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsAnyOf", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("b")

		// Act & Assert
		actual := args.Map{"result": sso.IsAnyOf("a", "b", "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsAnyOf("x", "y")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleStringOnce_IsAnyOf_Empty_Core(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsAnyOf_Empty_Core", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert — empty values returns true
		actual := args.Map{"result": sso.IsAnyOf()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty values", actual)
	})
}

func Test_SimpleStringOnce_IsContains_SsoCore(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsContains", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello world")

		// Act & Assert
		actual := args.Map{"result": sso.IsContains("world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsContains("xyz")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleStringOnce_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsAnyContains", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello world")

		// Act & Assert
		actual := args.Map{"result": sso.IsAnyContains("xyz", "world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsAnyContains("abc", "def")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleStringOnce_IsAnyContains_Empty_Core(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsAnyContains_Empty_Core", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		actual := args.Map{"result": sso.IsAnyContains()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
	})
}

func Test_SimpleStringOnce_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsEqualNonSensitive", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("Hello")

		// Act & Assert
		actual := args.Map{"result": sso.IsEqualNonSensitive("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SimpleStringOnce_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_IsRegexMatches", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act & Assert
		actual := args.Map{"result": sso.IsRegexMatches(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsRegexMatches(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil regex", actual)
	})
}

func Test_SimpleStringOnce_RegexFindString(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_RegexFindString", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc123def")
		re := regexp.MustCompile(`\d+`)

		// Act
		result := sso.RegexFindString(re)

		// Assert
		actual := args.Map{"result": result != "123"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '123', got ''", actual)
	})
}

func Test_SimpleStringOnce_RegexFindString_Nil(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_RegexFindString_Nil", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": sso.RegexFindString(nil) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil regex", actual)
	})
}

func Test_SimpleStringOnce_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_RegexFindAllStringsWithFlag", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a1b2c3")
		re := regexp.MustCompile(`\d`)

		// Act
		items, hasAny := sso.RegexFindAllStringsWithFlag(re, -1)

		// Assert
		actual := args.Map{"result": hasAny || len(items) != 3}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 3 matches", actual)
	})
}

func Test_SimpleStringOnce_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_RegexFindAllStringsWithFlag_Nil", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		items, hasAny := sso.RegexFindAllStringsWithFlag(nil, -1)

		// Assert
		actual := args.Map{"result": hasAny || len(items) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil regex", actual)
	})
}

func Test_SimpleStringOnce_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_RegexFindAllStrings", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a1b2")
		re := regexp.MustCompile(`\d`)

		// Act
		items := sso.RegexFindAllStrings(re, -1)

		// Assert
		actual := args.Map{"result": len(items) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleStringOnce_RegexFindAllStrings_Nil(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_RegexFindAllStrings_Nil", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		items := sso.RegexFindAllStrings(nil, -1)

		// Assert
		actual := args.Map{"result": len(items) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}
