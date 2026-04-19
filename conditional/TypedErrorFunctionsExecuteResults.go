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

package conditional

import (
	"strconv"

	"github.com/alimtvnetwork/core/errcore"
)

// TypedErrorFunctionsExecuteResults is the generic version of ErrorFunctionsExecuteResults.
// It executes the appropriate set of functions based on the condition,
// collects results of type T, and aggregates any errors.
//
// Each function returns (result T, err error).
// Results are collected only when err is nil.
//
// Usage:
//
//	results, err := conditional.TypedErrorFunctionsExecuteResults[int](true, trueFuncs, falseFuncs)
func TypedErrorFunctionsExecuteResults[T any](
	isTrue bool,
	trueValueFunctions []func() (T, error),
	falseValueFunctions []func() (T, error),
) ([]T, error) {
	if isTrue {
		return executeTypedErrorFunctions[T](trueValueFunctions)
	}

	return executeTypedErrorFunctions[T](falseValueFunctions)
}

func executeTypedErrorFunctions[T any](
	functions []func() (T, error),
) ([]T, error) {
	if len(functions) == 0 {
		return nil, nil
	}

	results := make([]T, 0, len(functions))
	var sliceErr []string

	for index, currentFunction := range functions {
		if currentFunction == nil {
			continue
		}

		result, err := currentFunction()

		if err != nil {
			sliceErr = append(sliceErr, err.Error()+"- index of - "+strconv.Itoa(index))
			continue
		}

		results = append(results, result)
	}

	return results, errcore.SliceToError(sliceErr)
}
