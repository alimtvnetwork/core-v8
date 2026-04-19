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

	"github.com/alimtvnetwork/core/coredata/coredynamic"
)

func jsonResultUnmarshallingTest02() {
	mapAnyItems := getMapAnyItems()
	jsonResult := mapAnyItems.JsonPtr()
	var emptyMapResult *coredynamic.MapAnyItems

	err := jsonResult.Unmarshal(emptyMapResult)
	fmt.Println("err:", err)
	fmt.Println(emptyMapResult)

	var emptyMapResult2 coredynamic.MapAnyItems

	unmarshalValueErr := jsonResult.Unmarshal(emptyMapResult2)
	fmt.Println("unmarshalValueErr:", unmarshalValueErr)
	fmt.Println(emptyMapResult2)

	var emptyMapResult3 coredynamic.MapAnyItems

	unmarshalPtrErr := jsonResult.Unmarshal(&emptyMapResult3)
	fmt.Println("unmarshalPtrErr:", unmarshalPtrErr)
	fmt.Println(emptyMapResult3)

	selfInjectErr := emptyMapResult3.JsonParseSelfInject(jsonResult)
	fmt.Println("selfInjectErr:", selfInjectErr)
	fmt.Println("emptyMapResult3:", emptyMapResult3)

	selfInjectNilErr := emptyMapResult3.JsonParseSelfInject(nil)
	fmt.Println("selfInjectNilErr:", selfInjectNilErr)
	fmt.Println("emptyMapResult3:", emptyMapResult3)

	nilMapInjectErr := emptyMapResult.JsonParseSelfInject(jsonResult)
	fmt.Println("nil emptyMapResult nilMapInjectErr:", nilMapInjectErr)

	jsonResult = nil
	nilJsonUnmarshalErr := jsonResult.Unmarshal(&emptyMapResult3)
	fmt.Println("json Result nil, nilJsonUnmarshalErr:", nilJsonUnmarshalErr)

	jsonResult = nil
	nilBothUnmarshalErr := jsonResult.Unmarshal(emptyMapResult)
	fmt.Println("json Result, object nil, nilBothUnmarshalErr:", nilBothUnmarshalErr)
}
