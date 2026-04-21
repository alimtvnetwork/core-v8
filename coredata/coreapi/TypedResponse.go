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

package coreapi

import "github.com/alimtvnetwork/core-v8/constants"

// TypedResponse is the generic API response type.
//
// T represents the strongly-typed response payload.
//
// Usage:
//
//	resp := coreapi.NewTypedResponse[MyResult](attr, result)
//	resp.Response.SomeField // fully typed
type TypedResponse[T any] struct {
	Attribute *ResponseAttribute `json:"Attribute,omitempty"`
	Response  T                  `json:"Response,omitempty"`
}

// NewTypedResponse creates a valid TypedResponse with the given attribute and response.
func NewTypedResponse[T any](
	attribute *ResponseAttribute,
	response T,
) *TypedResponse[T] {
	return &TypedResponse[T]{
		Attribute: attribute,
		Response:  response,
	}
}

// InvalidTypedResponse creates an invalid TypedResponse with a zero-value response.
func InvalidTypedResponse[T any](
	attribute *ResponseAttribute,
) *TypedResponse[T] {
	if attribute == nil {
		attribute = InvalidResponseAttribute(constants.EmptyString)
	}

	return &TypedResponse[T]{
		Attribute: attribute,
	}
}

// Clone returns a deep copy of the TypedResponse.
func (it *TypedResponse[T]) Clone() *TypedResponse[T] {
	if it == nil {
		return nil
	}

	return &TypedResponse[T]{
		Attribute: it.Attribute.Clone(),
		Response:  it.Response,
	}
}

// TypedResponseResult converts to a TypedResponseResult[T].
func (it *TypedResponse[T]) TypedResponseResult() *TypedResponseResult[T] {
	if it == nil {
		return nil
	}

	return &TypedResponseResult[T]{
		Attribute: it.Attribute,
		Response:  it.Response,
	}
}
