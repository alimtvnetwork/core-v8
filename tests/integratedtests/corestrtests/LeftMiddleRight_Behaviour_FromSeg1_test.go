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
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args")

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ══════════════════════════════════════════════════════════════════════════════

func Test_LMR_Creators_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_Creators_FromSeg1", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
			"valid": lmr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewLeftMiddleRight -- valid triple", actual)
	})
}

func Test_LMR_Invalid_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_Invalid_FromSeg1", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRight("err")

		// Act
		actual := args.Map{
			"valid": lmr.IsValid,
			"msg": lmr.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "err",
		}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRight -- invalid with message", actual)
	})
}

func Test_LMR_InvalidNoMessage_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_InvalidNoMessage_FromSeg1", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRightNoMessage()

		// Act
		actual := args.Map{"valid": lmr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRightNoMessage -- invalid", actual)
	})
}

func Test_LMR_BytesMethods_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_BytesMethods_FromSeg1", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("L", "M", "R")

		// Act
		actual := args.Map{
			"leftB":  string(lmr.LeftBytes()),
			"midB":   string(lmr.MiddleBytes()),
			"rightB": string(lmr.RightBytes()),
		}

		// Assert
		expected := args.Map{
			"leftB":  "L",
			"midB":   "M",
			"rightB": "R",
		}
		expected.ShouldBeEqual(t, 0, "LMR Bytes methods -- correct bytes", actual)
	})
}

func Test_LMR_TrimMethods_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_TrimMethods_FromSeg1", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight(" L ", " M ", " R ")

		// Act
		actual := args.Map{
			"leftTrim":  lmr.LeftTrim(),
			"midTrim":   lmr.MiddleTrim(),
			"rightTrim": lmr.RightTrim(),
		}

		// Assert
		expected := args.Map{
			"leftTrim":  "L",
			"midTrim":   "M",
			"rightTrim": "R",
		}
		expected.ShouldBeEqual(t, 0, "LMR Trim methods -- trimmed", actual)
	})
}

func Test_LMR_EmptyChecks_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_EmptyChecks_FromSeg1", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("", "", "")

		// Act
		actual := args.Map{
			"leftEmpty":  lmr.IsLeftEmpty(),
			"midEmpty":   lmr.IsMiddleEmpty(),
			"rightEmpty": lmr.IsRightEmpty(),
			"leftWS":     lmr.IsLeftWhitespace(),
			"midWS":      lmr.IsMiddleWhitespace(),
			"rightWS":    lmr.IsRightWhitespace(),
		}

		// Assert
		expected := args.Map{
			"leftEmpty":  true,
			"midEmpty":   true,
			"rightEmpty": true,
			"leftWS":     true,
			"midWS":      true,
			"rightWS":    true,
		}
		expected.ShouldBeEqual(t, 0, "LMR empty checks -- all empty", actual)
	})
}

func Test_LMR_ValidNonEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_ValidNonEmpty_FromSeg1", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"validLeft":    lmr.HasValidNonEmptyLeft(),
			"validMid":     lmr.HasValidNonEmptyMiddle(),
			"validRight":   lmr.HasValidNonEmptyRight(),
			"validWSLeft":  lmr.HasValidNonWhitespaceLeft(),
			"validWSMid":   lmr.HasValidNonWhitespaceMiddle(),
			"validWSRight": lmr.HasValidNonWhitespaceRight(),
			"safeNonEmpty": lmr.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"validLeft":    true,
			"validMid":     true,
			"validRight":   true,
			"validWSLeft":  true,
			"validWSMid":   true,
			"validWSRight": true,
			"safeNonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "LMR HasValidNonEmpty -- all valid", actual)
	})
}

func Test_LMR_IsAll_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_IsAll_FromSeg1", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"isAll": lmr.IsAll("a", "b", "c"),
			"is":    lmr.Is("a", "c"),
		}

		// Assert
		expected := args.Map{
			"isAll": true,
			"is":    true,
		}
		expected.ShouldBeEqual(t, 0, "LMR IsAll and Is -- match", actual)
	})
}

func Test_LMR_Clone_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_Clone_FromSeg1", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		c := lmr.Clone()

		// Act
		actual := args.Map{
			"left": c.Left,
			"mid": c.Middle,
			"right": c.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LMR Clone -- same values", actual)
	})
}

func Test_LMR_ToLeftRight_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_ToLeftRight_FromSeg1", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "c",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LMR ToLeftRight -- left and right preserved", actual)
	})
}

func Test_LMR_ClearDispose_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_ClearDispose_FromSeg1", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "",
			"mid": "",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "LMR Clear -- emptied", actual)
	})
}

func Test_LMR_DisposeNil_FromSeg1(t *testing.T) {
	safeTest(t, "Test_LMR_DisposeNil_FromSeg1", func() {
		var lmr *corestr.LeftMiddleRight
		lmr.Dispose() // should not panic
		lmr.Clear()   // should not panic
	})
}

