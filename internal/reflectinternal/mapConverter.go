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
	"sort"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/internal/convertinternal"
)

type mapConverter struct{}

func (it mapConverter) Length(i any) int {
	return SliceConverter.Length(i)
}

// ToStringsRv
//
//	expectation : map[key:string]...value don't care.
func (it mapConverter) ToStringsRv(reflectVal reflect.Value) ([]string, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ToStringsRv(
			reflect.Indirect(reflectVal),
		)
	}

	if reflectVal.Kind() != reflect.Map {
		return []string{},
			it.notAMapErr(reflectVal)
	}

	mapKeys := reflectVal.MapKeys()
	length := len(mapKeys)
	keys := make([]string, length)

	for i, key := range reflectVal.MapKeys() {
		keyAny := key.Interface()
		keyAsString, isString := keyAny.(string)

		isNotString := !isString

		if isNotString {
			return keys, it.notStringErr(keyAny)
		}

		keys[i] = keyAsString
	}

	return keys, nil
}

func (it mapConverter) notStringErr(keyAny any) error {
	return fmt.Errorf("not string type : %T", keyAny)
}

func (it mapConverter) notAMapErr(reflectVal reflect.Value) error {
	return fmt.Errorf("reflection is not map but %s", reflectVal.String())
}

// ToKeysStrings
//
//	expectation : map[key:string]...value don't care.
func (it mapConverter) ToKeysStrings(i any) ([]string, error) {
	return it.ToStrings(i)
}

// ToValuesAny
//
//	expectation : map[...]...value don't care.
func (it mapConverter) ToValuesAny(i any) ([]any, error) {
	if Is.Null(i) {
		return []any{}, nil
	}

	rv := reflect.ValueOf(i)

	var list []any

	err := Looper.MapForRv(
		rv, func(total int, index int, key, v any) (err error) {
			list = append(list, v)

			return nil
		},
	)

	return list, err
}

// ToKeysAny
//
//	expectation : map[...]...value don't care.
func (it mapConverter) ToKeysAny(i any) ([]any, error) {
	if Is.Null(i) {
		return []any{}, nil
	}

	rv := reflect.ValueOf(i)

	var list []any

	err := Looper.MapForRv(
		rv, func(total int, index int, key, v any) (err error) {
			list = append(list, key)

			return nil
		},
	)

	return list, err
}

// ToKeysValuesAny
//
//	expectation : map[string]...value don't care.
func (it mapConverter) ToKeysValuesAny(i any) (keys []string, values []any, err error) {
	if Is.Null(i) {
		return []string{}, []any{}, nil
	}

	rv := reflect.ValueOf(i)
	toStringFunc := convertinternal.AnyTo.SmartString

	err = Looper.MapForRv(
		rv, func(total int, index int, key, v any) (err error) {
			keys = append(keys, toStringFunc(key))
			values = append(values, v)

			return nil
		},
	)

	return keys, values, err
}

// ToStrings
//
//	expectation : map[key:string]don't care values
func (it mapConverter) ToStrings(anyItem any) ([]string, error) {
	if Is.Null(anyItem) {
		return []string{}, nil
	}

	reflectVal := reflect.ValueOf(anyItem)

	return it.ToStringsRv(reflectVal)
}

func (it mapConverter) ToStringsMust(anyItem any) []string {
	var reflectVal reflect.Value
	if rv, ok := anyItem.(reflect.Value); ok {
		reflectVal = rv
	} else {
		reflectVal = reflect.ValueOf(anyItem)
	}

	mapKeys, err := it.ToStringsRv(reflectVal)

	if err != nil {
		panic(err)
	}

	return mapKeys
}

func (it mapConverter) ToSortedStrings(anyItem any) ([]string, error) {
	if Is.Null(anyItem) {
		return []string{}, nil
	}

	reflectVal := reflect.ValueOf(anyItem)

	keys, err := it.ToStringsRv(reflectVal)

	if err != nil {
		return keys, err
	}

	sort.Strings(keys)

	return keys, nil
}

func (it mapConverter) ToSortedStringsMust(anyItem any) []string {
	if Is.Null(anyItem) {
		return []string{}
	}

	var reflectVal reflect.Value
	if rv, ok := anyItem.(reflect.Value); ok {
		reflectVal = rv
	} else {
		reflectVal = reflect.ValueOf(anyItem)
	}

	keys, err := it.ToStringsRv(reflectVal)
	if err != nil {
		panic(err)
	}
	sort.Strings(keys)

	return keys
}

// ToMapStringAnyRv
//
//	expectation : map[key:any]any to map[string]any
func (it mapConverter) ToMapStringAnyRv(
	reflectVal reflect.Value,
) (map[string]any, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return it.ToMapStringAnyRv(
			reflect.Indirect(reflectVal),
		)
	}

	if reflectVal.Kind() != reflect.Map {
		return map[string]any{},
			it.notAMapErr(reflectVal)
	}

	mapKeys := reflectVal.MapKeys()
	newMap := make(
		map[string]any,
		reflectVal.Len()+1,
	)

	for _, key := range mapKeys {
		value := reflectVal.MapIndex(key)
		keyAny := key.Interface()
		var keyAsString string

		keyAsString, isString := keyAny.(string)
		if isString {
			newMap[keyAsString] = value.Interface()
			continue
		}

		keyAsString = fmt.Sprintf(
			constants.SprintValueFormat,
			keyAny,
		)
		newMap[keyAsString] = value.Interface()
	}

	return newMap, nil
}

// ToMapStringAny
//
//	expectation : map[key:any]any to map[string]any
func (it mapConverter) ToMapStringAny(
	i any,
) (map[string]any, error) {
	if Is.Null(i) {
		return map[string]any{}, nil
	}

	return it.ToMapStringAnyRv(
		reflect.ValueOf(i),
	)
}
