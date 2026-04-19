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

package coresorttests

import "fmt"

// toIntPtrs converts []int to []*int (each element gets its own pointer).
func toIntPtrs(values []int) []*int {
	ptrs := make([]*int, len(values))

	for i := range values {
		v := values[i]
		ptrs[i] = &v
	}

	return ptrs
}

// formatIntPtrs formats []*int to "[v1 v2 v3]" matching fmt.Sprintf("%v", []int{...}).
func formatIntPtrs(ptrs []*int) string {
	vals := make([]int, len(ptrs))

	for i, p := range ptrs {
		vals[i] = *p
	}

	return fmt.Sprintf("%v", vals)
}

// toStrPtrs converts []string to []*string.
func toStrPtrs(values []string) []*string {
	ptrs := make([]*string, len(values))

	for i := range values {
		v := values[i]
		ptrs[i] = &v
	}

	return ptrs
}

// formatStrPtrs formats []*string to "[v1 v2 v3]" matching fmt.Sprintf("%v", []string{...}).
func formatStrPtrs(ptrs []*string) string {
	vals := make([]string, len(ptrs))

	for i, p := range ptrs {
		vals[i] = *p
	}

	return fmt.Sprintf("%v", vals)
}
