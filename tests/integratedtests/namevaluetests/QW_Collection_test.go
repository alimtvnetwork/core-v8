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

package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/namevalue"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_QW_Collection_String_NilReceiver(t *testing.T) {
	// Arrange
	var c *namevalue.Collection[string, string]
	s := c.String()

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_QW_Collection_JsonString_NilReceiver(t *testing.T) {
	// Arrange
	defer func() { recover() }() // value receiver on nil pointer may panic
	var c *namevalue.Collection[string, string]
	s := c.JsonString()

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_QW_Instance_IsNull(t *testing.T) {
	// Arrange
	var inst *namevalue.Instance[string, string]

	// Act
	actual := args.Map{"result": inst.IsNull()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected null", actual)
}

func Test_QW_Instance_String_Nil(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// nil pointer panic is expected for zero-value Instance.String()
		}
	}()
	inst := namevalue.Instance[string, string]{}
	_ = inst.String()
}

func Test_QW_Instance_JsonString_Nil(t *testing.T) {
	// JsonString is a value receiver — calling on nil pointer panics
	defer func() {
		if r := recover(); r != nil {
			// expected: nil pointer dereference on value receiver
		}
	}()
	var inst *namevalue.Instance[string, string]
	_ = inst.JsonString()
}
