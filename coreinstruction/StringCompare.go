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

package coreinstruction

import (
	"github.com/alimtvnetwork/core-v8/corecomparator"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v8/regexnew"
)

type StringCompare struct {
	StringSearch
	Content string
}

func NewStringCompare(
	method stringcompareas.Variant,
	isIgnoreCase bool,
	search,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: method,
			Search:        search,
			BaseIsIgnoreCase: corecomparator.BaseIsIgnoreCase{
				IsIgnoreCase: isIgnoreCase,
			},
		},
		Content: content,
	}
}

func NewStringCompareEqual(
	search,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: stringcompareas.Equal,
			Search:        search,
		},
		Content: content,
	}
}

func NewStringCompareRegex(
	regex,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: stringcompareas.Regex,
			Search:        regex,
		},
		Content: content,
	}
}

func NewStringCompareStartsWith(
	isIgnoreCase bool,
	search,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: stringcompareas.StartsWith,
			Search:        search,
			BaseIsIgnoreCase: corecomparator.BaseIsIgnoreCase{
				IsIgnoreCase: isIgnoreCase,
			},
		},
		Content: content,
	}
}

func NewStringCompareEndsWith(
	isIgnoreCase bool,
	search,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: stringcompareas.EndsWith,
			Search:        search,
			BaseIsIgnoreCase: corecomparator.BaseIsIgnoreCase{
				IsIgnoreCase: isIgnoreCase,
			},
		},
		Content: content,
	}
}

func NewStringCompareContains(
	isIgnoreCase bool,
	search,
	content string,
) *StringCompare {
	return &StringCompare{
		StringSearch: StringSearch{
			CompareMethod: stringcompareas.Contains,
			Search:        search,
			BaseIsIgnoreCase: corecomparator.BaseIsIgnoreCase{
				IsIgnoreCase: isIgnoreCase,
			},
		},
		Content: content,
	}
}

func (it *StringCompare) IsInvalid() bool {
	return it == nil
}

func (it *StringCompare) IsDefined() bool {
	return it != nil
}

func (it *StringCompare) IsMatch() bool {
	if it == nil {
		return true
	}

	return it.CompareMethod.IsCompareSuccess(
		it.IsIgnoreCase,
		it.Content,
		it.Search,
	)
}

func (it *StringCompare) IsMatchFailed() bool {
	return !it.IsMatch()
}

func (it *StringCompare) VerifyError() error {
	if it == nil {
		return nil
	}

	if it.CompareMethod.IsRegex() {
		return regexnew.MatchErrorLock(
			it.Search,
			it.Content)
	}

	return it.CompareMethod.VerifyError(
		it.IsIgnoreCase,
		it.Content,
		it.Search,
	)
}
