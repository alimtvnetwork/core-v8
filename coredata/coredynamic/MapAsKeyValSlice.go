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

package coredynamic

import (
	"reflect"

	"github.com/alimtvnetwork/core-v8/errcore"
)

// MapAsKeyValSlice
//
//	expectation : map[key:any]any
func MapAsKeyValSlice(reflectVal reflect.Value) (*KeyValCollection, error) {
	if reflectVal.Kind() == reflect.Ptr {
		return MapAsKeyValSlice(
			reflect.Indirect(reflect.ValueOf(reflectVal)),
		)
	}

	if reflectVal.Kind() != reflect.Map {
		return EmptyKeyValCollection(),
			errcore.TypeMismatchType.Error("Reflection is not Map", reflectVal)
	}

	mapKeys := reflectVal.MapKeys()
	keyValCollection := NewKeyValCollection(len(mapKeys))

	for _, key := range reflectVal.MapKeys() {
		value := reflectVal.MapIndex(key)

		dynamicKV := &KeyVal{
			Key:   key.Interface(),
			Value: value.Interface(),
		}

		keyValCollection.AddPtr(dynamicKV)
	}

	return keyValCollection, nil
}
