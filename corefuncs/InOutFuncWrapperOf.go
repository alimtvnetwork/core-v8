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

import "github.com/alimtvnetwork/core-v8/errcore"

// InOutFuncWrapperOf is the generic version of an input-output function wrapper.
//
// It wraps a typed InOutFuncOf[TIn, TOut] with a name for identification
// and provides conversion to ActionFunc/ActionReturnsErrorFunc.
//
// Usage:
//
//	wrapper := corefuncs.InOutFuncWrapperOf[string, int]{
//	    Name:   "strlen",
//	    Action: func(s string) int { return len(s) },
//	}
//	result := wrapper.Exec("hello") // result is int(5)
type InOutFuncWrapperOf[TIn any, TOut any] struct {
	Name   string
	Action InOutFuncOf[TIn, TOut]
}

// Exec runs the wrapped function with the given typed input.
func (it InOutFuncWrapperOf[TIn, TOut]) Exec(input TIn) TOut {
	return it.Action(input)
}

// AsActionFunc returns an ActionFunc that executes with the given input (discarding output).
func (it InOutFuncWrapperOf[TIn, TOut]) AsActionFunc(input TIn) ActionFunc {
	return func() {
		it.Action(input)
	}
}

// AsActionReturnsErrorFunc returns an ActionReturnsErrorFunc (never errors).
func (it InOutFuncWrapperOf[TIn, TOut]) AsActionReturnsErrorFunc(
	input TIn,
) ActionReturnsErrorFunc {
	return func() error {
		it.Action(input)

		return nil
	}
}

// ToLegacy converts to the non-generic InOutErrFuncWrapper for backward compatibility.
func (it InOutFuncWrapperOf[TIn, TOut]) ToLegacy() InOutErrFuncWrapper {
	return InOutErrFuncWrapper{
		Name: it.Name,
		Action: func(input any) (any, error) {
			return it.Action(input.(TIn)), nil
		},
	}
}

// SerializeOutputFuncWrapperOf wraps a typed serialization function.
//
// Usage:
//
//	wrapper := corefuncs.SerializeOutputFuncWrapperOf[MyStruct]{
//	    Name:   "json-marshal",
//	    Action: func(m MyStruct) ([]byte, error) { return json.Marshal(m) },
//	}
//	bytes, err := wrapper.Exec(myStruct)
type SerializeOutputFuncWrapperOf[TIn any] struct {
	Name   string
	Action SerializeOutputFuncOf[TIn]
}

// Exec runs the serialization function with typed input.
func (it SerializeOutputFuncWrapperOf[TIn]) Exec(input TIn) ([]byte, error) {
	return it.Action(input)
}

// AsActionReturnsErrorFunc returns an ActionReturnsErrorFunc that captures the input.
func (it SerializeOutputFuncWrapperOf[TIn]) AsActionReturnsErrorFunc(
	input TIn,
) ActionReturnsErrorFunc {
	return func() error {
		_, err := it.Action(input)

		if err != nil {
			return errcore.
				FailedToExecuteType.
				Error(err.Error()+", function name:", it.Name)
		}

		return nil
	}
}
