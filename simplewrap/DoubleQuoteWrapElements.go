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

package simplewrap

// DoubleQuoteWrapElements Returns new empty slice if nil or empty slice given.
//
// Reference : https://play.golang.org/p/s_uN2-ckk2F | https://stackoverflow.com/a/48832120
func DoubleQuoteWrapElements(
	isSkipQuoteOnlyOnExistence bool,
	inputElements ...string,
) (doubleQuoteWrappedItems []string) {
	if inputElements == nil {
		return []string{}
	}

	length := len(inputElements)
	newSlice := make([]string, length)

	if length == 0 {
		return newSlice
	}

	if isSkipQuoteOnlyOnExistence {
		return wrapDoubleQuoteByExistenceCheck(inputElements, newSlice)
	}

	for i, item := range inputElements {
		newSlice[i] = WithDoubleQuote(item)
	}

	return newSlice
}
