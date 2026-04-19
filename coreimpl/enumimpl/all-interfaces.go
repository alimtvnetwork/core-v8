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

package enumimpl

type valueByter interface {
	Value() byte
}

type exactValueByter interface {
	ValueByte() byte
}

type valueInter interface {
	Value() int
}

type exactValueInter interface {
	ValueInt() int
}

type valueInt8er interface {
	Value() int8
}

type exactValueInt8er interface {
	ValueInt8() int8
}

type valueUInt16er interface {
	Value() uint16
}

type exactValueUInt16er interface {
	ValueUInt16() uint16
}

type formatter interface {
	TypeName() string
	Name() string
	ValueString() string
}

type DifferChecker interface {
	GetSingleDiffResult(isLeft bool, l, r any) any
	GetResultOnKeyMissingInRightExistInLeft(lKey string, lVal any) any
	IsEqual(isRegardless bool, l, r any) bool
}
