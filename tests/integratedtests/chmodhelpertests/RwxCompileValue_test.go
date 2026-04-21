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

package chmodhelpertests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/chmodhelper"
)

func Test_RwxCompileValue(t *testing.T) {
	for caseIndex, testCase := range rwxCompileValueTestCases {
		// Arrange
		existingRwxWrapper, _ :=
			chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
				&testCase.Existing,
			)
		expectedVariableWrapper, _ :=
			chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
				&testCase.Expected,
			)

		expectedFullRwx := expectedVariableWrapper.
			ToCompileFixedPtr().
			ToFullRwxValueString()

		// Act
		actualVarWrapper, _ :=
			chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
				&testCase.Input,
			)
		actualRwxWrapper := actualVarWrapper.
			ToCompileWrapper(existingRwxWrapper.ToCompileFixedPtr())
		actualFullRwx := actualRwxWrapper.ToFullRwxValueString()

		actLines := []string{actualFullRwx}
		expectedLines := []string{expectedFullRwx}

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			actLines,
			expectedLines,
			fmt.Sprintf("  Existing: %s", testCase.Existing.ToString(false)),
			fmt.Sprintf("  Input:    %s", testCase.Input.ToString(false)),
			fmt.Sprintf("  Expected: %s", testCase.Expected.ToString(false)),
		)
	}
}
