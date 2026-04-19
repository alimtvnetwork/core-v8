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

	"github.com/alimtvnetwork/core/keymk"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage6 — appendStringsWithBaseAnyItems skip-empty branch
//
// Target: keymk/appendStringsWithBaseAnyItems.go:13-14
//   isSkipOnEmpty && item == "" → continue
//
// Exercise via CompileKeys with IsSkipEmptyEntry=true and a sub-key
// that has an empty string in its keyChains.
// ══════════════════════════════════════════════════════════════════════════════

func Test_CompileKeys_SkipsEmptyKeyChain(t *testing.T) {
	// Arrange — sub-key built with IsSkipEmptyEntry=false so "" enters keyChains
	noSkipOption := &keymk.Option{
		Joiner:           "-",
		IsSkipEmptyEntry: false,
	}
	mainKey := keymk.NewKey.Default("root", "chain1")
	subKey := keymk.NewKey.All(noSkipOption, "sub", "", "val")

	// Act — CompileKeys uses mainKey.option.IsSkipEmptyEntry (true) to filter sub keyChains
	result := mainKey.CompileKeys(subKey)

	// Assert
	convey.Convey("CompileKeys skips empty keyChain entries when IsSkipEmptyEntry is true", t, func() {
		convey.So(result, convey.ShouldContainSubstring, "sub")
		convey.So(result, convey.ShouldContainSubstring, "val")
	})
}
