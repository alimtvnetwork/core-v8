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
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/mutexbykey"
)

func Test_Get_Verification(t *testing.T) {
	for caseIndex, testCase := range getAndDeleteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")

		// Act
		mutex := mutexbykey.Get(key)
		isNotNil := fmt.Sprintf("%v", mutex != nil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isNotNil)

		// Cleanup
		mutexbykey.Delete(key)
	}
}

func Test_Delete_Verification(t *testing.T) {
	for caseIndex, testCase := range deleteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		key, _ := input.GetAsString("key")

		// pre-create only for the "existing key" case
		when := input.When()
		if fmt.Sprintf("%v", when) == "given existing key to delete" {
			mutexbykey.Get(key)
		}

		// Act
		deleted := mutexbykey.Delete(key)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", deleted))
	}
}
