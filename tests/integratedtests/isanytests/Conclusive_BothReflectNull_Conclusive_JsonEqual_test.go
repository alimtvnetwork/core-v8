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

package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/isany"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Conclusive_BothReflectNull(t *testing.T) {
	// Arrange
	var p1 *string
	var p2 *string
	isEqual, isConclusive := isany.Conclusive(p1, p2)

	// Act
	actual := args.Map{"result": isEqual || !isConclusive}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil ptr should be equal conclusive", actual)
}

func Test_Conclusive_OneReflectNull(t *testing.T) {
	// Arrange
	var p1 *string
	s := "hello"
	isEqual, isConclusive := isany.Conclusive(p1, &s)

	// Act
	actual := args.Map{"result": isEqual || !isConclusive}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should be not equal but conclusive", actual)
}

func Test_Conclusive_DiffTypes_FromConclusiveBothReflec(t *testing.T) {
	// Arrange
	isEqual, isConclusive := isany.Conclusive(42, "hello")

	// Act
	actual := args.Map{"result": isEqual || !isConclusive}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "diff types should be not equal but conclusive", actual)
}

func Test_Conclusive_Inconclusive_FromConclusiveBothReflec(t *testing.T) {
	// Arrange
	a := "hello"
	b := "world"
	isEqual, isConclusive := isany.Conclusive(&a, &b)

	// Act
	actual := args.Map{"result": isEqual || isConclusive}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same type diff values should be inconclusive", actual)
}

func Test_JsonEqual_IntEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": isany.JsonEqual(42, 42)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	actual = args.Map{"result": isany.JsonEqual(42, 43)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_JsonEqual_JsonMarshal(t *testing.T) {
	// Arrange
	a := map[string]int{"a": 1}
	b := map[string]int{"a": 1}

	// Act
	actual := args.Map{"result": isany.JsonEqual(a, b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}
