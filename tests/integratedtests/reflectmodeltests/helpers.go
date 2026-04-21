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

package reflectmodeltests

import (
	"reflect"

	"github.com/alimtvnetwork/core-v8/reflectcore/reflectmodel"
)

// sampleStruct is a test helper struct used to create
// real reflect.Method and reflect.StructField instances.
type sampleStruct struct {
	Name    string
	Age     int
	Active  bool
	private string
}

func (s sampleStruct) PublicMethod(a string, b int) (string, error) {
	return a, nil
}

func (s sampleStruct) NoArgsMethod() string {
	return "hello"
}

func (s sampleStruct) MultiReturn() (int, string, error) {
	return 0, "", nil
}

// newFieldProcessor creates a FieldProcessor from a sampleStruct field by name.
func newFieldProcessor(fieldName string, index int) *reflectmodel.FieldProcessor {
	t := reflect.TypeOf(sampleStruct{})
	field, ok := t.FieldByName(fieldName)
	if !ok {
		return nil
	}

	return &reflectmodel.FieldProcessor{
		Name:      field.Name,
		Index:     index,
		Field:     field,
		FieldType: field.Type,
	}
}

// newMethodProcessor creates a MethodProcessor from sampleStruct by method name.
func newMethodProcessor(methodName string) *reflectmodel.MethodProcessor {
	t := reflect.TypeOf(sampleStruct{})

	method, ok := t.MethodByName(methodName)
	if !ok {
		return nil
	}

	return &reflectmodel.MethodProcessor{
		Name:          method.Name,
		Index:         method.Index,
		ReflectMethod: method,
	}
}
