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
	"errors"
	"fmt"
)

// MustBeValid panics if the FuncWrap is nil or invalid.
func (it *FuncWrap[T]) MustBeValid() {
	if it == nil {
		panic("cannot execute on nil func-wrap")
	}

	if it.IsInvalid() {
		panic("func-wrap invalid - " + it.Name)
	}
}

// ValidationError returns an error if the FuncWrap is nil or invalid,
// or nil if it is valid.
func (it *FuncWrap[T]) ValidationError() error {
	if it == nil {
		return errors.New("cannot execute on nil func-wrap")
	}

	if it.IsInvalid() {
		return fmt.Errorf(
			"func-wrap is invalid:\n    given type: %T\n    name: %s",
			it.Func,
			it.Name,
		)
	}

	return nil
}

// ValidateMethodArgs validates that the given arguments match the
// expected count and types of the wrapped function's parameters.
func (it *FuncWrap[T]) ValidateMethodArgs(args []any) error {
	expectedCount := it.ArgsCount()
	given := len(args)

	if given != expectedCount {
		return errors.New(
			it.argsCountMismatchErrorMessage(
				expectedCount,
				given,
				args,
			),
		)
	}

	_, err := it.VerifyInArgs(args)

	return err
}

// InvalidError returns a descriptive error explaining why the FuncWrap is invalid,
// or nil if it is valid.
func (it *FuncWrap[T]) InvalidError() error {
	if it == nil {
		return errors.New("func-wrap is nil")
	}

	if !it.rv.IsValid() {
		return errors.New("reflect value is invalid")
	}

	if !it.HasValidFunc() {
		return errors.New("func-wrap request doesn't hold a valid func reflect")
	}

	return nil
}
