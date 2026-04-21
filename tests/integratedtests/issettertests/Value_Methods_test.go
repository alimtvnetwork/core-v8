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

package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/issetter"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_Value_Methods_Ext2(t *testing.T) {
	// Arrange
	v, err := issetter.New("Set")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)

	// Assert
	actual = args.Map{"result": v.IsUnset()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be unset", actual)
	actual = args.Map{"result": v.IsSet()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be set", actual)
}

func Test_NewBool_Ext2(t *testing.T) {
	// Arrange
	v := issetter.NewBool(true)

	// Assert
	actual := args.Map{"result": v.IsUnset()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be unset", actual)
	actual = args.Map{"result": v.Boolean()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be true", actual)
}

func Test_NewMust_Ext2(t *testing.T) {
	// Act
	v := issetter.NewMust("True")

	// Assert
	actual := args.Map{"result": v.IsUnset()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be unset", actual)
	actual = args.Map{"result": v.IsTrue()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be True", actual)
}

func Test_Max_Ext2(t *testing.T) {
	// Act
	result := issetter.Max()

	// Assert
	actual := args.Map{"result": result != issetter.Wildcard}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Wildcard to be max", actual)
}

func Test_Min_Ext2(t *testing.T) {
	// Act
	result := issetter.Min()

	// Assert
	actual := args.Map{"result": result != issetter.Uninitialized}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Uninitialized to be min", actual)
}
