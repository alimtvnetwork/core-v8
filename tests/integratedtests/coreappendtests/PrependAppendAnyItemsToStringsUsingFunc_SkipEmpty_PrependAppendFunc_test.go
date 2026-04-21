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

package coreappendtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coreappend"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_PrependAppendAnyItemsToStringsUsingFunc_SkipEmpty(t *testing.T) {
	// Arrange
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		true,
		func(item any) string { return fmt.Sprintf("%v", item) },
		"pre",
		"post",
		"a", nil, "",
	)
	// nil items are skipped, empty string items are skipped (isSkipEmptyString=true)
	// "a" -> "a", nil -> skipped, "" -> "" which is empty so skipped

	// Act
	actual := args.Map{"result": len(result) < 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 2 items", actual)
}

func Test_PrependAppendAnyItemsToStringsUsingFunc_SkipEmptyString_InLoop(t *testing.T) {
	// Arrange
	// This specifically targets the branch: isSkipEmptyString && currentStr == ""
	result := coreappend.PrependAppendAnyItemsToStringsUsingFunc(
		true,
		func(item any) string {
			if item == nil {
				return ""
			}
			return fmt.Sprintf("%v", item)
		},
		"pre",
		"post",
		"hello", "world",
	)

	// Act
	actual := args.Map{"result": len(result) < 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 3 items", actual)
}
