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

type BasicByter interface {
	IsAnyOf(
		value byte,
		givenBytes ...byte,
	) bool
	Max() byte
	Min() byte
	GetValueByString(
		jsonValueString string,
	) byte
	GetValueByName(
		name string,
	) (byte, error)
	GetStringValue(
		input byte,
	) string
	Ranges() []byte
	Hashmap() map[string]byte
	HashmapPtr() *map[string]byte
	IsValidRange(
		value byte,
	) bool
	ToEnumJsonBytes(
		value byte,
	) ([]byte, error)
	ToEnumString(
		value byte,
	) string
	AppendPrependJoinValue(
		joiner string,
		appendVal, prependVal byte,
	) string
	AppendPrependJoinNamer(
		joiner string,
		appendVal, prependVal toNamer,
	) string
	ToNumberString(
		valueInRawFormat any,
	) string
	// UnmarshallToValue
	//
	//  isMappedToFirstIfEmpty: maps invalid values to first item
	UnmarshallToValue(
		isMappedToFirstIfEmpty bool,
		jsonUnmarshallingValue []byte,
	) (byte, error)
}
