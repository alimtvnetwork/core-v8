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

package coregenerictests

import "testing"

// ==========================================
// LinkedList — CaseNilSafe pattern
// ==========================================

func Test_LinkedList_NilReceiver(t *testing.T) {
	for caseIndex, tc := range linkedListNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// Hashset — CaseNilSafe pattern
// ==========================================

func Test_Hashset_NilReceiver(t *testing.T) {
	for caseIndex, tc := range hashsetNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// Hashmap — CaseNilSafe pattern
// ==========================================

func Test_Hashmap_NilReceiver(t *testing.T) {
	for caseIndex, tc := range hashmapNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// Pair — CaseNilSafe pattern
// ==========================================

func Test_Pair_NilReceiver(t *testing.T) {
	for caseIndex, tc := range pairNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

// ==========================================
// Triple — CaseNilSafe pattern
// ==========================================

func Test_Triple_NilReceiver(t *testing.T) {
	for caseIndex, tc := range tripleNilSafeTestCases {
		// Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}
