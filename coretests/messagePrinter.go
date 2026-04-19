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

package coretests

import (
	"log/slog"
)

type printMessage struct{}

func (it printMessage) FailedExpected(
	isFailed bool,
	when,
	actual,
	expected any,
	counter int,
) {
	if isFailed {
		message := GetAssert.Quick(when, actual, expected, counter)

		slog.Warn("test failed", "message", message)
	}
}

// PrintNameValue
//
// Print using msgformats.PrintValuesFormat
func (it printMessage) NameValue(header string, anyItem any) {
	toString := ToStringNameValues(anyItem)

	slog.Info("name-value",
		"header", header,
		"value", anyItem,
		"toString", toString,
	)
}

// PrintValue
//
// Print values using msgformats.PrintValuesFormat
func (it printMessage) Value(header string, anyItem any) {
	toString := ToStringValues(anyItem)

	slog.Info("value",
		"header", header,
		"value", anyItem,
		"toString", toString,
	)
}
