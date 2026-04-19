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

package coreinstructiontests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// ClonePtr
// ==========================================================================

var fromToClonePtrCopiesTestCase = coretestcases.CaseV1{
	Title: "ClonePtr - copies From and To",
	ExpectedInput: args.Map{
		"isNotNil": true,
		"from":     "source",
		"to":       "destination",
	},
}

var fromToClonePtrNilTestCase = coretestcases.CaseV1{
	Title:         "ClonePtr - nil receiver returns nil",
	ExpectedInput: args.Map{"isNil": true},
}

// ==========================================================================
// Clone
// ==========================================================================

var fromToCloneCopiesTestCase = coretestcases.CaseV1{
	Title: "Clone - copies values",
	ExpectedInput: args.Map{
		"from": "a",
		"to":   "b",
	},
}

// ==========================================================================
// IsNull
// ==========================================================================

var fromToIsNullNilTestCase = coretestcases.CaseV1{
	Title:         "IsNull - nil returns true",
	ExpectedInput: args.Map{"result": true},
}

var fromToIsNullNonNilTestCase = coretestcases.CaseV1{
	Title:         "IsNull - non-nil returns false",
	ExpectedInput: args.Map{"result": false},
}

// ==========================================================================
// IsFromEmpty
// ==========================================================================

var fromToIsFromEmptyEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsFromEmpty - empty From returns true",
	ExpectedInput: args.Map{"result": true},
}

var fromToIsFromEmptyNilTestCase = coretestcases.CaseV1{
	Title:         "IsFromEmpty - nil receiver returns true",
	ExpectedInput: args.Map{"result": true},
}

// ==========================================================================
// IsToEmpty
// ==========================================================================

var fromToIsToEmptyEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsToEmpty - empty To returns true",
	ExpectedInput: args.Map{"result": true},
}

var fromToIsToEmptyNonEmptyTestCase = coretestcases.CaseV1{
	Title:         "IsToEmpty - non-empty returns false",
	ExpectedInput: args.Map{"result": false},
}

// ==========================================================================
// String
// ==========================================================================

var fromToStringContainsTestCase = coretestcases.CaseV1{
	Title: "String - contains From and To",
	ExpectedInput: args.Map{
		"containsFrom": true,
		"containsTo":   true,
	},
}

// ==========================================================================
// FromName / ToName
// ==========================================================================

var fromToNamesTestCase = coretestcases.CaseV1{
	Title: "FromName/ToName return field values",
	ExpectedInput: args.Map{
		"fromName": "src",
		"toName":   "dst",
	},
}

// ==========================================================================
// SetFromName
// ==========================================================================

var fromToSetFromNameUpdatesTestCase = coretestcases.CaseV1{
	Title:         "SetFromName - updates From",
	ExpectedInput: args.Map{"from": "new"},
}

var fromToSetFromNameNilTestCase = coretestcases.CaseV1{
	Title:         "SetFromName - nil receiver no panic",
	ExpectedInput: args.Map{"noPanic": true},
}

// ==========================================================================
// SetToName
// ==========================================================================

var fromToSetToNameUpdatesTestCase = coretestcases.CaseV1{
	Title:         "SetToName - updates To",
	ExpectedInput: args.Map{"to": "new"},
}

// ==========================================================================
// SourceDestination
// ==========================================================================

var fromToSourceDestMapsTestCase = coretestcases.CaseV1{
	Title: "SourceDestination - maps From->Source To->Destination",
	ExpectedInput: args.Map{
		"isNotNil":    true,
		"source":      "src",
		"destination": "dst",
	},
}

var fromToSourceDestNilTestCase = coretestcases.CaseV1{
	Title:         "SourceDestination - nil returns nil",
	ExpectedInput: args.Map{"isNil": true},
}

// ==========================================================================
// Rename
// ==========================================================================

var fromToRenameMapsTestCase = coretestcases.CaseV1{
	Title: "Rename - maps From->Existing To->New",
	ExpectedInput: args.Map{
		"isNotNil": true,
		"existing": "old",
		"newName":  "new",
	},
}

var fromToRenameNilTestCase = coretestcases.CaseV1{
	Title:         "Rename - nil returns nil",
	ExpectedInput: args.Map{"isNil": true},
}
