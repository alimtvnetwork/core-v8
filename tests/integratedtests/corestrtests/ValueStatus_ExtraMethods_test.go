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

	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args"
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

