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

package coreinterface

// TypedValueGetter is a generic interface for types that expose a typed Value().
//
// This replaces the need for type-specific getter interfaces like
// ValueIntegerGetter, ValueStringGetter, ValueFloat64Getter, etc.
//
// Existing type-specific getters are retained for backward compatibility
// but new code should prefer TypedValueGetter[T].
//
// Usage:
//
//	type MyConfig struct { val string }
//	func (c MyConfig) Value() string { return c.val }
//	// MyConfig satisfies TypedValueGetter[string]
type TypedValueGetter[T any] interface {
	Value() T
}

// TypedValuesGetter is a generic interface for types that expose
// a typed slice via Values().
//
// Usage:
//
//	type StringList struct { items []string }
//	func (s StringList) Values() []string { return s.items }
//	// StringList satisfies TypedValuesGetter[string]
type TypedValuesGetter[T any] interface {
	Values() []T
}

// TypedKeyValueGetter is a generic interface for key-value accessors.
//
// Usage:
//
//	type Entry struct { k string; v int }
//	func (e Entry) Key() string { return e.k }
//	func (e Entry) Value() int  { return e.v }
//	// Entry satisfies TypedKeyValueGetter[string, int]
type TypedKeyValueGetter[K comparable, V any] interface {
	Key() K
	TypedValueGetter[V]
}
