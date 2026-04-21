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

package mutexbykeytests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/mutexbykey"
)

// ============================================================================
// Get same key twice returns same mutex
// ============================================================================

func Test_Get_SameKeyTwice_Ext_Verification(t *testing.T) {
	for caseIndex, tc := range extGetSameKeyTwiceTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")

		// Act
		mutex1 := mutexbykey.Get(key)
		mutex2 := mutexbykey.Get(key)

		actual := args.Map{
			"isSame": mutex1 == mutex2,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)

		// Cleanup
		mutexbykey.Delete(key)
	}
}
