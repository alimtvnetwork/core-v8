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

package corecomparator

var (
	CompareNames = [...]string{
		Equal:            "Equal",
		LeftGreater:      "LeftGreater",
		LeftGreaterEqual: "LeftGreaterEqual",
		LeftLess:         "LeftLess",
		LeftLessEqual:    "LeftLessEqual",
		NotEqual:         "NotEqual",
		Inconclusive:     "Inconclusive",
	}

	CompareOperatorsSymbols = [...]string{
		Equal:            "=",
		LeftGreater:      ">",
		LeftGreaterEqual: ">=",
		LeftLess:         "<",
		LeftLessEqual:    "<=",
		NotEqual:         "!=",
		Inconclusive:     "?!",
	}

	CompareOperatorsShotNames = [...]string{
		Equal:            "eq",
		LeftGreater:      "gt",
		LeftGreaterEqual: "ge",
		LeftLess:         "lt",
		LeftLessEqual:    "le",
		NotEqual:         "ne",
		Inconclusive:     "i",
	}

	SqlCompareOperators = [...]string{
		Equal:            "=",
		LeftGreater:      ">",
		LeftGreaterEqual: ">=",
		LeftLess:         "<",
		LeftLessEqual:    "<=",
		NotEqual:         "<>",
		Inconclusive:     "i",
	}

	RangesMap = map[string]Compare{
		Equal.Name():                         Equal,
		LeftGreater.Name():                   LeftGreater,
		LeftGreaterEqual.Name():              LeftGreaterEqual,
		LeftLess.Name():                      LeftLess,
		LeftLessEqual.Name():                 LeftLessEqual,
		NotEqual.Name():                      NotEqual,
		Inconclusive.Name():                  Inconclusive,
		Equal.NumberString():                 Equal,
		LeftGreater.NumberString():           LeftGreater,
		LeftGreaterEqual.NumberString():      LeftGreaterEqual,
		LeftLess.NumberString():              LeftLess,
		LeftLessEqual.NumberString():         LeftLessEqual,
		NotEqual.NumberString():              NotEqual,
		Inconclusive.NumberString():          Inconclusive,
		Equal.NumberJsonString():             Equal,
		LeftGreater.NumberJsonString():       LeftGreater,
		LeftGreaterEqual.NumberJsonString():  LeftGreaterEqual,
		LeftLess.NumberJsonString():          LeftLess,
		LeftLessEqual.NumberJsonString():     LeftLessEqual,
		NotEqual.NumberJsonString():          NotEqual,
		Inconclusive.NumberJsonString():      Inconclusive,
		Equal.OperatorSymbol():               Equal,
		LeftGreater.OperatorSymbol():         LeftGreater,
		LeftGreaterEqual.OperatorSymbol():    LeftGreaterEqual,
		LeftLess.OperatorSymbol():            LeftLess,
		LeftLessEqual.OperatorSymbol():       LeftLessEqual,
		NotEqual.OperatorSymbol():            NotEqual,
		Inconclusive.OperatorSymbol():        Inconclusive,
		Equal.OperatorShortForm():            Equal,
		LeftGreater.OperatorShortForm():      LeftGreater,
		LeftGreaterEqual.OperatorShortForm(): LeftGreaterEqual,
		LeftLess.OperatorShortForm():         LeftLess,
		LeftLessEqual.OperatorShortForm():    LeftLessEqual,
		NotEqual.OperatorShortForm():         NotEqual,
		Inconclusive.OperatorShortForm():     Inconclusive,
	}
)
