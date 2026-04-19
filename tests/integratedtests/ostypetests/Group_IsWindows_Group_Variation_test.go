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

package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/ostype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Group_IsWindows(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.WindowsGroup.IsWindows()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected windows", actual)
}

func Test_Group_IsUnix(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.UnixGroup.IsUnix()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected unix", actual)
}

func Test_Group_IsAndroid(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.AndroidGroup.IsAndroid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected android", actual)
}

func Test_Group_IsInvalidGroup(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.InvalidGroup.IsInvalidGroup()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_Variation_Group_Android(t *testing.T) {
	// Arrange
	g := ostype.Android.Group()

	// Act
	actual := args.Map{"result": g.IsAndroid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected android group", actual)
}

func Test_Variation_Group_Unix(t *testing.T) {
	// Arrange
	g := ostype.Linux.Group()

	// Act
	actual := args.Map{"result": g.IsUnix()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected unix group", actual)
}

func Test_Variation_IsActualGroupUnix(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.Linux.IsActualGroupUnix()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected actual group unix", actual)
}

func Test_Variation_IsPossibleUnixGroup(t *testing.T) {
	// Act
	actual := args.Map{"result": ostype.Linux.IsPossibleUnixGroup()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected possible unix", actual)
	actual = args.Map{"result": ostype.Windows.IsPossibleUnixGroup()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "windows should not be unix", actual)
}
