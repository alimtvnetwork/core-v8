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

package reflectinternal

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/alimtvnetwork/core-v8/constants"
)

type reflectTypeConverter struct{}

func (it reflectTypeConverter) SafeName(anyItem any) string {
	rt := reflect.TypeOf(anyItem)

	if Is.Null(rt) {
		return ""
	}

	return rt.String()
}

func (it reflectTypeConverter) SafeTypeNameOfSliceOrSingle(
	isSingle bool,
	anyItem any,
) string {
	if isSingle {
		return it.SafeName(anyItem)
	}

	return it.SliceFirstItemTypeName(anyItem)
}

// SliceFirstItemTypeName
//
// Gets slice element type name, reduce ptr slice as well.
func (it reflectTypeConverter) SliceFirstItemTypeName(slice any) string {
	rt := reflect.TypeOf(slice)

	if Is.Null(rt) {
		return ""
	}

	if rt.Kind() == reflect.Ptr || rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	return rt.Elem().String()
}

func (it reflectTypeConverter) NamesStringUsingReflectType(
	isFullName bool,
	reflectTypes ...reflect.Type,
) string {
	if len(reflectTypes) == 0 {
		return ""
	}

	return strings.Join(
		it.NamesUsingReflectType(isFullName, reflectTypes...),
		constants.CommaSpace,
	)
}

func (it reflectTypeConverter) TypeNamesString(
	isFullName bool,
	anyItems ...any,
) string {
	if len(anyItems) == 0 {
		return ""
	}

	return strings.Join(
		TypeNames(isFullName, anyItems...),
		constants.CommaSpace,
	)
}

func (it reflectTypeConverter) NamesUsingReflectType(
	isFullName bool,
	reflectTypes ...reflect.Type,
) []string {
	if len(reflectTypes) == 0 {
		return []string{}
	}

	slice := make([]string, len(reflectTypes))

	if isFullName {
		for i, item := range reflectTypes {
			slice[i] = item.String()
		}

		return slice
	}

	for i, item := range reflectTypes {
		slice[i] = item.Name()
	}

	return slice
}

func (it reflectTypeConverter) NamesReferenceString(
	isFullName bool,
	anyItems ...any,
) string {
	if len(anyItems) == 0 {
		return ""
	}

	return "Reference (Types): " + strings.Join(
		TypeNames(isFullName, anyItems...),
		constants.CommaSpace,
	)
}

func (it reflectTypeConverter) Names(
	isFullName bool,
	anyItems ...any,
) []string {
	if len(anyItems) == 0 {
		return []string{}
	}

	slice := make([]string, len(anyItems))

	if isFullName {
		for i, item := range anyItems {
			slice[i] = reflect.TypeOf(item).String()
		}

		return slice
	}

	for i, item := range anyItems {
		slice[i] = reflect.TypeOf(item).Name()
	}

	return slice
}

func (it reflectTypeConverter) Name(anyItem any) string {
	rf := reflect.TypeOf(anyItem)

	if rf == nil {
		return ""
	}

	return rf.String()
}

func (it reflectTypeConverter) NameUsingFmt(anyItem any) string {
	return fmt.Sprintf(constants.SprintTypeFormat, anyItem)
}
