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

package iserrortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/iserror"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_Equal_BothNonNilSameMessage(t *testing.T) {
	// Arrange
	// Cover the Error() comparison branch
	e1 := errors.New("same")
	e2 := errors.New("same")

	// Act
	actual := args.Map{"result": iserror.Equal(e1, e2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for same message", actual)
}

func Test_QW_Equal_BothNonNilDiffMessage(t *testing.T) {
	// Arrange
	e1 := errors.New("a")
	e2 := errors.New("b")

	// Act
	actual := args.Map{"result": iserror.Equal(e1, e2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for different message", actual)
}

func Test_QW_Equal_LeftNilRightNot(t *testing.T) {
	// Act
	actual := args.Map{"result": iserror.Equal(nil, errors.New("a"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}
