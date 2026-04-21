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

package corejson

import (
	"encoding/json"
	"errors"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata"
	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/alimtvnetwork/core-v8/internal/reflectinternal"
)

type newResultCreator struct{}

// UnmarshalUsingBytes
//
//	Aka. alias for DeserializeUsingBytes
//
//	Should be used when Result itself is Serialized
//	and save to somewhere and then unmarshal or deserialize
func (it newResultCreator) UnmarshalUsingBytes(
	deserializingBytes []byte,
) *Result {
	return it.DeserializeUsingBytes(deserializingBytes)
}

// DeserializeUsingBytes
//
//	Should be used when Result itself is Serialized
//	and save to somewhere and then unmarshal or deserialize
func (it newResultCreator) DeserializeUsingBytes(
	deserializingBytes []byte,
) *Result {
	empty := it.TypeName(resultTypeName)

	err := Deserialize.
		UsingBytes(deserializingBytes, empty)

	if err == nil {
		return empty
	}

	empty.Error = err

	return empty
}

func (it newResultCreator) DeserializeUsingResult(
	jsonResult *Result,
) *Result {
	if jsonResult.HasIssuesOrEmpty() {
		return it.ErrorPtr(jsonResult.Error)
	}

	empty := it.TypeName(resultTypeName)

	err := Deserialize.
		UsingBytes(
			jsonResult.SafeBytes(),
			empty)

	if err == nil {
		return empty
	}

	empty.Error = err

	return empty
}

func (it newResultCreator) UsingBytes(
	jsonBytes []byte,
) Result {
	return Result{
		Bytes: jsonBytes,
	}
}

