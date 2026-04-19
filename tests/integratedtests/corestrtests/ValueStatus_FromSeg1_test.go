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

func Test_ValueStatus_InvalidNoMessage_FromSeg1(t *testing.T) {
	safeTest(t, "Test_ValueStatus_InvalidNoMessage_FromSeg1", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{
			"notNil": vs != nil,
			"index": vs.Index,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"index": -1,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatusNoMessage returns valid struct -- default", actual)
	})
}

func Test_ValueStatus_Invalid_FromSeg1(t *testing.T) {
	safeTest(t, "Test_ValueStatus_Invalid_FromSeg1", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("err msg")

		// Act
		actual := args.Map{
			"notNil": vs != nil,
			"index": vs.Index,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"index": -1,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatus returns valid struct -- with message", actual)
	})
}

func Test_ValueStatus_Clone_FromSeg1(t *testing.T) {
	safeTest(t, "Test_ValueStatus_Clone_FromSeg1", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("clone me")
		cloned := vs.Clone()

		// Act
		actual := args.Map{
			"notNil": cloned != nil,
			"sameIdx": cloned.Index == vs.Index,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"sameIdx": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus Clone returns copy -- same index", actual)
	})
}

