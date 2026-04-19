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

import "github.com/alimtvnetwork/core/coredata/corejson"

type StandardEnumer interface {
	BasicEnumer
	StringRangesGetter
	RangeValidateChecker
	corejson.JsonContractsBinder
}

type BasicEnumValuer interface {
	ValueByte() byte
	ValueInt() int
	ValueInt8() int8
	ValueInt16() int16
	ValueUInt16() uint16
	ValueInt32() int32
	// ValueString
	//
	//  alias for ToNumberStringer
	//  returns the value number as string format (no name)
	//
	// Example:
	//  - "1", "2" NOT "Name"
	ValueString() string // value in string format
}

type BasicByteEnumer interface {
	UnmarshallEnumToValueByter
	MaxByte() byte
	MinByte() byte
	ValueByte() byte
	RangesByte() []byte
}

type BasicInt32Enumer interface {
	UnmarshallEnumToValueInt32(jsonUnmarshallingValue []byte) (int32, error)
	MaxInt32() int32
	MinInt32() int32
	ValueInt32() int32
	RangesInt32() []int32
	ToEnumString(input int32) string
}

type BasicInt16Enumer interface {
	UnmarshallEnumToValueInt16(jsonUnmarshallingValue []byte) (int16, error)
	MaxInt16() int16
	MinInt16() int16
	ValueInt16() int16
	RangesInt16() []int16
	ToEnumString(input int16) string
}

type BasicInt8Enumer interface {
	UnmarshallEnumToValueInt8(jsonUnmarshallingValue []byte) (int8, error)
	MaxInt8() int8
	MinInt8() int8
	ValueInt8() int8
	RangesInt8() []int8
	ToEnumString(input int8) string
}

type BasicIntEnumer interface {
	MaxInt() int
	MinInt() int
	ValueInt() int
	RangesInt() []int
	ToEnumString(input int) string
}

type BasicInt64Enumer interface {
	MaxInt64() int64
	MinInt64() int64
	ValueInt64() int64
	RangesInt64() []int64
}
