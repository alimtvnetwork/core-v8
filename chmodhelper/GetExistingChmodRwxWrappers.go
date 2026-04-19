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

package chmodhelper

import "github.com/alimtvnetwork/core/errcore"

func GetExistingChmodRwxWrappers(
	isContinueOnError bool,
	locations ...string,
) (filePathToRwxWrapper map[string]*RwxWrapper, err error) {
	results := make(
		map[string]*RwxWrapper,
		len(locations))

	if len(locations) == 0 {
		return results, nil
	}

	if isContinueOnError {
		var sliceErr []string

		for _, location := range locations {
			wrapperPtr, err2 := GetExistingChmodRwxWrapperPtr(
				location)

			if err2 != nil {
				sliceErr = append(
					sliceErr,
					err2.Error())
			} else {
				results[location] = wrapperPtr
			}
		}

		return results, errcore.SliceToError(sliceErr)
	}

	// immediate exit
	for _, location := range locations {
		wrapperPtr, err2 := GetExistingChmodRwxWrapperPtr(
			location)

		if err2 != nil {
			return results, err2
		} else {
			results[location] = wrapperPtr
		}
	}

	return results, nil
}
