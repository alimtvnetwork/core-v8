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

import (
	"reflect"
	"runtime"
	"strings"
)

// MethodName extracts the method name from a direct function reference.
//
// Example:
//
//	MethodName((*MyStruct).IsValid) → "IsValid"
//	MethodName((*MyStruct).String)  → "String"
func MethodName(funcRef any) string {
	if funcRef == nil {
		return ""
	}

	rv := reflect.ValueOf(funcRef)

	if rv.Kind() != reflect.Func {
		return ""
	}

	fullName := runtime.FuncForPC(rv.Pointer()).Name()

	// fullName looks like: "github.com/alimtvnetwork/core/pkg.(*Type).Method-fm"
	// We want just "Method"

	lastDot := strings.LastIndex(fullName, ".")

	if lastDot < 0 {
		return fullName
	}

	name := fullName[lastDot+1:]

	// Strip Go's "-fm" suffix for method expressions
	name = strings.TrimSuffix(name, "-fm")

	return name
}
