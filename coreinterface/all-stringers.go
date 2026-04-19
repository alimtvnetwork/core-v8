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

package coreinterface

type AllKeysStringer interface {
	AllKeys() []string
}

type AllKeysSortedStringer interface {
	AllKeysSorted() []string
}

type JsonCombineStringer interface {
	JsonStringer
	// MustJsonStringer panic if any error
	MustJsonStringer
}

type BuildStringer interface {
	Build() string
}

type ByteToStringer interface {
	String(input byte) string
}

type MustJsonStringer interface {
	JsonStringMust() string
}

type NameValueStringer interface {
	NameValue() string
}

type FullStringer interface {
	FullString() string
}

type FullStringWithTracer interface {
	FullStringWithTraces() string
}

type ToValueStringer interface {
	Value() string
}

// ToNumberStringer
//
//	It returns string number value.
//
// Examples:
//   - ToNumberString() -> "1"  if the value is 1
//   - ToNumberString() -> "10" if the value is 10
type ToNumberStringer interface {
	// ToNumberString
	//
	//  It returns string number value.
	//
	// Examples:
	//  - ToNumberString() -> "1"  if the value is 1
	//  - ToNumberString() -> "10" if the value is 10
	ToNumberString() string
}

type ValidationCheckerWithStringer interface {
	Stringer
	IsInvalidChecker
	IsValidChecker
}

type SafeStringer interface {
	SafeString() string
}
