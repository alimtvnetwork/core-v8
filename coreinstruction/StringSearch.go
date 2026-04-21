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

type StringSearch struct {
	corecomparator.BaseIsIgnoreCase
	CompareMethod stringcompareas.Variant
	Search        string
}

func (it *StringSearch) IsEmpty() bool {
	return it == nil
}

func (it *StringSearch) IsExist() bool {
	return it != nil
}

func (it *StringSearch) Has() bool {
	return it != nil
}

func (it *StringSearch) IsMatch(content string) bool {
	if it == nil {
		return true
	}

	return it.CompareMethod.IsCompareSuccess(
		it.IsIgnoreCase,
		content,
		it.Search,
	)
}

func (it *StringSearch) IsAllMatch(contents ...string) bool {
	if len(contents) == 0 {
		return true
	}

	for _, content := range contents {
		if it.IsMatchFailed(content) {
			return false
		}
	}

	return true
}

func (it *StringSearch) IsAnyMatchFailed(contents ...string) bool {
	return !it.IsAllMatch(contents...)
}

func (it *StringSearch) IsMatchFailed(content string) bool {
	return !it.IsMatch(content)
}

func (it *StringSearch) VerifyError(content string) error {
	if it == nil {
		return nil
	}

	if it.CompareMethod.IsRegex() {
		return regexnew.MatchErrorLock(
			it.Search,
			content)
	}

	return it.CompareMethod.VerifyError(
		it.IsIgnoreCase,
		content,
		it.Search,
	)
}
