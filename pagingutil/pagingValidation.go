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

package pagingutil

// isLengthEmpty returns true when there are no items to page through.
func isLengthEmpty(length int) bool {
	return length <= 0
}

// isPageSizeInvalid returns true when page size is zero or negative.
func isPageSizeInvalid(eachPageSize int) bool {
	return eachPageSize <= 0
}

// isPageIndexBelowMinimum returns true when page index is less than 1.
func isPageIndexBelowMinimum(pageIndex int) bool {
	return pageIndex < 1
}

// isPageIndexAboveMaximum returns true when page index exceeds total pages.
func isPageIndexAboveMaximum(pageIndex, totalPages int) bool {
	return pageIndex > totalPages
}

// isPagingOutOfRange returns true when item count fits within a single page.
func isPagingOutOfRange(length, eachPageSize int) bool {
	return length < eachPageSize
}

// clampedPageIndex returns a valid page index:
//   - negative or zero → 1 (first page)
//   - beyond total pages → totalPages (last page)
//   - otherwise → original value
func clampedPageIndex(pageIndex, totalPages int) int {
	if isPageIndexBelowMinimum(pageIndex) {
		return 1
	}

	if isPageIndexAboveMaximum(pageIndex, totalPages) {
		return totalPages
	}

	return pageIndex
}

// calculateSkipItems returns the number of items to skip for a given page.
func calculateSkipItems(pageIndex, eachPageSize int) int {
	return eachPageSize * (pageIndex - 1)
}

// clampedEndingLength ensures ending index does not exceed total length.
func clampedEndingLength(endingIndex, length int) int {
	if endingIndex > length {
		return length
	}

	return endingIndex
}
