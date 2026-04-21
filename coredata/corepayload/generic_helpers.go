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

package corepayload

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/defaulterr"
)

// DeserializePayloadTo deserializes the PayloadWrapper's Payloads bytes into T.
//
// Usage:
//
//	user, err := corepayload.DeserializePayloadTo[User](wrapper)
func DeserializePayloadTo[T any](wrapper *PayloadWrapper) (T, error) {
	var result T
	if wrapper == nil || len(wrapper.Payloads) == 0 {
		return result, defaulterr.NilResult
	}

	err := corejson.Deserialize.UsingBytes(wrapper.Payloads, &result)
	return result, err
}

// DeserializePayloadToMust deserializes the PayloadWrapper's Payloads bytes into T or panics.
//
// Usage:
//
//	user := corepayload.DeserializePayloadToMust[User](wrapper)
func DeserializePayloadToMust[T any](wrapper *PayloadWrapper) T {
	result, err := DeserializePayloadTo[T](wrapper)
	if err != nil {
		panic(err)
	}
	return result
}

// DeserializePayloadToSlice deserializes the PayloadWrapper's Payloads bytes into []T.
//
// Usage:
//
//	users, err := corepayload.DeserializePayloadToSlice[User](wrapper)
func DeserializePayloadToSlice[T any](wrapper *PayloadWrapper) ([]T, error) {
	var result []T
	if wrapper == nil || len(wrapper.Payloads) == 0 {
		return []T{}, defaulterr.NilResult
	}

	err := corejson.Deserialize.UsingBytes(wrapper.Payloads, &result)
	return result, err
}

// DeserializePayloadToSliceMust deserializes the PayloadWrapper's Payloads into []T or panics.
func DeserializePayloadToSliceMust[T any](wrapper *PayloadWrapper) []T {
	result, err := DeserializePayloadToSlice[T](wrapper)
	if err != nil {
		panic(err)
	}
	return result
}

// DeserializeAttributesPayloadTo deserializes the Attributes' DynamicPayloads bytes into T.
//
// Usage:
//
//	config, err := corepayload.DeserializeAttributesPayloadTo[AppConfig](attrs)
func DeserializeAttributesPayloadTo[T any](attr *Attributes) (T, error) {
	var result T
	if attr == nil || len(attr.DynamicPayloads) == 0 {
		return result, defaulterr.NilResult
	}

	err := corejson.Deserialize.UsingBytes(attr.DynamicPayloads, &result)
	return result, err
}

// DeserializeAttributesPayloadToMust deserializes Attributes' DynamicPayloads into T or panics.
func DeserializeAttributesPayloadToMust[T any](attr *Attributes) T {
	result, err := DeserializeAttributesPayloadTo[T](attr)
	if err != nil {
		panic(err)
	}
	return result
}

// DeserializeAttributesPayloadToSlice deserializes the Attributes' DynamicPayloads into []T.
func DeserializeAttributesPayloadToSlice[T any](attr *Attributes) ([]T, error) {
	var result []T
	if attr == nil || len(attr.DynamicPayloads) == 0 {
		return []T{}, defaulterr.NilResult
	}

	err := corejson.Deserialize.UsingBytes(attr.DynamicPayloads, &result)
	return result, err
}
