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

package csvinternal

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

// StringersToCsvStrings
//
// Formats :
//   - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//   - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//   - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func StringersToCsvStrings(
	isIncludeQuote,
	isIncludeSingleQuote bool,
	stringerFunctions ...fmt.Stringer,
) []string {
	if len(stringerFunctions) == 0 {
		return []string{}
	}

	slice := make([]string, len(stringerFunctions))

	if isIncludeQuote && isIncludeSingleQuote {
		// single quote
		for i, stringerFunc := range stringerFunctions {
			slice[i] = fmt.Sprintf(
				constants.ValueWithSingleQuoteFormat,
				stringerFunc.String())
		}

		return slice
	} else if isIncludeQuote && !isIncludeSingleQuote {
		// double quote
		for i, stringerFunc := range stringerFunctions {
			slice[i] = fmt.Sprintf(
				constants.ValueWithDoubleQuoteFormat,
				stringerFunc.String())
		}

		return slice
	}

	// no quote
	for i, stringerFunc := range stringerFunctions {
		slice[i] = fmt.Sprintf(
			constants.SprintValueFormat,
			stringerFunc.String())
	}

	return slice
}
