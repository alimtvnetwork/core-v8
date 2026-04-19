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

// GetPagingInfo calculates paging metadata from a PagingRequest.
//
// Validation rules:
//   - Zero or negative EachPageSize → empty PagingInfo (no paging possible).
//   - Zero Length → empty PagingInfo with PageIndex 0.
//   - PageIndex below 1 → clamped to first page (1).
//   - PageIndex above total pages → clamped to last page.
//   - Length < EachPageSize → single page, IsPagingPossible = false.
func GetPagingInfo(request PagingRequest) PagingInfo {
	// Guard: invalid page size
	if isPageSizeInvalid(request.EachPageSize) {
		return PagingInfo{TotalPages: 0}
	}

	length := request.Length

	// Guard: no items
	if isLengthEmpty(length) {
		return PagingInfo{
			PageIndex:        0,
			SkipItems:        0,
			EndingLength:     0,
			TotalPages:       0,
			IsPagingPossible: false,
		}
	}

	// Guard: everything fits in one page
	if isPagingOutOfRange(length, request.EachPageSize) {
		return PagingInfo{
			PageIndex:        1,
			SkipItems:        0,
			EndingLength:     length,
			TotalPages:       1,
			IsPagingPossible: false,
		}
	}

	// Calculate total pages for clamping
	totalPages := GetPagesSize(request.EachPageSize, length)

	// Clamp page index to valid range
	pageIndex := clampedPageIndex(request.PageIndex, totalPages)

	// Calculate offsets
	skipItems := calculateSkipItems(pageIndex, request.EachPageSize)
	endingIndex := clampedEndingLength(skipItems+request.EachPageSize, length)

	return PagingInfo{
		PageIndex:        pageIndex,
		SkipItems:        skipItems,
		EndingLength:     endingIndex,
		TotalPages:       totalPages,
		IsPagingPossible: true,
	}
}
