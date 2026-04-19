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

package regexnew

import "regexp"

type lazyRegexMap struct {
	items map[string]*LazyRegex
}

func (it *lazyRegexMap) IsEmpty() bool {
	return it == nil || len(it.items) == 0
}

func (it *lazyRegexMap) IsEmptyLock() bool {
	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	return it == nil || len(it.items) == 0
}

func (it *lazyRegexMap) HasAnyItem() bool {
	return it != nil && len(it.items) > 0
}

func (it *lazyRegexMap) HasAnyItemLock() bool {
	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	return it != nil && len(it.items) > 0
}

func (it *lazyRegexMap) Length() int {
	if it == nil {
		return 0
	}

	return len(it.items)
}

func (it *lazyRegexMap) LengthLock() int {
	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	return it.Length()
}

func (it *lazyRegexMap) Has(keyName string) bool {
	_, has := it.items[keyName]

	return has
}

func (it *lazyRegexMap) HasLock(keyName string) bool {
	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	_, has := it.items[keyName]

	return has
}

func (it *lazyRegexMap) CreateOrExisting(
	patternName string,
) (lazyRegex *LazyRegex, isExisting bool) {
	lazyRegEx, has := it.items[patternName]

	if has {
		return lazyRegEx, has
	}

	// create
	lazyRegex = it.createDefaultLazyRegex(
		patternName,
	)

	it.items[patternName] = lazyRegex

	return lazyRegex, false
}

func (it *lazyRegexMap) CreateOrExistingLock(
	patternName string,
) (lazyRegex *LazyRegex, isExisting bool) {
	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	return it.CreateOrExisting(patternName)
}

func (it *lazyRegexMap) CreateOrExistingLockIf(
	isLock bool,
	patternName string,
) (lazyRegex *LazyRegex, isExisting bool) {
	if isLock {
		lazyRegexLock.Lock()
		defer lazyRegexLock.Unlock()

	}

	return it.CreateOrExisting(patternName)
}

func (it *lazyRegexMap) createDefaultLazyRegex(
	patternName string,
) (lazyRegex *LazyRegex) {
	return &LazyRegex{
		pattern:  patternName,
		compiler: CreateLock, // must use lock func
	}
}

func (it *lazyRegexMap) createLazyRegex(
	patternName string,
	creatorFunc func(pattern string) (*regexp.Regexp, error),
) (lazyRegex *LazyRegex) {
	return &LazyRegex{
		pattern:  patternName,
		compiler: creatorFunc,
	}
}
