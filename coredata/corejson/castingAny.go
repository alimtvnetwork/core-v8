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
	"errors"
	"reflect"

	"github.com/alimtvnetwork/core-v8/internal/reflectinternal"
)

type castingAny struct{}

func (it castingAny) FromToDefault(
	fromAny,
	castedToPtr any,
) (failedOrDeserialized error) {
	return it.FromToOption(
		true,
		fromAny,
		castedToPtr,
	)
}

func (it castingAny) FromToReflection(
	fromAny,
	castedToPtr any,
) (failedOrDeserialized error) {
	return it.FromToOption(
		true,
		fromAny,
		castedToPtr,
	)
}

// FromToOption
//
// Giving nil is not support from to.
//
// Warning: must check nil before for from, to both.
//
// Casting from to steps:
//   - reflection first if equal type + right ptr and not nil.
//   - []byte
//   - string
//   - Jsoner
//   - Result
//   - *Result
//   - bytesSerializer
//   - serializerFunc
//   - error to string then cast from json string then to actual unmarshal
func (it castingAny) FromToOption(
	isUseReflection bool,
	fromAny,
	castedToPtr any,
) (failedOrDeserialized error) {
	err, isApplicable := it.reflectionCasting(
		isUseReflection,
		fromAny,
		castedToPtr,
	)
	if isApplicable {
		return err
	}

	switch castedFrom := fromAny.(type) {
	case []byte:
		return Deserialize.UsingBytes(
			castedFrom,
			castedToPtr,
		)
	case string:
		return Deserialize.UsingBytes(
			[]byte(castedFrom),
			castedToPtr,
		)
	case Jsoner:
		jsonResult := castedFrom.Json()

		return jsonResult.Deserialize(castedToPtr)
	case Result:
		return castedFrom.Deserialize(castedToPtr)
	case *Result:
		return castedFrom.Deserialize(castedToPtr)
	case bytesSerializer:
		allBytes, parsingErr := castedFrom.Serialize()

		if parsingErr != nil {
			// usually this error
			// contains all info
			return parsingErr
		}

		return Deserialize.UsingBytes(
			allBytes,
			castedToPtr,
		)
	case func() ([]byte, error): // serializer func
		jsonResult := NewResult.UsingSerializerFunc(
			castedFrom,
		)

		return jsonResult.Deserialize(castedToPtr)
	case error:
		if castedFrom == nil {
			return nil
		}

		parsingErr := Deserialize.UsingBytes(
			[]byte(castedFrom.Error()),
			castedToPtr,
		)

		if parsingErr != nil {
			return errors.New(
				castedFrom.Error() +
					parsingErr.Error(),
			)
		}

		return nil
	}

	// from
	serializeJsonResult := Serialize.Apply(
		fromAny,
	)

	// to
	return serializeJsonResult.Deserialize(
		castedToPtr,
	)
}

// reflectionCasting
//
//	todo refactor return err
func (it castingAny) reflectionCasting(
	isUseReflection bool,
	fromAny any,
	castedToPtr any,
) (err error, isApplicable bool) {
	isSkipReflection := !isUseReflection

	if isSkipReflection {
		return nil, false
	}

	if fromAny == nil || castedToPtr == nil {
		// represents interface nil
		// having type to nil will not be captured here.
		// intentionally not taking it -- not a mistake
		return errors.New(
			"cannot cast from to if any from or to is null",
		), false
	}

	leftType := reflect.TypeOf(fromAny)
	rightType := reflect.TypeOf(castedToPtr)

	if leftType != rightType {
		return nil, false
	}

	isRightPtr := rightType.Kind() == reflect.Ptr

	isNotPtr := !isRightPtr

	if isNotPtr {
		return nil, false
	}

	isLeftDefined := reflectinternal.Is.Defined(fromAny)

	isLeftUndefined := !isLeftDefined

	if isLeftUndefined {
		return nil, false
	}

	isRightDefined := reflectinternal.Is.Defined(castedToPtr)

	isRightUndefined := !isRightDefined

	if isRightUndefined {
		return nil, false
	}

	// ptr, same
	toVal := reflect.
		ValueOf(castedToPtr).
		Elem()
	reflect.
		ValueOf(fromAny).Elem().
		Set(toVal)

	return nil, true
}

func (it castingAny) OrDeserializeTo(
	fromAny,
	castedToPtr any,
) (failedOrDeserialized error) {
	return it.FromToDefault(fromAny, castedToPtr)
}
