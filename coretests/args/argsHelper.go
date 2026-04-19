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

package args

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// getByIndex safely retrieves an item from a slice by its index.
// Returns nil if the index is out of bounds.
func getByIndex(slice []any, index int) any {
	if len(slice)-1 < index {
		return nil
	}

	return slice[index]
}

// buildToString builds a formatted string representation for an arg type,
// using the given type name and cached SimpleStringOnce for memoization.
//
// Example output: "Three { first-val, second-val, third-val }"
func buildToString(
	typeName string,
	slice []any,
	cache *corestr.SimpleStringOnce,
) string {
	if cache.IsInitialized() {
		return cache.String()
	}

	var items []string

	for _, item := range slice {
		items = append(items, toString(item))
	}

	result := fmt.Sprintf(
		selfToStringFmt,
		typeName,
		strings.Join(items, constants.CommaSpace),
	)

	return cache.GetSetOnce(result)
}

// appendIfDefined appends the given value to the slice only if it is
// non-nil and non-zero as determined by reflectinternal.Is.Defined.
func appendIfDefined(args []any, value any) []any {
	if reflectinternal.Is.Defined(value) {
		return append(args, value)
	}

	return args
}

// invokeMustHelper invokes the given FuncWrapAny with args, panicking on error.
// This eliminates duplicate InvokeMust patterns across all Func arg types.
func invokeMustHelper(fw *FuncWrapAny, args ...any) []any {
	results, err := fw.Invoke(args...)

	if err != nil {
		panic(err)
	}

	return results
}
