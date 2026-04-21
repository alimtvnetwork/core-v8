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

package stringslice

import (
	"github.com/alimtvnetwork/core-v8/constants"
)

// LinesSimpleProcessNoEmpty split text using constants.NewLineUnix
// then returns lines processed by lineProcessor and don't take any empty string in the output.
// Empty string lineOut will be discarded from the final outputs.
func LinesSimpleProcessNoEmpty(
	splitsLines []string,
	lineProcessor func(lineIn string) (lineOut string),
) []string {
	if len(splitsLines) == 0 {
		return []string{}
	}

	slice := Make(constants.Zero, len(splitsLines))

	for _, lineIn := range splitsLines {
		lineOut := lineProcessor(lineIn)

		if lineOut == constants.EmptyString {
			continue
		}

		slice = append(slice, lineOut)
	}

	return slice
}
