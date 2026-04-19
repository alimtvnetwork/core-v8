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

import (
	"reflect"

	"github.com/alimtvnetwork/core/internal/internalinterface"
)

type StringRangesGetter interface {
	StringRangesPtr() []string
	StringRanges() []string
}

type ReflectTypeGetter interface {
	ReflectType() reflect.Type
	GetErrorOnTypeMismatch(
		typeMatch reflect.Type,
		isIncludeInvalidMessage bool,
	) error
}

type ErrorGetter interface {
	Error() error
}

type MetaGetter interface {
	Meta() string
}

type MessageGetter interface {
	Message() string
}

type MeaningFullErrorGetter interface {
	MeaningFullError() error
}

type RawPayloadsGetter interface {
	internalinterface.RawPayloadsGetter
}

type ListStringsGetter interface {
	ListStrings() []string
}

type LengthGetter interface {
	Length() int
}

type EntityTypeNameGetter interface {
	EntityTypeName() string
}

type TypeNameWithRangeNamesCsvGetter interface {
	StringRangeNamesCsvGetter
	TypeNameGetter
}

type CountGetter interface {
	Count() int
}

type StringsGetter interface {
	Strings() []string
}

type IdentifierGetter interface {
	Identifier() string
}

type InvalidMessageGetter interface {
	InvalidMessage() string
}

type TypeNameGetter interface {
	TypeName() string
}

type TypenameStringGetter interface {
	TypenameString() string
}

type StringsLinesGetter interface {
	Lines() []string
}

type NameGetter interface {
	Name() string
}

type CategoryNameGetter interface {
	CategoryName() string
}

type AnyAttributesGetter interface {
	internalinterface.AnyAttributesGetter
}

// DynamicPageItemsGetter Paging items related methods
type DynamicPageItemsGetter interface {
	GetDynamicPagedItems(perPageItems int, pageIndex int) any
}

type IntRangeLengthGetter interface {
	RangeLength() int
}

type IntRangesGetter interface {
	RangesInt() []int
	Ranges() []int
}

type InvalidDirectErrorGetter interface {
	GetInvalidDirectError() error
}

type InvalidErrorGetter interface {
	InvalidError() error
}

type LastIndexGetter interface {
	LastIndex() int
}

type RawMessageBytesGetter interface {
	RawMessageBytes() []byte
}

type ReflectKindGetter interface {
	ReflectKind() reflect.Kind
}

type SafeBytesGetter interface {
	SafeBytes() []byte
}

type SafeBytesPointerGetter interface {
	SafeBytesPtr() []byte
}

type StringRangeNamesCsvGetter interface {
	RangeNamesCsv() string
}

type ReflectValueGetter interface {
	ReflectValue() *reflect.Value
}

type SafeStringsGetter interface {
	SafeStrings() []string
}

type StringRangeNamesGetter interface {
	StringRangeNames() []string
}

type UintUserIdGetter interface {
	UserId() uint
}

type UserIdGetter interface {
	UserId() int
}

type UsernameGetter interface {
	Username() string
}

type ValidationMessageGetter interface {
	ValidationMessage() string
}

type ValidationErrorGetter interface {
	ValidationError() error
}

type ValueAnyItemGetter interface {
	Value() any
}

type ValueByteGetter interface {
	Value() byte
}

type ValueErrorGetter interface {
	Value() error
}

type ValueFloat32Getter interface {
	Value() float32
}

type ValueFloat64Getter interface {
	Value() float64
}

type ValueGetter interface {
	ValueDynamic() any
}

type ValueInt16Getter interface {
	Value() int16
}

type ValueInt32Getter interface {
	Value() int32
}

type ValueInt64Getter interface {
	Value() int64
}

type ValueIntegerGetter interface {
	Value() int
}

type ValueReflectSetter interface {
	ValueReflectSet(setterPtr any) error
}

type ValueStringGetter interface {
	Value() string
}

type ValueStringsGetter interface {
	Value() []string
}

type FilterTextGetter interface {
	FilterText() string
}

type CategoryRevealerGetter interface {
	CategoryRevealer() CategoryRevealer
}

type StackTracesBytesGetter interface {
	StackTracesBytes() []byte
}

type JsonErrorBytesGetter interface {
	JsonErrorBytes() []byte
}

type MapStringAnyGetter interface {
	MapStringAny() map[string]any
}

type MapStringStringGetter interface {
	MapStringString() map[string]string
}

type MapStringFuncGetter interface {
	MapStringString() map[string]func()
}

type MapStringErrorFuncGetter interface {
	MapStringString() map[string]func() error
}

type VariableNameStringGetter interface {
	VariableName() string
}

type VariableNameAnyGetter interface {
	VariableNameAny() any
}

type ValueAnyGetter interface {
	ValueAny() any
}

type ExplicitValueStringGetter interface {
	ValueString() string
}

type KeysHashsetGetter interface {
	KeysHashset() map[string]bool
}

type HashmapGetter interface {
	Hashmap() map[string]string
}

type AllValuesStringsGetter interface {
	AllValuesStrings() []string
}

type AllValuesIntegersGetter interface {
	AllValuesIntegers() []int
}

type KeyStringNameGetter interface {
	KeyName() string
}
