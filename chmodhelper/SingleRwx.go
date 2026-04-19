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

package chmodhelper

import (
	"errors"
	"strings"

	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
)

type SingleRwx struct {
	// Rwx Index Values
	//  - 0: 'r'/'*'/'-'
	//  - 1: 'w'/'*'/'-'
	//  - 2: 'x'/'*'/'-'
	// Examples can be :
	//  - "rwx" or
	//  - "*wx" or
	//  - "rw*" or
	//  - "***"
	//
	// Length must be 3. Not more not less.
	Rwx       string
	ClassType chmodclasstype.Variant
}

func NewSingleRwx(
	rwx string,
	classType chmodclasstype.Variant,
) (*SingleRwx, error) {
	err := GetRwxLengthError(rwx)

	if err != nil {
		return nil, err
	}

	return &SingleRwx{
		Rwx:       rwx,
		ClassType: classType,
	}, nil
}

func (it *SingleRwx) ToRwxOwnerGroupOther() *chmodins.RwxOwnerGroupOther {
	switch it.ClassType {
	case chmodclasstype.All:
		return &chmodins.RwxOwnerGroupOther{
			Owner: it.Rwx,
			Group: it.Rwx,
			Other: it.Rwx,
		}
	case chmodclasstype.Owner:
		return &chmodins.RwxOwnerGroupOther{
			Owner: it.Rwx,
			Group: AllWildcards,
			Other: AllWildcards,
		}
	case chmodclasstype.Group:
		return &chmodins.RwxOwnerGroupOther{
			Owner: AllWildcards,
			Group: it.Rwx,
			Other: AllWildcards,
		}

	case chmodclasstype.Other:
		return &chmodins.RwxOwnerGroupOther{
			Owner: AllWildcards,
			Group: AllWildcards,
			Other: it.Rwx,
		}

	case chmodclasstype.OwnerGroup:
		return &chmodins.RwxOwnerGroupOther{
			Owner: it.Rwx,
			Group: it.Rwx,
			Other: AllWildcards,
		}

	case chmodclasstype.GroupOther:
		return &chmodins.RwxOwnerGroupOther{
			Owner: AllWildcards,
			Group: it.Rwx,
			Other: it.Rwx,
		}

	case chmodclasstype.OwnerOther:
		return &chmodins.RwxOwnerGroupOther{
			Owner: it.Rwx,
			Group: AllWildcards,
			Other: it.Rwx,
		}

	default:
		panic(chmodclasstype.BasicEnumImpl.RangesInvalidErr())
	}
}

func (it *SingleRwx) ToRwxInstruction(
	conditionalIns *chmodins.Condition,
) *chmodins.RwxInstruction {
	rwxOwnerGroupOther := it.ToRwxOwnerGroupOther()

	return &chmodins.RwxInstruction{
		RwxOwnerGroupOther: *rwxOwnerGroupOther,
		Condition:          *conditionalIns,
	}
}

func (it *SingleRwx) ToVarRwxWrapper() (*RwxVariableWrapper, error) {
	rwxOwnerGroupOther := it.ToRwxOwnerGroupOther()

	return ParseRwxOwnerGroupOtherToRwxVariableWrapper(rwxOwnerGroupOther)
}

func (it *SingleRwx) ToDisabledRwxWrapper() (*RwxWrapper, error) {
	rwxOwnerGroupOther := it.ToRwxOwnerGroupOther()
	rwxFullString := rwxOwnerGroupOther.String()
	rwxFullString = strings.ReplaceAll(
		rwxFullString,
		constants.WildcardSymbol,
		constants.Hyphen)

	rwxWrapper, err := New.RwxWrapper.RwxFullString(
		rwxFullString)

	if err != nil {
		return nil, err
	}

	return &rwxWrapper, err
}

func (it *SingleRwx) ToRwxWrapper() (*RwxWrapper, error) {
	if !it.ClassType.IsAll() {
		return nil, errcore.MeaningfulError(
			errcore.CannotConvertToRwxWhereVarRwxPossibleType,
			"ToRwxWrapper",
			errors.New("use ToVarRwx"))
	}

	rwxWrapper, err := New.RwxWrapper.UsingRwxOwnerGroupOther(
		it.ToRwxOwnerGroupOther())

	if err != nil {
		return nil, err
	}

	return &rwxWrapper, err
}

func (it *SingleRwx) ApplyOnMany(
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	toRwxInstruction := it.ToRwxInstruction(condition)
	executor, err := ParseRwxInstructionToExecutor(toRwxInstruction)

	if err != nil {
		return err
	}

	return executor.ApplyOnPathsPtr(&locations)
}
