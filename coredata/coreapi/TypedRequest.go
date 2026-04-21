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

import "github.com/alimtvnetwork/core-v8/coredata/coredynamic"

// TypedRequest is the generic request type.
//
// T represents the strongly-typed request payload, providing compile-time safety.
//
// Usage:
//
//	req := coreapi.NewTypedRequest[MyInput](attr, input)
//	req.Request.Field // strongly typed access
type TypedRequest[T any] struct {
	Attribute *RequestAttribute `json:"Attribute,omitempty"`
	Request   T                 `json:"Request,omitempty"`
}

// NewTypedRequest creates a valid TypedRequest with the given attribute and request.
func NewTypedRequest[T any](
	attribute *RequestAttribute,
	request T,
) *TypedRequest[T] {
	return &TypedRequest[T]{
		Attribute: attribute,
		Request:   request,
	}
}

// InvalidTypedRequest creates an invalid TypedRequest with a zero-value request.
func InvalidTypedRequest[T any](
	attribute *RequestAttribute,
) *TypedRequest[T] {
	if attribute == nil {
		attribute = InvalidRequestAttribute()
	}

	return &TypedRequest[T]{
		Attribute: attribute,
	}
}

// Clone returns a deep copy of the TypedRequest.
func (it *TypedRequest[T]) Clone() *TypedRequest[T] {
	if it == nil {
		return nil
	}

	return &TypedRequest[T]{
		Attribute: it.Attribute.Clone(),
		Request:   it.Request,
	}
}

// ToTypedSimpleGenericRequest converts to TypedSimpleGenericRequest[T]
// by wrapping the typed request in a TypedSimpleRequest[T].
func (it *TypedRequest[T]) ToTypedSimpleGenericRequest(
	isValid bool,
	invalidMessage string,
) *TypedSimpleGenericRequest[T] {
	if it == nil {
		return nil
	}

	return &TypedSimpleGenericRequest[T]{
		Attribute: it.Attribute,
		Request: coredynamic.NewTypedSimpleRequest[T](
			it.Request,
			isValid,
			invalidMessage),
	}
}
