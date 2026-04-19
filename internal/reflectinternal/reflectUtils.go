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

package reflectinternal

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type reflectUtils struct{}

func (it reflectUtils) MaxLimit(currentLength, maxCapacity int) int {
	if maxCapacity <= -1 {
		return currentLength
	}

	if currentLength >= maxCapacity {
		return maxCapacity
	}

	return currentLength
}

func (it reflectUtils) AppendArgs(appendingItem any, args []any) []any {
	if len(args) == 0 {
		return []any{appendingItem}
	}

	list := make(
		[]any,
		len(args)+1,
	)

	list[0] = appendingItem

	for i, arg := range args {
		list[i+1] = reflect.ValueOf(arg)
	}

	return list
}

func (it reflectUtils) VerifyReflectTypesAny(left, right []any) (isOkay bool, err error) {
	leftLen := len(left)
	rightLen := len(right)

	if leftLen != rightLen {
		errMsg := fmt.Sprintf(
			"Left Len(%d) != Right Len (%d)",
			leftLen,
			rightLen,
		)

		return false, errors.New(errMsg)
	}

	var errSlice []string

	for i := 0; i < leftLen; i++ {
		l := left[i]
		r := right[i]

		isCurrTypeOkay, errFirst := it.IsReflectTypeMatchAny(l, r)

		if isCurrTypeOkay {
			continue
		}

		if errFirst != nil {
			errSlice = append(
				errSlice,
				it.errorMessageForTypeVerification(i, errFirst),
			)
		}
	}

	if len(errSlice) == 0 {
		return true, nil
	}

	return false, errors.New(strings.Join(errSlice, "\n"))
}

func (it reflectUtils) errorMessageForTypeVerification(i int, errFirst error) string {
	return fmt.Sprintf(
		"- Index {%d}, %s args : %s",
		i,
		indexToPositionFunc(i),
		errFirst.Error(),
	)
}

func (it reflectUtils) VerifyReflectTypes(
	rootName string,
	expectedArgs,
	givenArgs []reflect.Type,
) (isOkay bool, err error) {
	leftLen := len(expectedArgs)
	rightLen := len(givenArgs)

	if leftLen != rightLen {
		errMsg := fmt.Sprintf(
			"Expected Length (%d) != (%d) Given Length",
			leftLen,
			rightLen,
		)

		return false, errors.New(errMsg)
	}

	var errLines []string

	for i := 0; i < leftLen; i++ {
		expected := expectedArgs[i]
		given := givenArgs[i]

		isCurrTypeOkay, errFirst := it.IsReflectTypeMatch(expected, given)

		if isCurrTypeOkay {
			continue
		}

		if errFirst != nil {
			errLines = append(
				errLines,
				it.errorMessageForTypeVerification(i, errFirst),
			)
		}
	}

	if len(errLines) == 0 {
		return true, nil
	}

	finalErrMessage := prependWithSpacesFunc(
		4,
		errLines,
		0,
		fmt.Sprintf("%s =>", rootName),
	)

	return false, errors.New(finalErrMessage)
}

func (it reflectUtils) PkgNameOnly(
	i any,
) string {
	fullName := GetFunc.FullName(i)
	_, pkgName, _ := GetFunc.All(fullName)

	return GetFunc.fixFinalFuncName(pkgName)
}

func (it reflectUtils) FullNameToPkgName(
	fullName string,
) string {
	_, pkgName, _ := GetFunc.All(fullName)

	return GetFunc.fixFinalFuncName(pkgName)
}

func (it reflectUtils) IsReflectTypeMatch(expectedType, givenType reflect.Type) (isOkay bool, err error) {
	if expectedType == givenType {
		return true, nil
	}

	if expectedType.String() == "interface {}" {
		return true, nil
	}

	errMsg := fmt.Sprintf(
		"Expected Type (%s) != (%s) Given Type",
		expectedType.Name(),
		givenType.Name(),
	)

	return false, errors.New(errMsg)
}

func (it reflectUtils) IsReflectTypeMatchAny(expected, given any) (isOkay bool, err error) {
	ex := reflect.TypeOf(expected)
	gi := reflect.TypeOf(given)

	return it.IsReflectTypeMatch(ex, gi)
}
