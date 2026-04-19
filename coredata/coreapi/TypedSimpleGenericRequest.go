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

// TypedSimpleGenericRequest is the generic version of SimpleGenericRequest.
//
// It wraps a *coredynamic.TypedSimpleRequest[T] instead of the dynamic *coredynamic.SimpleRequest,
// providing compile-time type safety for the request payload.
//
// Usage:
//
//	type UserInput struct { Name string; Email string }
//
//	req := coreapi.NewTypedSimpleGenericRequest[UserInput](
//	    attr,
//	    coredynamic.NewTypedSimpleRequestValid[UserInput](
//	        UserInput{Name: "Alice", Email: "alice@example.com"},
//	    ),
//	)
//
//	fmt.Println(req.Request.Data().Name) // "Alice" — strongly typed
type TypedSimpleGenericRequest[T any] struct {
	Attribute *RequestAttribute                  `json:"Attribute,omitempty"`
	Request   *coredynamic.TypedSimpleRequest[T] `json:"Request,omitempty"`
}

// NewTypedSimpleGenericRequest creates a valid TypedSimpleGenericRequest.
func NewTypedSimpleGenericRequest[T any](
	attribute *RequestAttribute,
	request *coredynamic.TypedSimpleRequest[T],
) *TypedSimpleGenericRequest[T] {
	return &TypedSimpleGenericRequest[T]{
		Attribute: attribute,
		Request:   request,
	}
}

// InvalidTypedSimpleGenericRequest creates an invalid TypedSimpleGenericRequest
// with a nil request.
func InvalidTypedSimpleGenericRequest[T any](
	attribute *RequestAttribute,
) *TypedSimpleGenericRequest[T] {
	if attribute == nil {
		attribute = InvalidRequestAttribute()
	}

	return &TypedSimpleGenericRequest[T]{
		Attribute: attribute,
		Request:   nil,
	}
}

// IsValid returns true if both attribute and request are valid.
func (it *TypedSimpleGenericRequest[T]) IsValid() bool {
	if it == nil || it.Request == nil {
		return false
	}

	return it.Attribute != nil &&
		it.Attribute.IsValid &&
		it.Request.IsValid()
}

// IsInvalid returns true if the request is invalid.
func (it *TypedSimpleGenericRequest[T]) IsInvalid() bool {
	return !it.IsValid()
}

// Data returns the underlying typed data from the request.
// Panics if Request is nil.
func (it *TypedSimpleGenericRequest[T]) Data() T {
	return it.Request.Data()
}

// Message returns the request message (typically validation/error message).
func (it *TypedSimpleGenericRequest[T]) Message() string {
	if it == nil || it.Request == nil {
		return ""
	}

	return it.Request.Message()
}

// InvalidError returns the request's error if it has one.
func (it *TypedSimpleGenericRequest[T]) InvalidError() error {
	if it == nil || it.Request == nil {
		return nil
	}

	return it.Request.InvalidError()
}

// Clone returns a deep copy of the TypedSimpleGenericRequest.
func (it *TypedSimpleGenericRequest[T]) Clone() *TypedSimpleGenericRequest[T] {
	if it == nil {
		return nil
	}

	return &TypedSimpleGenericRequest[T]{
		Attribute: it.Attribute.Clone(),
		Request:   it.Request.Clone(),
	}
}
