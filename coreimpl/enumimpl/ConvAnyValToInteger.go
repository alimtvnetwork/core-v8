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

package enumimpl

import (
	"fmt"
	"strconv"

	"github.com/alimtvnetwork/core-v8/constants"
)

func ConvEnumAnyValToInteger(val any) int {
	_, isStr := val.(string)

	if isStr {
		// already a string represents string type enum
		return constants.MinInt
	}

	valInt, isInt := val.(int)

	if isInt {
		return valInt
	}

	switch casted := val.(type) {
	case valueByter:
		return int(casted.Value())
	case exactValueByter:
		return int(casted.ValueByte())
	case valueInter:
		return casted.Value()
	case exactValueInter:
		return casted.ValueInt()
	case valueInt8er:
		return int(casted.Value())
	case exactValueInt8er:
		return int(casted.ValueInt8())
	case valueUInt16er:
		return int(casted.Value())
	case exactValueUInt16er:
		return int(casted.ValueUInt16())
	}

	str := fmt.Sprintf(
		constants.SprintValueFormat,
		val)

	convValueInt, err := strconv.Atoi(str)

	if err != nil {
		return constants.MinInt
	}

	return convValueInt
}
