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

type newLazyRegexCreator struct{}

// New
//
//	used to create as vars
func (it newLazyRegexCreator) New(
	pattern string,
) *LazyRegex {
	lazyRegex, _ := lazyRegexOnceMap.CreateOrExisting(
		pattern)

	return lazyRegex
}

// NewLock
//
//	used to generate inside method
func (it newLazyRegexCreator) NewLock(
	pattern string,
) *LazyRegex {
	lazyRegex, _ := lazyRegexOnceMap.CreateOrExistingLock(
		pattern)

	return lazyRegex
}

func (it newLazyRegexCreator) TwoLock(
	pattern, secondPattern string,
) (first, second *LazyRegex) {
	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	first = it.New(pattern)
	second = it.New(secondPattern)

	return first, second
}

func (it newLazyRegexCreator) ManyUsingLock(
	patterns ...string,
) (patternsKeyAsMap map[string]*LazyRegex) {
	if len(patterns) == 0 {
		return map[string]*LazyRegex{}
	}

	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	patternsKeyAsMap = make(
		map[string]*LazyRegex,
		len(patterns))

	for _, pattern := range patterns {
		patternsKeyAsMap[pattern] = it.New(pattern)
	}

	return patternsKeyAsMap
}

func (it newLazyRegexCreator) AllPatternsMap() map[string]*LazyRegex {
	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	return lazyRegexOnceMap.items
}

// NewLockIf
//
//	used to generate inside method
//	lock must be performed when called from method.
func (it newLazyRegexCreator) NewLockIf(
	isLock bool,
	pattern string,
) *LazyRegex {
	if isLock {
		return it.NewLock(pattern)
	}

	return it.New(pattern)
}
