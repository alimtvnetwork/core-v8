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

package corefuncs

// Generic function types — type-safe versions of the any-based func types.
// These allow callers to work with concrete types instead of any,
// eliminating the need for type assertions at call sites.
//
// Each generic type mirrors a corresponding any-based type in funcs.go:
//
//	InOutFunc           → InOutFuncOf[TIn, TOut]
//	InOutErrFunc        → InOutErrFuncOf[TIn, TOut]
//	SerializeOutputFunc → SerializeOutputFuncOf[TIn]
//	InActionReturnsErrFunc → InActionReturnsErrFuncOf[TIn]
//	ResultDelegatingFunc   → ResultDelegatingFuncOf[T]
type (
	// InOutFuncOf is a generic version of InOutFunc.
	//
	//	any-based:  func(input any) (output any)
	//	generic:    func(input TIn) (output TOut)
	InOutFuncOf[TIn any, TOut any] func(input TIn) (output TOut)

	// InOutErrFuncOf is a generic version of InOutErrFunc.
	//
	//	any-based:  func(input any) (output any, err error)
	//	generic:    func(input TIn) (output TOut, err error)
	InOutErrFuncOf[TIn any, TOut any] func(input TIn) (output TOut, err error)

	// SerializeOutputFuncOf is a generic version of SerializeOutputFunc.
	//
	//	any-based:  func(input any) (serializedBytes []byte, err error)
	//	generic:    func(input TIn) (serializedBytes []byte, err error)
	SerializeOutputFuncOf[TIn any] func(input TIn) (serializedBytes []byte, err error)

	// InActionReturnsErrFuncOf is a generic version of InActionReturnsErrFunc.
	//
	//	any-based:  func(input any) (err error)
	//	generic:    func(input TIn) (err error)
	InActionReturnsErrFuncOf[TIn any] func(input TIn) (err error)

	// ResultDelegatingFuncOf is a generic version of ResultDelegatingFunc.
	//
	//	any-based:  func(resultDelegatedTo any) error
	//	generic:    func(resultDelegatedTo T) error
	//
	// resultDelegatedTo can be unmarshal or marshal or reflect set target.
	ResultDelegatingFuncOf[T any] func(resultDelegatedTo T) error
)
