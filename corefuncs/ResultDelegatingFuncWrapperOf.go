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

import "github.com/alimtvnetwork/core/errcore"

// ResultDelegatingFuncWrapperOf is the generic version of ResultDelegatingFuncWrapper.
//
// It wraps a typed ResultDelegatingFuncOf[T] with a name for error reporting
// and provides conversion to ActionFunc/ActionReturnsErrorFunc.
//
// Usage:
//
//	wrapper := corefuncs.ResultDelegatingFuncWrapperOf[*MyStruct]{
//	    Name:   "unmarshal-user",
//	    Action: func(target *MyStruct) error { return json.Unmarshal(data, target) },
//	}
//	var user MyStruct
//	err := wrapper.Exec(&user)
type ResultDelegatingFuncWrapperOf[T any] struct {
	Name   string
	Action ResultDelegatingFuncOf[T]
}

// Exec runs the wrapped function with the given typed target.
func (it ResultDelegatingFuncWrapperOf[T]) Exec(
	toPointer T,
) error {
	return it.Action(toPointer)
}

// AsActionFunc returns an ActionFunc that executes with the given target.
// Errors are handled via errcore.HandleErr.
func (it ResultDelegatingFuncWrapperOf[T]) AsActionFunc(toPointer T) ActionFunc {
	return func() {
		actionReturnsErrFunc := it.AsActionReturnsErrorFunc(toPointer)
		errcore.HandleErr(actionReturnsErrFunc())
	}
}

// AsActionReturnsErrorFunc returns an ActionReturnsErrorFunc that captures the target.
func (it ResultDelegatingFuncWrapperOf[T]) AsActionReturnsErrorFunc(
	toPointer T,
) ActionReturnsErrorFunc {
	return func() error {
		err := it.Action(toPointer)

		if err != nil {
			return errcore.
				FailedToExecuteType.
				Error(err.Error()+", function name:", it.Name)
		}

		return nil
	}
}

// ToLegacy converts to the non-generic ResultDelegatingFuncWrapper for backward compatibility.
func (it ResultDelegatingFuncWrapperOf[T]) ToLegacy() ResultDelegatingFuncWrapper {
	return ResultDelegatingFuncWrapper{
		Name: it.Name,
		Action: func(resultDelegatedTo any) error {
			return it.Action(resultDelegatedTo.(T))
		},
	}
}
