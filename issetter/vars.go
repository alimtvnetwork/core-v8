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

package issetter

import (
	"reflect"

	"github.com/alimtvnetwork/core/internal/csvinternal"
	"github.com/alimtvnetwork/core/simplewrap"
)

var (
	valuesNames = []string{"Uninitialized", "True", "False", "Unset", "Set", "Wildcard"}

	jsonValuesMap = map[string]Value{
		simplewrap.WithDoubleQuote("Uninitialized"): Uninitialized,
		simplewrap.WithDoubleQuote("0"):             Uninitialized,
		simplewrap.WithDoubleQuote(""):              Uninitialized,
		simplewrap.WithDoubleQuote("-"):             Uninitialized,
		simplewrap.WithDoubleQuote("-1"):            Uninitialized,
		simplewrap.WithDoubleQuote("1"):             True,
		simplewrap.WithDoubleQuote("yes"):           True,
		simplewrap.WithDoubleQuote("Yes"):           True,
		simplewrap.WithDoubleQuote("true"):          True,
		simplewrap.WithDoubleQuote("True"):          True,
		simplewrap.WithDoubleQuote("no"):            False,
		simplewrap.WithDoubleQuote("No"):            False,
		simplewrap.WithDoubleQuote("Nop"):           False,
		simplewrap.WithDoubleQuote("None"):          False,
		simplewrap.WithDoubleQuote("false"):         False,
		simplewrap.WithDoubleQuote("False"):         False,
		simplewrap.WithDoubleQuote("set"):           Set,
		simplewrap.WithDoubleQuote("Set"):           Set,
		simplewrap.WithDoubleQuote("Unset"):         Unset,
		simplewrap.WithDoubleQuote("unset"):         Unset,
		simplewrap.WithDoubleQuote("*"):             Wildcard,
		simplewrap.WithDoubleQuote("%"):             Wildcard,
		simplewrap.WithDoubleQuote("Wildcard"):      Wildcard,
		simplewrap.WithDoubleQuote("WildCard"):      Wildcard,
		simplewrap.WithDoubleQuote("wildcard"):      Wildcard, // all small
		"Uninitialized":                        Uninitialized,
		"0":                                    Uninitialized,
		"":                                     Uninitialized,
		"-":                                    Uninitialized,
		"true":                                 True,
		"True":                                 True,
		"yes":                                  True,
		"Yes":                                  True,
		"y":                                    True,
		"Y":                                    True,
		"1":                                    True,
		"false":                                False,
		"False":                                False,
		"no":                                   False,
		"No":                                   False,
		"n":                                    False,
		"N":                                    False,
		"2":                                    True,
		"*":                                    Wildcard,
		"Wildcard":                             Wildcard,
		"wildcard":                             Wildcard,
		"%":                                    Wildcard,
		"set":                                  Set,
		"Set":                                  Set,
		"Unset":                                Unset,
		"unset":                                Unset,
	}

	valuesToJsonBytesMap = map[Value][]byte{
		Uninitialized: jsonBytes("Uninitialized"),
		True:          jsonBytes("True"),
		False:         jsonBytes("False"),
		Unset:         jsonBytes("Unset"),
		Set:           jsonBytes("Set"),
		Wildcard:      jsonBytes("Wildcard"),
	}

	undefinedMap = map[Value]bool{
		Uninitialized: true,
		Wildcard:      true,
	}

	falseMap = map[Value]bool{
		False: true,
		Unset: true,
	}

	trueMap = map[Value]bool{
		True: true,
		Set:  true,
	}

	valuesToNameMap = map[Value]string{
		Uninitialized: "Uninitialized",
		True:          "True",
		False:         "False",
		Unset:         "Unset",
		Set:           "Set",
		Wildcard:      "Wildcard",
	}

	lowerCaseYesNoNames = map[Value]string{
		Uninitialized: "-",
		True:          "yes",
		False:         "no",
		Set:           "yes",
		Unset:         "no",
		Wildcard:      "*",
	}

	yesNoNames = map[Value]string{
		Uninitialized: "-",
		True:          "Yes",
		False:         "No",
		Set:           "Yes",
		Unset:         "No",
		Wildcard:      "*",
	}

	lowerCaseOnOffNames = map[Value]string{
		Uninitialized: "-",
		True:          "on",
		False:         "off",
		Set:           "on",
		Unset:         "off",
		Wildcard:      "*",
	}

	onOffNames = map[Value]string{
		Uninitialized: "-",
		True:          "On",
		False:         "Off",
		Set:           "On",
		Unset:         "Off",
		Wildcard:      "*",
	}

	trueFalseNames = map[Value]string{
		Uninitialized: "-",
		True:          "True",
		False:         "False",
		Set:           "True",
		Unset:         "False",
		Wildcard:      "*",
	}

	trueFalseLowerNames = map[Value]string{
		Uninitialized: "-",
		True:          "true",
		False:         "false",
		Set:           "true",
		Unset:         "false",
		Wildcard:      "*",
	}

	setUnsetLowerNames = map[Value]string{
		Uninitialized: "-",
		True:          "set",
		False:         "unset",
		Set:           "set",
		Unset:         "unset",
		Wildcard:      "*",
	}

	convSetUnsetToTrueFalseMap = map[Value]Value{
		Uninitialized: Uninitialized,
		True:          True,
		False:         False,
		Set:           True,
		Unset:         False,
		Wildcard:      Wildcard,
	}

	convTrueFalseToSetUnsetMap = map[Value]Value{
		Uninitialized: Uninitialized,
		True:          Set,
		False:         Unset,
		Set:           Set,
		Unset:         Unset,
		Wildcard:      Wildcard,
	}

	rangesCsvString = csvinternal.RangeNamesWithValuesIndexesCsvString(
		valuesNames...)

	dynamicRangesMap = generateDynamicRangesMap()
	integerRanges    = IntegerEnumRanges()

	typeName = reflect.TypeOf(Uninitialized).String()
)