func (it newResultCreator) UsingBytesType(
	jsonBytes []byte,
	typeName string,
) Result {
	return Result{
		Bytes:    jsonBytes,
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingBytesTypePtr(
	jsonBytes []byte,
	typeName string,
) *Result {
	return &Result{
		Bytes:    jsonBytes,
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingTypeBytesPtr(
	typeName string,
	jsonBytes []byte,
) *Result {
	return &Result{
		Bytes:    jsonBytes,
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingBytesPtr(
	jsonBytes []byte,
) *Result {
	if jsonBytes == nil {
		return &Result{}
	}

	return &Result{
		Bytes: jsonBytes,
	}
}

func (it newResultCreator) UsingBytesPtrErrPtr(
	jsonBytes []byte, err error, typeName string,
) *Result {
	if jsonBytes == nil {
		return &Result{
			Error:    err,
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingBytesErrPtr(
	jsonBytes []byte, err error, typeName string,
) *Result {
	if len(jsonBytes) == 0 {
		return &Result{
			Bytes:    []byte{},
			Error:    err,
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) PtrUsingStringPtr(
	jsonStringPtr *string,
	typeName string,
) *Result {
	if jsonStringPtr == nil {
		return it.PtrUsingBytesPtr(
			nil,
			errors.New("json string ptr is nil cannot process further"),
			typeName)
	}

	return &Result{
		Bytes:    []byte(*jsonStringPtr),
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingErrorStringPtr(
	err error,
	jsonStringPtr *string,
	typeName string,
) *Result {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	if jsonStringPtr == nil {
		return it.PtrUsingBytesPtr(
			nil,
			errors.New("json string ptr is nil cannot process further"+errMsg),
			typeName)
	}

	return &Result{
		Bytes:    []byte(*jsonStringPtr),
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) Ptr(
	jsonBytes []byte,
	err error,
	typeName string,
) *Result {
	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingJsonBytesTypeError(
	jsonBytes []byte,
	err error,
	typeName string,
) *Result {
	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingJsonBytesError(
	jsonBytes []byte,
	err error,
) *Result {
	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: constants.UnknownType,
	}
}

func (it newResultCreator) UsingTypePlusString(
	typeName string,
	jsonString string,
) *Result {
	return &Result{
		Bytes:    []byte(jsonString),
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingTypePlusStringPtr(
	typeName string,
	jsonStringPtr *string,
) *Result {
	if jsonStringPtr == nil || len(*jsonStringPtr) == 0 {
		return &Result{
			Bytes: []byte{},
		}
	}

	return &Result{
		Bytes:    []byte(*jsonStringPtr),
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingStringWithType(
	jsonString string,
	typeName string,
) *Result {
	return &Result{
		Bytes:    []byte(jsonString),
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingString(
	jsonString string,
) *Result {
	return &Result{
		Bytes:    []byte(jsonString),
		TypeName: constants.UnknownStringType,
	}
}

func (it newResultCreator) UsingStringPtr(
	jsonStringPtr *string,
) *Result {
	if jsonStringPtr == nil || len(*jsonStringPtr) == 0 {
		return &Result{
			Bytes: []byte{},
		}
	}

	return &Result{
		Bytes:    []byte(*jsonStringPtr),
		TypeName: constants.UnknownStringType,
	}
}

func (it newResultCreator) CreatePtr(
	jsonBytes []byte,
	err error,
	typeName string,
) *Result {
	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) NonPtr(
	jsonBytes []byte,
	err error,
	typeName string,
) Result {
	return Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) Create(
	jsonBytes []byte,
	err error,
	typeName string,
) Result {
	return Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) PtrUsingBytesPtr(
	jsonBytes []byte,
	err error,
	typeName string,
) *Result {
	if err != nil {
		return &Result{
			Bytes:    []byte{},
			Error:    err,
			TypeName: typeName,
		}
	}

	if jsonBytes == nil {
		return &Result{
			Bytes:    []byte{},
			Error:    nil,
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    nil,
		TypeName: typeName,
	}
}

// CastingAny
//
//	if already in JsonResult then returns it
func (it newResultCreator) CastingAny(
	castingAnyToJsonResultPtr any,
) *Result {
	return AnyTo.SerializedJsonResult(castingAnyToJsonResultPtr)
}

func (it newResultCreator) Any(
	anyItem any,
) Result {
	jsonBytes, err := json.Marshal(anyItem)
	typeName := reflectinternal.TypeName(anyItem)

	if err != nil {
		return Result{
			Bytes: jsonBytes,
			Error: errcore.MarshallingFailedType.Error(
				err.Error(),
				typeName),
			TypeName: typeName,
		}
	}

	return Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) AnyPtr(
	anyItem any,
) *Result {
	jsonBytes, err := json.Marshal(anyItem)
	typeName := reflectinternal.TypeName(anyItem)

	if err != nil {
		return &Result{
			Bytes: jsonBytes,
			Error: errcore.MarshallingFailedType.Error(
				err.Error(),
				typeName),
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

// UsingBytesError Get created with nil.
func (it newResultCreator) UsingBytesError(
	bytesError *coredata.BytesError,
) Result {
	if bytesError == nil {
		return Result{}
	}

	return Result{
		Bytes: bytesError.Bytes,
		Error: bytesError.Error,
	}
}

func (it newResultCreator) Error(err error) Result {
	return Result{
		Bytes: nil,
		Error: err,
	}
}

func (it newResultCreator) ErrorPtr(err error) *Result {
	return &Result{
		Bytes: nil,
		Error: err,
	}
}

func (it newResultCreator) Empty() Result {
	return Result{}
}

func (it newResultCreator) EmptyPtr() *Result {
	return &Result{}
}

func (it newResultCreator) TypeName(typeName string) *Result {
	return &Result{
		TypeName: typeName,
	}
}

func (it newResultCreator) TypeNameBytes(typeName string) *Result {
	return &Result{
		TypeName: typeName,
	}
}

func (it newResultCreator) Many(
	anyItems ...any,
) *Result {
	return it.AnyPtr(anyItems)
}

func (it newResultCreator) Serialize(
	anyItem any,
) *Result {
	jsonBytes, err := json.Marshal(anyItem)
	typeName := reflectinternal.TypeName(anyItem)

	if err != nil {
		return &Result{
			Bytes: jsonBytes,
			Error: errcore.MarshallingFailedType.Error(
				err.Error(),
				typeName),
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) Marshal(
	anyItem any,
) *Result {
	jsonBytes, err := json.Marshal(anyItem)
	typeName := reflectinternal.TypeName(anyItem)

	if err != nil {
		return &Result{
			Bytes: jsonBytes,
			Error: errcore.MarshallingFailedType.Error(
				err.Error(),
				typeName),
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func (it newResultCreator) UsingSerializer(
	serializer bytesSerializer,
) *Result {
	if serializer == nil {
		return nil
	}

	allBytes, err := serializer.Serialize()

	return &Result{
		Bytes: allBytes,
		Error: err,
		TypeName: reflectinternal.TypeName(
			serializer),
	}
}

func (it newResultCreator) UsingSerializerFunc(
	serializerFunc func() ([]byte, error),
) *Result {
	if serializerFunc == nil {
		return nil
	}

	allBytes, err := serializerFunc()

	return &Result{
		Bytes:    allBytes,
		Error:    err,
		TypeName: reflectinternal.TypeName(serializerFunc),
	}
}

func (it newResultCreator) UsingJsoner(
	jsoner Jsoner,
) *Result {
	if jsoner == nil {
		return nil
	}

	return jsoner.JsonPtr()
}

// AnyToCastingResult
//
// accepted types (usages anyTo.SerializedJsonResult):
//   - Result, *Result
//   - []byte
//   - string
//   - jsoner
//   - bytesSerializer
//   - anyItem
func (it newResultCreator) AnyToCastingResult(
	anyItem any,
) *Result {
	return AnyTo.SerializedJsonResult(anyItem)
}
