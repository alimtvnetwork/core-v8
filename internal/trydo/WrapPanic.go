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

package trydo

func WrapPanic(voidFunc func()) Exception {
	var exception Exception

	Block{
		Try: func() {
			voidFunc()
		},
		Catch: func(e Exception) {
			exception = e
		},
		Finally: nil,
	}.Do()

	return exception
}

func ErrorFuncWrapPanic(errFunc func() error) WrappedErr {
	var exception Exception
	var err error
	var hasThrown bool

	Block{
		Try: func() {
			err = errFunc()
		},
		Catch: func(e Exception) {
			exception = e
			hasThrown = e != nil
		},
		Finally: nil,
	}.Do()

	return WrappedErr{
		Error:     err,
		Exception: exception,
		HasThrown: hasThrown,
		HasError:  err != nil,
	}
}
