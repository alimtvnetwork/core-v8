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

// FunctionsExecuteResults is the generic version of AnyFunctionsExecuteResults.
// It executes the appropriate set of functions based on the condition
// and collects results of type T.
//
// Each function returns (result T, isTake bool, isBreak bool):
//   - isTake: include this result in the output slice
//   - isBreak: stop executing further functions
//
// Usage:
//
//	results := conditional.FunctionsExecuteResults[int](true, trueFuncs, falseFuncs)
func FunctionsExecuteResults[T any](
	isTrue bool,
	trueValueFunctions, falseValueFunctions []func() (
		result T,
		isTake,
		isBreak bool,
	),
) []T {
	if isTrue {
		return executeFunctions[T](trueValueFunctions)
	}

	return executeFunctions[T](falseValueFunctions)
}

func executeFunctions[T any](
	functions []func() (
		result T,
		isTake,
		isBreak bool,
	),
) []T {
	if len(functions) == 0 {
		return nil
	}

	results := make([]T, 0, len(functions))

	for _, curFunc := range functions {
		if curFunc == nil {
			continue
		}

		result, isTake, isBreak := curFunc()

		if isTake {
			results = append(results, result)
		}

		if isBreak {
			return results
		}
	}

	return results
}
