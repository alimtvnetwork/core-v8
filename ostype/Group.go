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

package ostype

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
)

type Group byte

const (
	WindowsGroup Group = iota
	UnixGroup
	AndroidGroup
	InvalidGroup
)

func (it Group) AllNameValues() []string {
	return basicEnumImplOsGroup.AllNameValues()
}

func (it Group) OnlySupportedErr(names ...string) error {
	return basicEnumImplOsGroup.OnlySupportedErr(names...)
}

func (it Group) OnlySupportedMsgErr(message string, names ...string) error {
	return basicEnumImplOsGroup.OnlySupportedMsgErr(message, names...)
}

func (it Group) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Group) IntegerEnumRanges() []int {
	return basicEnumImplOsGroup.IntegerEnumRanges()
}

func (it Group) MinMaxAny() (min, max any) {
	return basicEnumImplOsGroup.MinMaxAny()
}

func (it Group) MinValueString() string {
	return basicEnumImplOsGroup.MinValueString()
}

func (it Group) MaxValueString() string {
	return basicEnumImplOsGroup.MaxValueString()
}

func (it Group) MaxInt() int {
	return basicEnumImplOsGroup.MaxInt()
}

func (it Group) MinInt() int {
	return basicEnumImplOsGroup.MinInt()
}

func (it Group) RangesDynamicMap() map[string]any {
	return basicEnumImplOsGroup.RangesDynamicMap()
}

func (it Group) IsByteValueEqual(value byte) bool {
	return byte(it) == value
}

func (it Group) Format(format string) (compiled string) {
	return basicEnumImplOsGroup.Format(format, it)
}

func (it Group) IsEnumEqual(enum enuminf.BasicEnumer) bool {
	return it.Value() == enum.ValueByte()
}

func (it *Group) IsAnyEnumsEqual(enums ...enuminf.BasicEnumer) bool {
	for _, enum := range enums {
		if it.IsEnumEqual(enum) {
			return true
		}
	}

	return false
}

func (it Group) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Group) IsAnyNamesOf(names ...string) bool {
	for _, name := range names {
		if it.IsNameEqual(name) {
			return true
		}
	}

	return false
}

func (it Group) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Group) IsAnyValuesEqual(anyByteValues ...byte) bool {
	for _, currentVal := range anyByteValues {
		if it.IsValueEqual(currentVal) {
			return true
		}
	}

	return false
}

func (it Group) ValueInt() int {
	return int(it)
}

func (it Group) ValueInt8() int8 {
	return int8(it)
}

func (it Group) ValueInt16() int16 {
	return int16(it)
}

func (it Group) ValueInt32() int32 {
	return int32(it)
}

func (it Group) ValueString() string {
	return it.ToNumberString()
}

func (it Group) Is(another Group) bool {
	return it == another
}

func (it Group) IsWindows() bool {
	return it == WindowsGroup
}

func (it Group) IsUnix() bool {
	return it == UnixGroup
}

func (it Group) IsAndroid() bool {
	return it == AndroidGroup
}

func (it Group) IsInvalidGroup() bool {
	return it == InvalidGroup
}

func (it Group) Byte() byte {
	return byte(it)
}

func (it Group) MarshalJSON() ([]byte, error) {
	return basicEnumImplOsGroup.ToEnumJsonBytes(it.Value())
}

func (it *Group) UnmarshalJSON(data []byte) error {
	valueByte, err := basicEnumImplOsGroup.UnmarshallToValue(
		true,
		data)

	if err == nil {
		*it = Group(valueByte)
	}

	return err
}

func (it Group) Name() string {
	return basicEnumImplOsGroup.ToEnumString(it.Value())
}

func (it Group) NameValue() string {
	return basicEnumImplOsGroup.NameWithValue(it.Value())
}

func (it Group) ToNumberString() string {
	return basicEnumImplOsGroup.ToNumberString(it.Value())
}

func (it Group) RangeNamesCsv() string {
	return basicEnumImplOsGroup.RangeNamesCsv()
}

func (it Group) TypeName() string {
	return basicEnumImplOsGroup.TypeName()
}

func (it Group) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return basicEnumImplOsGroup.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it Group) MaxByte() byte {
	return basicEnumImplOsGroup.Max()
}

func (it Group) MinByte() byte {
	return basicEnumImplOsGroup.Min()
}

func (it Group) ValueByte() byte {
	return byte(it)
}

func (it Group) RangesByte() []byte {
	return basicEnumImplOsGroup.Ranges()
}

func (it Group) Value() byte {
	return byte(it)
}

func (it Group) String() string {
	return basicEnumImplOsGroup.ToEnumString(it.Value())
}

func (it Group) IsValid() bool {
	return it != InvalidGroup
}

func (it Group) IsInvalid() bool {
	return it == InvalidGroup
}

func (it Group) EnumType() enuminf.EnumTyper {
	return basicEnumImplOsGroup.EnumType()
}

func (it *Group) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return it
}

func (it *Group) AsJsonContractsBinder() corejson.JsonMarshaller {
	return it
}

func (it Group) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Group) ToPtr() *Group {
	return &it
}
