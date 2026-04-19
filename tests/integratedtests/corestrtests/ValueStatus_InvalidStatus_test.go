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

