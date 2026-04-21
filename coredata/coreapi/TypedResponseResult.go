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

// TypedResponseResult is the generic response result type.
//
// T represents the strongly-typed response data.
//
// Usage:
//
//	result := coreapi.NewTypedResponseResult[MyOutput](attr, output)
//	result.Response.Field // strongly typed
type TypedResponseResult[T any] struct {
	Attribute *ResponseAttribute `json:"Attribute,omitempty"`
	Response  T                  `json:"Response,omitempty"`
}

// NewTypedResponseResult creates a valid TypedResponseResult.
func NewTypedResponseResult[T any](
	attribute *ResponseAttribute,
	response T,
) *TypedResponseResult[T] {
	return &TypedResponseResult[T]{
		Attribute: attribute,
		Response:  response,
	}
}

// InvalidTypedResponseResult creates an invalid TypedResponseResult with a zero-value response.
func InvalidTypedResponseResult[T any](
	attribute *ResponseAttribute,
) *TypedResponseResult[T] {
	if attribute == nil {
		attribute = InvalidResponseAttribute(constants.EmptyString)
	}

	return &TypedResponseResult[T]{
		Attribute: attribute,
	}
}

// IsValid returns true if the attribute is present and valid.
func (it *TypedResponseResult[T]) IsValid() bool {
	return it != nil &&
		it.Attribute != nil &&
		it.Attribute.IsValid
}

// IsInvalid returns true if the result is invalid.
func (it *TypedResponseResult[T]) IsInvalid() bool {
	return !it.IsValid()
}

// Message returns the attribute message, or empty string if nil.
func (it *TypedResponseResult[T]) Message() string {
	if it == nil || it.Attribute == nil {
		return constants.EmptyString
	}

	return it.Attribute.Message
}

// Clone returns a deep copy of the TypedResponseResult.
func (it TypedResponseResult[T]) Clone() TypedResponseResult[T] {
	return TypedResponseResult[T]{
		Attribute: it.Attribute.Clone(),
		Response:  it.Response,
	}
}

// ClonePtr returns a pointer to a deep copy.
func (it *TypedResponseResult[T]) ClonePtr() *TypedResponseResult[T] {
	if it == nil {
		return nil
	}

	cloned := it.Clone()

	return &cloned
}

// ToTypedResponse converts to TypedResponse[T].
func (it *TypedResponseResult[T]) ToTypedResponse() *TypedResponse[T] {
	if it == nil {
		return nil
	}

	return &TypedResponse[T]{
		Attribute: it.Attribute,
		Response:  it.Response,
	}
}
