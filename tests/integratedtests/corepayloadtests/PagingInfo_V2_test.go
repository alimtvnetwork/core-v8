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

package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── PagingInfo ──

func Test_PagingInfo_FromPagingInfoV2(t *testing.T) {
	// Arrange
	p := corepayload.PagingInfo{
		CurrentPageIndex: 1,
		PerPageItems:     10,
		TotalItems:       25,
		TotalPages:       3,
	}

	// Act
	actual := args.Map{
		"totalPages": p.TotalPages,
		"hasTotalPages": p.HasTotalPages(),
		"hasCurrentPage": p.HasCurrentPageIndex(),
	}

	// Assert
	expected := args.Map{
		"totalPages": 3,
		"hasTotalPages": true,
		"hasCurrentPage": true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns correct value -- with args", actual)
}

func Test_PagingInfo_Empty_FromPagingInfoV2(t *testing.T) {
	// Arrange
	p := corepayload.PagingInfo{}

	// Act
	actual := args.Map{
		"isEmpty":    p.IsEmpty(),
		"hasTotalPages": p.HasTotalPages(),
	}

	// Assert
	expected := args.Map{
		"isEmpty":    true,
		"hasTotalPages": false,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns empty -- empty", actual)
}

func Test_PagingInfo_Clone(t *testing.T) {
	// Arrange
	p := corepayload.PagingInfo{
		CurrentPageIndex: 2,
		PerPageItems:     10,
		TotalItems:       50,
		TotalPages:       5,
	}
	cloned := p.Clone()

	// Act
	actual := args.Map{
		"isEqual": p.IsEqual(&cloned),
		"totalPages": cloned.TotalPages,
	}

	// Assert
	expected := args.Map{
		"isEqual": true,
		"totalPages": 5,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns correct value -- clone", actual)
}

func Test_PagingInfo_ClonePtr(t *testing.T) {
	// Arrange
	p := &corepayload.PagingInfo{TotalPages: 3, TotalItems: 30}
	cloned := p.ClonePtr()

	// Act
	actual := args.Map{
		"notNil":  cloned != nil,
		"isEqual": p.IsEqual(cloned),
	}

	// Assert
	expected := args.Map{
		"notNil":  true,
		"isEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns correct value -- clonePtr", actual)
}

func Test_PagingInfo_ClonePtr_Nil_FromPagingInfoV2(t *testing.T) {
	// Arrange
	var p *corepayload.PagingInfo
	cloned := p.ClonePtr()

	// Act
	actual := args.Map{"isNil": cloned == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns nil -- clonePtr nil", actual)
}

func Test_PagingInfo_InvalidChecks(t *testing.T) {
	// Arrange
	p := corepayload.PagingInfo{}

	// Act
	actual := args.Map{
		"invalidTotalPages":       p.IsInvalidTotalPages(),
		"invalidCurrentPageIndex": p.IsInvalidCurrentPageIndex(),
		"invalidPerPageItems":     p.IsInvalidPerPageItems(),
		"invalidTotalItems":       p.IsInvalidTotalItems(),
	}

	// Assert
	expected := args.Map{
		"invalidTotalPages":       true,
		"invalidCurrentPageIndex": true,
		"invalidPerPageItems":     true,
		"invalidTotalItems":       true,
	}
	expected.ShouldBeEqual(t, 0, "PagingInfo returns error -- invalid checks", actual)
}

// ── SessionInfo ──

func Test_SessionInfo_FromPagingInfoV2(t *testing.T) {
	// Arrange
	s := corepayload.SessionInfo{Id: "abc123"}

	// Act
	actual := args.Map{
		"isValid": s.IsValid(),
		"id":      s.Id,
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"id":      "abc123",
	}
	expected.ShouldBeEqual(t, 0, "SessionInfo returns correct value -- with args", actual)
}

func Test_SessionInfo_Empty(t *testing.T) {
	// Arrange
	s := corepayload.SessionInfo{}

	// Act
	actual := args.Map{"isEmpty": s.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "SessionInfo empty -- all zero fields", actual)
}

// ── AuthInfo ──

func Test_AuthInfo_FromPagingInfoV2(t *testing.T) {
	// Arrange
	a := corepayload.AuthInfo{Identifier: "id1", ActionType: "login", ResourceName: "/api"}

	// Act
	actual := args.Map{
		"hasAction":   a.HasActionType(),
		"hasResource": a.HasResourceName(),
		"identifier":  a.Identifier,
	}

	// Assert
	expected := args.Map{
		"hasAction":   true,
		"hasResource": true,
		"identifier":  "id1",
	}
	expected.ShouldBeEqual(t, 0, "AuthInfo returns correct value -- with args", actual)
}

// ── PayloadWrapper ──

func Test_PayloadWrapper_Basic_FromPagingInfoV2(t *testing.T) {
	// Arrange
	pw, _ := corepayload.New.PayloadWrapper.Create(
		"test", "id1", "task", "cat",
		map[string]string{"k": "v"},
	)

	// Act
	actual := args.Map{
		"notNil":  pw != nil,
		"hasAny":  pw.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"notNil":  true,
		"hasAny":  true,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns correct value -- basic", actual)
}

func Test_PayloadWrapper_Empty(t *testing.T) {
	// Arrange
	pw := corepayload.New.PayloadWrapper.Empty()

	// Act
	actual := args.Map{
		"notNil":  pw != nil,
		"hasAny":  pw.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"notNil":  true,
		"hasAny":  false,
	}
	expected.ShouldBeEqual(t, 0, "PayloadWrapper returns empty -- empty", actual)
}
