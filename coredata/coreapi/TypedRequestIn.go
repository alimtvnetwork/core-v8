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

import "github.com/alimtvnetwork/core/coredata/coredynamic"

// TypedRequestIn is the generic API request type.
//
// T represents the strongly-typed request payload.
//
// Usage:
//
//	req := coreapi.NewTypedRequestIn[MyPayload](attr, payload)
//	req.Request.SomeField // fully typed, no assertion needed
type TypedRequestIn[T any] struct {
	Attribute *RequestAttribute `json:"Attribute,omitempty"`
	Request   T                 `json:"Request,omitempty"`
}

// NewTypedRequestIn creates a valid TypedRequestIn with the given attribute and request.
func NewTypedRequestIn[T any](
	attribute *RequestAttribute,
	request T,
) *TypedRequestIn[T] {
	return &TypedRequestIn[T]{
		Attribute: attribute,
		Request:   request,
	}
}

// InvalidTypedRequestIn creates an invalid TypedRequestIn with a zero-value request.
func InvalidTypedRequestIn[T any](
	attribute *RequestAttribute,
) *TypedRequestIn[T] {
	if attribute == nil {
		attribute = InvalidRequestAttribute()
	}

	return &TypedRequestIn[T]{
		Attribute: attribute,
	}
}

// Clone returns a deep copy of the TypedRequestIn.
func (it *TypedRequestIn[T]) Clone() *TypedRequestIn[T] {
	if it == nil {
		return nil
	}

	return &TypedRequestIn[T]{
		Attribute: it.Attribute.Clone(),
		Request:   it.Request,
	}
}

// TypedSimpleGenericRequest converts to a TypedSimpleGenericRequest[T]
// by wrapping the request in a TypedSimpleRequest.
func (it *TypedRequestIn[T]) TypedSimpleGenericRequest(
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
