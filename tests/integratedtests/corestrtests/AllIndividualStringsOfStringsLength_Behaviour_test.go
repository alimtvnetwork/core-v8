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

package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args")

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ══════════════════════════════════════════════════════════════════════════════

func Test_AllIndividualStringsOfStringsLength_Nil_FromSeg1(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_Nil_FromSeg1", func() {
		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(nil)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- nil ptr", actual)
	})
}

func Test_AllIndividualStringsOfStringsLength_NilSlice_FromSeg1(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_NilSlice_FromSeg1", func() {
		// Arrange
		var s [][]string

		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- nil inner slice", actual)
	})
}

func Test_AllIndividualStringsOfStringsLength_WithItems_FromSeg1(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_WithItems_FromSeg1", func() {
		// Arrange
		s := [][]string{{"a", "b"}, {"c"}}

		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns total -- multi", actual)
	})
}

func Test_AllIndividualStringsOfStringsLength_Empty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_Empty_FromSeg1", func() {
		// Arrange
		s := [][]string{}

		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&s)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns 0 -- empty outer", actual)
	})
}

