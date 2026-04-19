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

package coregeneric

import "testing"

// Test_IndexAt_ReturnsNil_WhenInternalChainShorterThanLength
// covers the defensive guard at LinkedList.go:319-321 where
// the node chain ends before reaching the requested index
// (length field is inconsistent with actual chain).
func Test_IndexAt_ReturnsNil_WhenInternalChainShorterThanLength(t *testing.T) {
	// Arrange
	ll := EmptyLinkedList[string]()
	ll.Add("a")
	// Corrupt internal state: set length larger than actual chain
	ll.length = 5

	// Act
	result := ll.IndexAt(3)

	// Assert
	if result != nil {
		t.Errorf("expected nil when chain is shorter than length, got %v", result)
	}
}
