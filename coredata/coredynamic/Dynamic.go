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

package coredynamic

import (
	"fmt"
	"reflect"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata/coreonce"
	"github.com/alimtvnetwork/core-v8/issetter"
)

// Dynamic wraps an arbitrary value with cached reflection metadata.
//
// Getters and type checks are in DynamicGetters.go.
// Reflect-based operations and loops are in DynamicReflect.go.
// JSON serialization/deserialization is in DynamicJson.go.
type Dynamic struct {
	innerData       any
	isValid         bool
	reflectType     reflect.Type
	reflectVal      *reflect.Value
	innerDataString *string
	typeName        coreonce.StringOnce
	length          coreonce.IntegerOnce
	isPointer       issetter.Value
}

// =============================================================================
// Constructors
// =============================================================================

func InvalidDynamic() Dynamic {
	return *InvalidDynamicPtr()
}

func InvalidDynamicPtr() *Dynamic {
	return NewDynamicPtr(
		nil,
		false,
	)
}

func NewDynamicValid(
	data any,
) Dynamic {
	return *NewDynamicPtr(data, true)
}

func NewDynamic(
	data any,
	isValid bool,
) Dynamic {
	return *NewDynamicPtr(data, isValid)
}

func NewDynamicPtr(
	data any,
	isValid bool,
) *Dynamic {
	return &Dynamic{
		innerData: data,
		isValid:   isValid,
		typeName: coreonce.NewStringOnce(
			func() string {
				return fmt.Sprintf(constants.SprintTypeFormat, data)
			},
		),
		length: coreonce.NewIntegerOnce(
			func() int {
				if data == nil {
					return 0
				}

				return LengthOfReflect(reflect.ValueOf(data))
			},
		),
	}
}

// =============================================================================
// Clone
// =============================================================================

func (it Dynamic) Clone() Dynamic {
	return NewDynamic(
		it.innerData,
		it.isValid,
	)
}

func (it *Dynamic) ClonePtr() *Dynamic {
	if it == nil {
		return nil
	}

	return NewDynamicPtr(
		it.innerData,
		it.isValid,
	)
}

func (it Dynamic) NonPtr() Dynamic {
	return it
}

func (it Dynamic) Ptr() *Dynamic {
	return &it
}
