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

package results

import "fmt"

// Results is a two-value typed result for functions returning (T1, T2).
//
// Embeds Result[T1] for the primary value and adds Result2 for the second.
type Results[T1, T2 any] struct {
	Result[T1]
	Result2 T2
}

// String returns a human-readable summary including both values.
func (it Results[T1, T2]) String() string {
	if it.Panicked {
		return fmt.Sprintf(
			"Results{panicked: %v, panicValue: %v}",
			it.Panicked,
			it.PanicValue,
		)
	}

	if it.Error != nil {
		return fmt.Sprintf(
			"Results{value: %v, result2: %v, error: %s}",
			it.Value,
			it.Result2,
			it.Error.Error(),
		)
	}

	return fmt.Sprintf(
		"Results{value: %v, result2: %v}",
		it.Value,
		it.Result2,
	)
}

// IsResult2 checks whether Result2 matches the given expected value.
func (it Results[T1, T2]) IsResult2(expected any) bool {
	return fmt.Sprintf("%v", it.Result2) == fmt.Sprintf("%v", expected)
}

// Result2String returns Result2 formatted via %v.
func (it Results[T1, T2]) Result2String() string {
	return fmt.Sprintf("%v", it.Result2)
}

// FromResultAny converts a ResultAny into a typed Results[T1, T2]
// by type-asserting AllResults[0] and AllResults[1].
//
// If type assertion fails, zero values are used.
func FromResultAny[T1, T2 any](r ResultAny) Results[T1, T2] {
	var res Results[T1, T2]

	res.Panicked = r.Panicked
	res.PanicValue = r.PanicValue
	res.Error = r.Error
	res.AllResults = r.AllResults
	res.ReturnCount = r.ReturnCount

	if len(r.AllResults) > 0 {
		if v, ok := r.AllResults[0].(T1); ok {
			res.Value = v
		}
	}

	if len(r.AllResults) > 1 {
		if v, ok := r.AllResults[1].(T2); ok {
			res.Result2 = v
		}
	}

	return res
}
