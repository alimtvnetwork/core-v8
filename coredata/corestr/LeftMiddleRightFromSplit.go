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

// LeftMiddleRightFromSplit splits a string by separator into a LeftMiddleRight.
//
// Delegates to coregeneric.TripleFromSplit.
//
// Usage:
//
//	lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
//	lmr.Left   // "a"
//	lmr.Middle // "b"
//	lmr.Right  // "c"
func LeftMiddleRightFromSplit(input, sep string) *LeftMiddleRight {
	triple := coregeneric.TripleFromSplit(input, sep)

	return &LeftMiddleRight{Triple: *triple}
}

// LeftMiddleRightFromSplitTrimmed splits and trims whitespace from all parts.
//
// Delegates to coregeneric.TripleFromSplitTrimmed.
func LeftMiddleRightFromSplitTrimmed(input, sep string) *LeftMiddleRight {
	triple := coregeneric.TripleFromSplitTrimmed(input, sep)

	return &LeftMiddleRight{Triple: *triple}
}

// LeftMiddleRightFromSplitN splits into at most 3 parts.
// The third part contains the remainder after the second separator.
//
// Delegates to coregeneric.TripleFromSplitN.
//
// Usage:
//
//	lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
//	lmr.Left   // "a"
//	lmr.Middle // "b"
//	lmr.Right  // "c:d:e"
func LeftMiddleRightFromSplitN(input, sep string) *LeftMiddleRight {
	triple := coregeneric.TripleFromSplitN(input, sep)

	return &LeftMiddleRight{Triple: *triple}
}

// LeftMiddleRightFromSplitNTrimmed splits into at most 3 parts with trimming.
//
// Delegates to coregeneric.TripleFromSplitNTrimmed.
func LeftMiddleRightFromSplitNTrimmed(input, sep string) *LeftMiddleRight {
	triple := coregeneric.TripleFromSplitNTrimmed(input, sep)

	return &LeftMiddleRight{Triple: *triple}
}
