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

package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/keymk"
)

// Test_Cov5_AppendStringsWithBaseAnyItems_SkipEmpty covers
// keymk/appendStringsWithBaseAnyItems.go L13-14: skip empty entry with IsSkipEmptyEntry.
func Test_AppendStringsWithBaseAnyItems_SkipEmpty(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Create(
		keymk.JoinerOption, // IsSkipEmptyEntry=true
		"root",
	)

	// Act — pass empty string as any item; it should be skipped
	result := key.Compile("a", "", "b")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "root-a-b"}
	expected.ShouldBeEqual(t, 0, "appendStringsWithBaseAnyItems returns empty -- skip empty", actual)
}
