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
	"reflect"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

func CastTo(
	isOutputPointer bool,
	input any,
	acceptedTypes ...reflect.Type,
) CastedResult {
	currentRfType := reflect.TypeOf(input)
	rv := reflect.ValueOf(input)
	kind := rv.Kind()
	var sliceErr []string

	isMatchingAcceptedType := IsAnyTypesOf(
		currentRfType,
		acceptedTypes...,
	)

	isNoMatch := !isMatchingAcceptedType

	if isNoMatch {
		// not matching
		sliceErr = append(
			sliceErr,
			errcore.UnsupportedType.Combine(
				"none matches, current type:"+currentRfType.String(),
				getTypeNamesUsingReflectFunc(true, acceptedTypes...),
			),
		)
	}

	isNull := input == nil || reflectinternal.Is.NullRv(
		rv,
	)
	isOutNonPointer := !isOutputPointer
	hasNonPointerIssue := isNull && isOutNonPointer

	if hasNonPointerIssue {
		sliceErr = append(
			sliceErr,
			errcore.
				InvalidNullPointerType.
				SrcDestination(
					"cannot output non pointer if pointer is null",
					"Value", constants.NilAngelBracket,
					"Type", currentRfType.String(),
				),
		)

		return CastedResult{
			Casted:                 nil,
			SourceReflectType:      currentRfType,
			SourceKind:             kind,
			Error:                  errcore.SliceToError(sliceErr),
			IsNull:                 isNull,
			IsMatchingAcceptedType: isMatchingAcceptedType,
			IsPointer:              isOutNonPointer,
			IsSourcePointer:        kind == reflect.Ptr,
			IsValid:                rv.IsValid(),
		}
	}

	val, _ := PointerOrNonPointerUsingReflectValue(
		isOutputPointer,
		rv,
	)

	return CastedResult{
		Casted:                 val,
		SourceReflectType:      currentRfType,
		SourceKind:             kind,
		Error:                  errcore.SliceToError(sliceErr),
		IsNull:                 isNull,
		IsMatchingAcceptedType: isMatchingAcceptedType,
		IsPointer:              isOutNonPointer,
		IsSourcePointer:        kind == reflect.Ptr,
		IsValid:                rv.IsValid(),
	}
}
