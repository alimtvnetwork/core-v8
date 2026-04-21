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

package main

import (
	"fmt"

	"github.com/alimtvnetwork/core-v8/coretaskinfo"
	"github.com/alimtvnetwork/core-v8/errcore"
)

func infoCreateExample01() {
	info := coretaskinfo.New.Info.Default(
		"some name",
		"some desc",
		"some url")

	fmt.Println(info.LazyMapPrettyJsonString())

	infoExamples := coretaskinfo.New.Info.Examples(
		"some name",
		"some desc",
		"some url",
		"some examples1 \"some command\"", "some examples 2")

	fmt.Println(infoExamples.LazyMapPrettyJsonString())

	infoNoExamples := coretaskinfo.New.Info.Examples(
		"no exmaple some name",
		"some desc",
		"some url",
	)

	fmt.Println(infoNoExamples.LazyMapPrettyJsonString())

	infoNoExamples2, parseErr := coretaskinfo.New.Info.Deserialized(
		infoNoExamples.JsonPtr().Bytes)

	errcore.HandleErr(parseErr)
	fmt.Println(infoNoExamples2.PrettyJsonStringWithPayloads([]byte("some payloads2")))

	infoNoExamples2 = nil

	fmt.Println(infoNoExamples2.PrettyJsonStringWithPayloads([]byte("some payloads3")))
}
