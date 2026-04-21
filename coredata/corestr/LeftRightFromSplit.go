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

package corestr

import "github.com/alimtvnetwork/core-v8/coredata/coregeneric"

// LeftRightFromSplit splits a string by separator into a LeftRight.
//
// Delegates to coregeneric.PairFromSplit.
//
// Usage:
//
//	lr := corestr.LeftRightFromSplit("key=value", "=")
//	lr.Left  // "key"
//	lr.Right // "value"
func LeftRightFromSplit(input, sep string) *LeftRight {
	pair := coregeneric.PairFromSplit(input, sep)

	return &LeftRight{Pair: *pair}
}

// LeftRightFromSplitTrimmed splits and trims whitespace from both parts.
//
// Delegates to coregeneric.PairFromSplitTrimmed.
func LeftRightFromSplitTrimmed(input, sep string) *LeftRight {
	pair := coregeneric.PairFromSplitTrimmed(input, sep)

	return &LeftRight{Pair: *pair}
}

// LeftRightFromSplitFull splits at the first separator only.
// Right gets everything after the first separator.
//
// Delegates to coregeneric.PairFromSplitFull.
//
// Usage:
//
//	lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
//	lr.Left  // "a"
//	lr.Right // "b:c:d"
func LeftRightFromSplitFull(input, sep string) *LeftRight {
	pair := coregeneric.PairFromSplitFull(input, sep)

	return &LeftRight{Pair: *pair}
}

// LeftRightFromSplitFullTrimmed is LeftRightFromSplitFull with trimmed output.
//
// Delegates to coregeneric.PairFromSplitFullTrimmed.
func LeftRightFromSplitFullTrimmed(input, sep string) *LeftRight {
	pair := coregeneric.PairFromSplitFullTrimmed(input, sep)

	return &LeftRight{Pair: *pair}
}
