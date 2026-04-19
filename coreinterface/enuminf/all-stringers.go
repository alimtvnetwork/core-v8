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

package enuminf

import "fmt"

type ByteToEnumStringer interface {
	ToByteEnumString(input byte) string
}

type IntToEnumStringer interface {
	ToIntEnumString(input int) string
}

type Int8ToEnumStringer interface {
	ToInt8EnumString(input int8) string
}

type Int16ToEnumStringer interface {
	ToInt16EnumString(input int16) string
}

type Int32ToEnumStringer interface {
	ToInt32EnumString(input int32) string
}

type enumNameStinger interface {
	Namer
	fmt.Stringer
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
