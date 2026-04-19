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

import "github.com/alimtvnetwork/core/internal/internalinterface"

type SplitNameValueByteGetter interface {
	enumNameStinger
	Value() byte
}

type SplitNameValueIntegerGetter interface {
	enumNameStinger
	Value() int
}

type SplitNameValueInteger8Getter interface {
	enumNameStinger
	Value() int8
}

type SplitNameValueInteger16Getter interface {
	enumNameStinger
	Value() int16
}

type SplitNameValueInteger32Getter interface {
	enumNameStinger
	Value() int32
}

type TypeNameGetter interface {
	TypeName() string
}

type StringRangeNamesCsvGetter interface {
	RangeNamesCsv() string
}

type TypeNameWithRangeNamesCsvGetter interface {
	StringRangeNamesCsvGetter
	TypeNameGetter
}

type ByteTypeEnumGetter interface {
	TypeEnum() BasicByteEnumContractsBinder
}

type StringRangesGetter interface {
	StringRangesPtr() []string
	StringRanges() []string
}

type BasicEnumerGetter interface {
	TypeBasicEnum() BasicEnumer
}

type RangeNamesCsvGetter interface {
	RangeNamesCsv() string
}

type RangesIntegerStringMapGetter interface {
	RangesIntegerStringMap() map[int]string
}

type RangesDynamicMapGetter interface {
	RangesDynamicMap() map[string]any
}

type IntegerEnumRangesGetter interface {
	IntegerEnumRanges() []int
}

type LoggerTyperGetter interface {
	LoggerTyper() LoggerTyper
}

type EventTyperGetter interface {
	EventTyper() EventTyper
}

type ErrorStringGetter interface {
	internalinterface.ErrorStringGetter
}
