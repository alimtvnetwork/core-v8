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

// aliases.go defines common result type aliases for frequently used
// return types. These are type aliases (=) to preserve %T output.
package results

// ResultAny is the untyped version of Result[T].
type ResultAny = Result[any]

// ResultBool is a Result for functions returning bool.
type ResultBool = Result[bool]

// ResultString is a Result for functions returning string.
type ResultString = Result[string]

// ResultInt is a Result for functions returning int.
type ResultInt = Result[int]

// ResultError is a Result for functions returning error.
type ResultError = Result[error]

// ResultsAny is the untyped version of Results[T1, T2].
type ResultsAny = Results[any, any]

// ResultsAnyError is for functions returning (any, error).
type ResultsAnyError = Results[any, error]

// ResultsStringError is for functions returning (string, error).
type ResultsStringError = Results[string, error]

// ResultsBoolError is for functions returning (bool, error).
type ResultsBoolError = Results[bool, error]

// ResultsIntError is for functions returning (int, error).
type ResultsIntError = Results[int, error]
