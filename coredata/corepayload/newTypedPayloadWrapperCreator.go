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
	"github.com/alimtvnetwork/core-v8/internal/reflectinternal"
)

// newTypedPayloadWrapperCreator provides typed factory methods for TypedPayloadWrapper[T].
//
// Because Go does not allow generic methods on non-generic types,
// these are package-level generic functions rather than methods on a creator struct.
// They mirror the newPayloadWrapperCreator API with compile-time type safety.

// TypedPayloadWrapperFrom creates a TypedPayloadWrapper[T] from typed data
// with name, identifier, and entity type.
//
// Usage:
//
//	typed, err := corepayload.TypedPayloadWrapperFrom[User](
//	    "user-create", "usr-123", "User",
//	    User{Name: "Alice"},
//	)
func TypedPayloadWrapperFrom[T any](
	name string,
	identifier string,
	entityType string,
	data T,
) (*TypedPayloadWrapper[T], error) {
	return NewTypedPayloadWrapperFrom[T](name, identifier, entityType, data)
}

// TypedPayloadWrapperRecord creates a TypedPayloadWrapper[T] with auto-detected entity type.
//
// Usage:
//
//	typed, err := corepayload.TypedPayloadWrapperRecord[User](
//	    "user-create", "usr-123", "task", "category",
//	    User{Name: "Alice"},
//	)
func TypedPayloadWrapperRecord[T any](
	name string,
	identifier string,
	taskTypeName string,
	categoryName string,
	data T,
) (*TypedPayloadWrapper[T], error) {
	entityType := reflectinternal.ReflectType.SafeName(data)

	return NewTypedPayloadWrapperFromInstruction[T](
		name,
		identifier,
		taskTypeName,
		entityType,
		categoryName,
		false,
		data,
		nil,
	)
}

// TypedPayloadWrapperRecords creates a TypedPayloadWrapper[T] for multiple records.
//
// Usage:
//
//	typed, err := corepayload.TypedPayloadWrapperRecords[[]User](
//	    "users-list", "batch-1", "task", "category",
//	    users,
//	)
func TypedPayloadWrapperRecords[T any](
	name string,
	identifier string,
	taskTypeName string,
	categoryName string,
	data T,
) (*TypedPayloadWrapper[T], error) {
	entityType := reflectinternal.ReflectType.SafeTypeNameOfSliceOrSingle(false, data)

	return NewTypedPayloadWrapperFromInstruction[T](
		name,
		identifier,
		taskTypeName,
		entityType,
		categoryName,
		true,
		data,
		nil,
	)
}

// TypedPayloadWrapperNameIdRecord creates a TypedPayloadWrapper[T] with name, ID, and data.
//
// Usage:
//
//	typed, err := corepayload.TypedPayloadWrapperNameIdRecord[User](
//	    "user-create", "usr-123",
//	    User{Name: "Alice"},
//	)
func TypedPayloadWrapperNameIdRecord[T any](
	name string,
	identifier string,
	data T,
) (*TypedPayloadWrapper[T], error) {
	entityType := reflectinternal.ReflectType.SafeName(data)

	return NewTypedPayloadWrapperFromInstruction[T](
		name,
		identifier,
		entityType,
		entityType,
		"",
		false,
		data,
		nil,
	)
}

// TypedPayloadWrapperNameIdCategory creates a TypedPayloadWrapper[T] with name, ID, category, and data.
func TypedPayloadWrapperNameIdCategory[T any](
	name string,
	identifier string,
	categoryName string,
	data T,
) (*TypedPayloadWrapper[T], error) {
	entityType := reflectinternal.ReflectType.SafeName(data)

	return NewTypedPayloadWrapperFromInstruction[T](
		name,
		identifier,
		entityType,
		entityType,
		categoryName,
		false,
		data,
		nil,
	)
}

// TypedPayloadWrapperAll creates a TypedPayloadWrapper[T] with all fields specified.
func TypedPayloadWrapperAll[T any](
	name string,
	identifier string,
	taskTypeName string,
	entityTypeName string,
	categoryName string,
	hasManyRecords bool,
	data T,
	attributes *Attributes,
) (*TypedPayloadWrapper[T], error) {
	return NewTypedPayloadWrapperFromInstruction[T](
		name,
		identifier,
		taskTypeName,
		entityTypeName,
		categoryName,
		hasManyRecords,
		data,
		attributes,
	)
}

// TypedPayloadWrapperDeserialize deserializes raw JSON bytes into a TypedPayloadWrapper[T].
//
// First deserializes to PayloadWrapper, then parses Payloads into T.
func TypedPayloadWrapperDeserialize[T any](rawBytes []byte) (*TypedPayloadWrapper[T], error) {
	wrapper, err := New.PayloadWrapper.Deserialize(rawBytes)

	if err != nil {
		return nil, err
	}

	return NewTypedPayloadWrapper[T](wrapper)
}

// TypedPayloadWrapperDeserializeUsingJsonResult deserializes a corejson.Result into TypedPayloadWrapper[T].
func TypedPayloadWrapperDeserializeUsingJsonResult[T any](
	jsonResult *corejson.Result,
) (*TypedPayloadWrapper[T], error) {
	wrapper, err := New.PayloadWrapper.DeserializeUsingJsonResult(jsonResult)

	if err != nil {
		return nil, err
	}

	return NewTypedPayloadWrapper[T](wrapper)
}

// TypedPayloadWrapperDeserializeToMany deserializes raw JSON bytes into a slice
// of TypedPayloadWrapper[T].
func TypedPayloadWrapperDeserializeToMany[T any](
	rawBytes []byte,
) ([]*TypedPayloadWrapper[T], error) {
	wrappers, err := New.PayloadWrapper.DeserializeToMany(rawBytes)

	if err != nil {
		return nil, err
	}

	typedWrappers := make([]*TypedPayloadWrapper[T], 0, len(wrappers))

	for _, wrapper := range wrappers {
		typed, typedErr := NewTypedPayloadWrapper[T](wrapper)

		if typedErr != nil {
			return nil, typedErr
		}

		typedWrappers = append(typedWrappers, typed)
	}

	return typedWrappers, nil
}
