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

package corepayload

type PagingInfo struct {
	CurrentPageIndex                     int // -- 1 based index
	TotalPages, PerPageItems, TotalItems int
}

func (it *PagingInfo) IsEmpty() bool {
	return it == nil ||
		it.TotalPages == 0 && it.TotalItems == 0
}

func (it *PagingInfo) IsEqual(right *PagingInfo) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	if it.TotalPages != right.TotalPages {
		return false
	}

	if it.CurrentPageIndex != right.CurrentPageIndex {
		return false
	}

	if it.PerPageItems != right.PerPageItems {
		return false
	}

	return it.TotalItems == right.TotalItems
}

func (it *PagingInfo) HasTotalPages() bool {
	return it != nil && it.TotalPages > 0
}

func (it *PagingInfo) HasCurrentPageIndex() bool {
	return it != nil && it.CurrentPageIndex > 0
}

func (it *PagingInfo) HasPerPageItems() bool {
	return it != nil && it.PerPageItems > 0
}

func (it *PagingInfo) HasTotalItems() bool {
	return it != nil && it.TotalItems > 0
}

func (it *PagingInfo) IsInvalidTotalPages() bool {
	return it == nil || it.TotalPages <= 0
}

func (it *PagingInfo) IsInvalidCurrentPageIndex() bool {
	return it == nil || it.CurrentPageIndex <= 0
}

func (it *PagingInfo) IsInvalidPerPageItems() bool {
	return it == nil || it.PerPageItems <= 0
}

func (it *PagingInfo) IsInvalidTotalItems() bool {
	return it == nil || it.TotalItems <= 0
}

func (it PagingInfo) Clone() PagingInfo {
	return PagingInfo{
		TotalPages:       it.TotalPages,
		CurrentPageIndex: it.CurrentPageIndex,
		PerPageItems:     it.PerPageItems,
		TotalItems:       it.TotalItems,
	}
}

func (it *PagingInfo) ClonePtr() *PagingInfo {
	if it == nil {
		return nil
	}

	return &PagingInfo{
		TotalPages:       it.TotalPages,
		CurrentPageIndex: it.CurrentPageIndex,
		PerPageItems:     it.PerPageItems,
		TotalItems:       it.TotalItems,
	}
}
