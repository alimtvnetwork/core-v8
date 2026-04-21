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

package typesconv

import (
	"strconv"

	"github.com/alimtvnetwork/core-v8/constants"
)

func StringPtr(val string) *string {
	return &val
}

// StringPtrToSimple if nil then 0
func StringPtrToSimple(val *string) string {
	if val == nil {
		return constants.EmptyString
	}

	return *val
}

// StringPtrToSimpleDef if nil then 0
func StringPtrToSimpleDef(val *string, defVal string) string {
	if val == nil {
		return defVal
	}

	return *val
}

// StringPtrToDefPtr if nil then 0
func StringPtrToDefPtr(val *string, defVal string) *string {
	if val == nil {
		return &defVal
	}

	return val
}

// StringPtrDefValFunc if nil then executes returns defValFunc result as pointer
func StringPtrDefValFunc(val *string, defValFunc func() string) *string {
	if val == nil {
		result := defValFunc()

		return &result
	}

	return val
}

func StringToBool(s string) bool {
	if s == "" {
		return false
	}

	switch s {
	case "yes", "Yes", "YES":
		return true
	case "no", "NO", "No":
		return false
	}

	isBool, err := strconv.ParseBool(s)

	if err != nil {
		return false
	}

	return isBool
}

func StringPointerToBool(s *string) bool {
	if s == nil || *s == "" {
		return false
	}

	return StringToBool(*s)
}

func StringPointerToBoolPtr(s *string) *bool {
	if s == nil || *s == "" {
		toFalse := false

		return &toFalse
	}

	result := StringToBool(*s)

	return &result
}

func StringToBoolPtr(s string) *bool {
	if s == "" {
		toFalse := false

		return &toFalse
	}

	result := StringToBool(s)

	return &result
}
