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

package errcoretests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/alimtvnetwork/core-v8/namevalue"
)

// ══════════════════════════════════════════════════════════════════════════════
// Formatters (Coverage06 - exported only)
// ══════════════════════════════════════════════════════════════════════════════

func Test_VarTwo_WithType(t *testing.T) {
	tc := varTwoTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarTwo(true, "a", 1, "b", 2) != ""})
}

func Test_VarTwo_WithoutType(t *testing.T) {
	tc := varTwoTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarTwo(false, "a", 1, "b", 2) != ""})
}

func Test_VarTwoNoType_Formatterstypes(t *testing.T) {
	tc := varTwoTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarTwoNoType("a", 1, "b", 2) != ""})
}

func Test_VarThree_WithType(t *testing.T) {
	tc := varTwoTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarThree(true, "a", 1, "b", 2, "c", 3) != ""})
}

func Test_VarThree_WithoutType(t *testing.T) {
	tc := varTwoTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarThree(false, "a", 1, "b", 2, "c", 3) != ""})
}

func Test_VarThreeNoType_Formatterstypes(t *testing.T) {
	tc := varTwoTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarThreeNoType("a", 1, "b", 2, "c", 3) != ""})
}

func Test_VarMap_Empty(t *testing.T) {
	tc := varMapTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.VarMap(nil) == ""})
}

func Test_VarMap_WithItems(t *testing.T) {
	tc := varMapTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarMap(map[string]any{"a": 1}) != ""})
}

func Test_VarMapStrings_Empty_Formatterstypes(t *testing.T) {
	tc := varMapTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.VarMapStrings(nil))})
}

func Test_VarMapStrings_WithItems(t *testing.T) {
	tc := varMapTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.VarMapStrings(map[string]any{"a": 1}))})
}

func Test_VarNameValues_Empty_Formatterstypes(t *testing.T) {
	tc := varNameValuesTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.VarNameValues() == ""})
}

func Test_VarNameValues_WithItems(t *testing.T) {
	tc := varNameValuesTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarNameValues(namevalue.StringAny{Name: "a", Value: 1}) != ""})
}

func Test_VarNameValuesJoiner_Empty_Formatterstypes(t *testing.T) {
	tc := varNameValuesTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.VarNameValuesJoiner(",") == ""})
}

func Test_VarNameValuesJoiner_WithItems(t *testing.T) {
	tc := varNameValuesTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.VarNameValuesJoiner(",", namevalue.StringAny{Name: "a", Value: 1}) != ""})
}

func Test_VarNameValuesStrings_Empty_Formatterstypes(t *testing.T) {
	tc := varNameValuesTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.VarNameValuesStrings())})
}

func Test_VarNameValuesStrings_WithItems(t *testing.T) {
	tc := varNameValuesTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.VarNameValuesStrings(namevalue.StringAny{Name: "a", Value: 1}))})
}

func Test_MessageVarTwo_Formatterstypes(t *testing.T) {
	tc := messageVarTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MessageVarTwo("msg", "a", 1, "b", 2) != ""})
}

func Test_MessageVarThree_Formatterstypes(t *testing.T) {
	tc := messageVarTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MessageVarThree("msg", "a", 1, "b", 2, "c", 3) != ""})
}

func Test_MessageVarMap_Empty(t *testing.T) {
	tc := messageVarTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.MessageVarMap("msg", nil)})
}

func Test_MessageVarMap_WithItems(t *testing.T) {
	tc := messageVarTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MessageVarMap("msg", map[string]any{"a": 1}) != ""})
}

func Test_MessageNameValues_Empty_Formatterstypes(t *testing.T) {
	tc := varNameValuesTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.MessageNameValues("msg")})
}

func Test_MessageNameValues_WithItems(t *testing.T) {
	tc := varNameValuesTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MessageNameValues("msg", namevalue.StringAny{Name: "a", Value: 1}) != ""})
}

func Test_MessageWithRef_Formatterstypes(t *testing.T) {
	tc := varNameValuesTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MessageWithRef("msg", "ref") != ""})
}

func Test_MessageWithRefToError_Formatterstypes(t *testing.T) {
	tc := varNameValuesTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.MessageWithRefToError("msg", "ref") != nil})
}

func Test_Ref_Nil(t *testing.T) {
	tc := refTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.Ref(nil) == ""})
}

func Test_Ref_WithRef(t *testing.T) {
	tc := refTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.Ref("ref") != ""})
}

func Test_RefToError_Nil_Formatterstypes(t *testing.T) {
	tc := refTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.RefToError(nil) == nil})
}

func Test_RefToError_WithRef(t *testing.T) {
	tc := refTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.RefToError("ref") != nil})
}

func Test_ToError_Empty(t *testing.T) {
	tc := refTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.ToError("") == nil})
}

func Test_ToError_WithMsg(t *testing.T) {
	tc := refTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ToError("msg") != nil})
}

func Test_ToString_Nil(t *testing.T) {
	tc := refTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.ToString(nil)})
}

func Test_ToString_WithErr(t *testing.T) {
	tc := refTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.ToString(errors.New("e"))})
}

func Test_ToStringPtr_Nil(t *testing.T) {
	tc := toStringPtrTestCases[0]
	p := errcore.ToStringPtr(nil)

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"ptrEmpty": p != nil && *p == ""})
}

func Test_ToStringPtr_WithErr(t *testing.T) {
	tc := toStringPtrTestCases[1]
	p := errcore.ToStringPtr(errors.New("e"))

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"ptrValue": *p})
}

func Test_ToValueString_Formatterstypes(t *testing.T) {
	tc := toStringPtrTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.ToValueString("hello") != ""})
}

func Test_ToExitError_Nil_Formatterstypes(t *testing.T) {
	tc := toStringPtrTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.ToExitError(nil) == nil})
}

func Test_ToExitError_NonExit(t *testing.T) {
	tc := toStringPtrTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.ToExitError(errors.New("e")) == nil})
}

func Test_SourceDestination_Formatterstypes(t *testing.T) {
	tc := sourceDestTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"nonEmpty1": errcore.SourceDestination(true, "src", "dst") != "",
		"nonEmpty2": errcore.SourceDestination(false, "src", "dst") != "",
	})
}

func Test_SourceDestinationNoType_Formatterstypes(t *testing.T) {
	tc := sourceDestTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.SourceDestinationNoType("src", "dst") != ""})
}

func Test_SourceDestinationErr_Formatterstypes(t *testing.T) {
	tc := sourceDestTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.SourceDestinationErr(true, "src", "dst") != nil})
}

func Test_Combine_Formatterstypes(t *testing.T) {
	tc := combineTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.Combine("gen", "other", "ref") != ""})
}

func Test_CombineWithMsgTypeNoStack_Formatterstypes(t *testing.T) {
	tc := combineTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"nonEmpty1": errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "", nil) != "",
		"nonEmpty2": errcore.CombineWithMsgTypeNoStack(errcore.InvalidType, "msg", "ref") != "",
	})
}

func Test_CombineWithMsgTypeStackTrace_Formatterstypes(t *testing.T) {
	tc := combineTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.CombineWithMsgTypeStackTrace(errcore.InvalidType, "msg", nil) != ""})
}

func Test_StackTracesCompiled_Formatterstypes(t *testing.T) {
	tc := combineTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackTracesCompiled([]string{"a", "b"}) != ""})
}

func Test_GherkinsString_Formatterstypes(t *testing.T) {
	tc := combineTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GherkinsString(0, "f", "g", "w", "t") != ""})
}

func Test_GherkinsStringWithExpectation_Formatterstypes(t *testing.T) {
	tc := combineTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GherkinsStringWithExpectation(0, "f", "g", "w", "t", "a", "e") != ""})
}

func Test_RangeNotMeet_WithRange_Formatterstypes(t *testing.T) {
	tc := rangeTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.RangeNotMeet("msg", 0, 10, []int{1, 2}) != ""})
}

func Test_RangeNotMeet_WithoutRange_Formatterstypes(t *testing.T) {
	tc := rangeTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.RangeNotMeet("msg", 0, 10, nil) != ""})
}

func Test_PanicRangeNotMeet_WithRange_Formatterstypes(t *testing.T) {
	tc := rangeTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.PanicRangeNotMeet("msg", 0, 10, []int{1}) != ""})
}

func Test_PanicRangeNotMeet_WithoutRange_Formatterstypes(t *testing.T) {
	tc := rangeTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.PanicRangeNotMeet("msg", 0, 10, nil) != ""})
}

func Test_EnumRangeNotMeet_WithRange_Formatterstypes(t *testing.T) {
	tc := rangeTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.EnumRangeNotMeet(0, 10, "range") != ""})
}

func Test_EnumRangeNotMeet_WithoutRange_Formatterstypes(t *testing.T) {
	tc := rangeTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.EnumRangeNotMeet(0, 10, nil) != ""})
}

func Test_MsgHeader_Formatterstypes(t *testing.T) {
	tc := msgHeaderTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MsgHeader("test") != ""})
}

func Test_MsgHeaderIf_True_Formatterstypes(t *testing.T) {
	tc := msgHeaderTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MsgHeaderIf(true, "test") != ""})
}

func Test_MsgHeaderIf_False_Formatterstypes(t *testing.T) {
	tc := msgHeaderTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MsgHeaderIf(false, "test") != ""})
}

func Test_MsgHeaderPlusEnding_Formatterstypes(t *testing.T) {
	tc := msgHeaderTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MsgHeaderPlusEnding("header", "msg") != ""})
}

func Test_SliceError_Nil(t *testing.T) {
	tc := sliceErrorTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.SliceError(",", nil) == nil})
}

func Test_SliceError_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.SliceError(",", []string{"a", "b"}) != nil})
}

func Test_SliceErrorDefault_Formatterstypes(t *testing.T) {
	tc := sliceErrorTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.SliceErrorDefault([]string{"a"}) != nil})
}

func Test_SliceToError_Nil(t *testing.T) {
	tc := sliceErrorTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.SliceToError(nil) == nil})
}

func Test_SliceToError_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.SliceToError([]string{"a"}) != nil})
}

func Test_SliceToErrorPtr_Nil(t *testing.T) {
	tc := sliceErrorTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.SliceToErrorPtr(nil) == nil})
}

func Test_SliceToErrorPtr_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.SliceToErrorPtr([]string{"a"}) != nil})
}

func Test_SliceErrorsToStrings_Nil(t *testing.T) {
	tc := sliceErrorTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.SliceErrorsToStrings(nil...))})
}

func Test_SliceErrorsToStrings_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.SliceErrorsToStrings(errors.New("a"), nil, errors.New("b")))})
}

func Test_ManyErrorToSingle_Formatterstypes(t *testing.T) {
	tc := sliceErrorTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ManyErrorToSingle([]error{errors.New("a"), nil}) != nil})
}

func Test_ManyErrorToSingleDirect_Formatterstypes(t *testing.T) {
	tc := sliceErrorTestCases[10]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ManyErrorToSingleDirect(errors.New("a")) != nil})
}

func Test_MergeErrors_Formatterstypes(t *testing.T) {
	tc := sliceErrorTestCases[11]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.MergeErrors(errors.New("a"), errors.New("b")) != nil})
}

func Test_MergeErrorsToString_Nil(t *testing.T) {
	tc := sliceErrorTestCases[12]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.MergeErrorsToString(",", nil...) == ""})
}

func Test_MergeErrorsToString_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[13]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MergeErrorsToString(",", errors.New("a")) != ""})
}

func Test_MergeErrorsToStringDefault_Nil(t *testing.T) {
	tc := sliceErrorTestCases[14]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.MergeErrorsToStringDefault(nil...) == ""})
}

func Test_MergeErrorsToStringDefault_WithItems(t *testing.T) {
	tc := sliceErrorTestCases[15]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.MergeErrorsToStringDefault(errors.New("a")) != ""})
}

func Test_StringLinesToQuoteLines_Empty_Formatterstypes(t *testing.T) {
	tc := stringLinesTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.StringLinesToQuoteLines(nil))})
}

func Test_StringLinesToQuoteLines_WithItems(t *testing.T) {
	tc := stringLinesTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.StringLinesToQuoteLines([]string{"a"}))})
}

func Test_StringLinesToQuoteLinesToSingle_Formatterstypes(t *testing.T) {
	tc := stringLinesTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StringLinesToQuoteLinesToSingle([]string{"a", "b"}) != ""})
}

func Test_LinesToDoubleQuoteLinesWithTabs_Empty_Formatterstypes(t *testing.T) {
	tc := stringLinesTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.LinesToDoubleQuoteLinesWithTabs(2, nil))})
}

func Test_LinesToDoubleQuoteLinesWithTabs_WithItems(t *testing.T) {
	tc := stringLinesTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"len": len(errcore.LinesToDoubleQuoteLinesWithTabs(4, []string{"a"}))})
}

func Test_FmtDebug_Formatterstypes(t *testing.T) {
	tc := debugPrintTestCases[0]
	noPanic := !callPanicsErrcore(func() { errcore.FmtDebug("test %s", "v") })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_FmtDebugIf_False_Formatterstypes(t *testing.T) {
	tc := debugPrintTestCases[1]
	noPanic := !callPanicsErrcore(func() { errcore.FmtDebugIf(false, "skip") })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_FmtDebugIf_True_Formatterstypes(t *testing.T) {
	tc := debugPrintTestCases[2]
	noPanic := !callPanicsErrcore(func() { errcore.FmtDebugIf(true, "test %s", "v") })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_ValidPrint(t *testing.T) {
	tc := debugPrintTestCases[3]
	noPanic := !callPanicsErrcore(func() {
		errcore.ValidPrint(false, "skip")
		errcore.ValidPrint(true, "show")
	})

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_FailedPrint(t *testing.T) {
	tc := debugPrintTestCases[4]
	noPanic := !callPanicsErrcore(func() {
		errcore.FailedPrint(false, "skip")
		errcore.FailedPrint(true, "show")
	})

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_GetActualAndExpectProcessedMessage_Formatterstypes(t *testing.T) {
	tc := getActualExpectTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GetActualAndExpectProcessedMessage(0, "a", "e", "ap", "ep") != ""})
}

func Test_GetSearchLineNumberExpectationMessage_Formatterstypes(t *testing.T) {
	tc := getActualExpectTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GetSearchLineNumberExpectationMessage(0, 1, 2, "content", "search", "info") != ""})
}

func Test_GetSearchTermExpectationMessage_WithInfo_Formatterstypes(t *testing.T) {
	tc := getActualExpectTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GetSearchTermExpectationMessage(0, "h", "e", 1, "a", "e", "info") != ""})
}

func Test_GetSearchTermExpectationMessage_NilInfo_Formatterstypes(t *testing.T) {
	tc := getActualExpectTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GetSearchTermExpectationMessage(0, "h", "e", 1, "a", "e", nil) != ""})
}

func Test_GetSearchTermExpectationSimpleMessage_Formatterstypes(t *testing.T) {
	tc := getActualExpectTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.GetSearchTermExpectationSimpleMessage(0, "e", 1, "c", "s") != ""})
}

func Test_Expected_But_Formatterstypes(t *testing.T) {
	tc := expectedButTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.Expected.But("t", "e", "a") != nil})
}

func Test_Expected_ButFoundAsMsg_Formatterstypes(t *testing.T) {
	tc := expectedButTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.Expected.ButFoundAsMsg("t", "e", "a") != ""})
}

func Test_Expected_ButFoundWithTypeAsMsg_Formatterstypes(t *testing.T) {
	tc := expectedButTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.Expected.ButFoundWithTypeAsMsg("t", "e", "a") != ""})
}

func Test_Expected_ButUsingType_Formatterstypes(t *testing.T) {
	tc := expectedButTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.Expected.ButUsingType("t", "e", "a") != nil})
}

func Test_Expected_ReflectButFound_Formatterstypes(t *testing.T) {
	tc := expectedReflectTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.Expected.ReflectButFound(reflect.Int, reflect.String) != nil})
}

func Test_Expected_PrimitiveButFound_Formatterstypes(t *testing.T) {
	tc := expectedReflectTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.Expected.PrimitiveButFound(reflect.Map) != nil})
}

func Test_Expected_ValueHasNoElements_Formatterstypes(t *testing.T) {
	tc := expectedReflectTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.Expected.ValueHasNoElements(reflect.Slice) != nil})
}

func Test_ShouldBe_StrEqMsg_Formatterstypes(t *testing.T) {
	tc := shouldBeTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.ShouldBe.StrEqMsg("a", "b") != ""})
}

func Test_ShouldBe_StrEqErr_Formatterstypes(t *testing.T) {
	tc := shouldBeTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ShouldBe.StrEqErr("a", "b") != nil})
}

func Test_ShouldBe_AnyEqMsg_Formatterstypes(t *testing.T) {
	tc := shouldBeTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.ShouldBe.AnyEqMsg(1, 2) != ""})
}

func Test_ShouldBe_AnyEqErr_Formatterstypes(t *testing.T) {
	tc := shouldBeTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ShouldBe.AnyEqErr(1, 2) != nil})
}

func Test_ShouldBe_JsonEqMsg_Formatterstypes(t *testing.T) {
	tc := shouldBeTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.ShouldBe.JsonEqMsg("a", "b") != ""})
}

func Test_ShouldBe_JsonEqErr_Formatterstypes(t *testing.T) {
	tc := shouldBeTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.ShouldBe.JsonEqErr("a", "b") != nil})
}

// ══════════════════════════════════════════════════════════════════════════════
// RawErrorType (Coverage07)
// ══════════════════════════════════════════════════════════════════════════════

func Test_RawErrorType_String(t *testing.T) {
	tc := rawErrorTypeTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.String() != ""})
}

func Test_RawErrorType_Combine(t *testing.T) {
	tc := rawErrorTypeTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.Combine("msg", "ref") != ""})
}

func Test_RawErrorType_CombineWithAnother_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeTestCases[2]
	r := errcore.InvalidType.CombineWithAnother(errcore.NotFound, "msg", "ref")

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": r.String() != ""})
}

func Test_RawErrorType_TypesAttach_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.TypesAttach("msg", "hello", 42) != ""})
}

func Test_RawErrorType_TypesAttachErr_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.TypesAttachErr("msg", "hello") != nil})
}

func Test_RawErrorType_SrcDestination_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.SrcDestination("msg", "src", "sv", "dst", "dv") != ""})
}

func Test_RawErrorType_SrcDestinationErr_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.SrcDestinationErr("msg", "src", "sv", "dst", "dv") != nil})
}

func Test_RawErrorType_Error_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.Error("msg", "ref") != nil})
}

func Test_RawErrorType_ErrorSkip_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.ErrorSkip(0, "msg", "ref") != nil})
}

func Test_RawErrorType_Fmt_Empty_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.Fmt("") != nil})
}

func Test_RawErrorType_Fmt_WithFormat_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeTestCases[10]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.Fmt("val=%d", 42) != nil})
}

func Test_RawErrorType_FmtIf_False_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeTestCases[11]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.InvalidType.FmtIf(false, "val=%d", 42) == nil})
}

func Test_RawErrorType_FmtIf_True_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeTestCases[12]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.FmtIf(true, "val=%d", 42) != nil})
}

func Test_RawErrorType_MergeError_Nil_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.InvalidType.MergeError(nil) == nil})
}

func Test_RawErrorType_MergeError_WithErr_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.MergeError(errors.New("e")) != nil})
}

func Test_RawErrorType_MergeErrorWithMessage_Nil_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.InvalidType.MergeErrorWithMessage(nil, "msg") == nil})
}

func Test_RawErrorType_MergeErrorWithMessage_WithErr_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.MergeErrorWithMessage(errors.New("e"), "msg") != nil})
}

func Test_RawErrorType_MergeErrorWithMessageRef_Nil_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.InvalidType.MergeErrorWithMessageRef(nil, "msg", "ref") == nil})
}

func Test_RawErrorType_MergeErrorWithMessageRef_WithErr_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.MergeErrorWithMessageRef(errors.New("e"), "msg", "ref") != nil})
}

func Test_RawErrorType_MergeErrorWithRef_Nil_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.InvalidType.MergeErrorWithRef(nil, "ref") == nil})
}

func Test_RawErrorType_MergeErrorWithRef_WithErr_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMergeTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.MergeErrorWithRef(errors.New("e"), "ref") != nil})
}

func Test_RawErrorType_MsgCsvRef_MsgOnly(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.MsgCsvRef("msg") != ""})
}

func Test_RawErrorType_MsgCsvRef_EmptyMsgWithRef(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.MsgCsvRef("", "ref") != ""})
}

func Test_RawErrorType_MsgCsvRef_MsgWithRefs(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.MsgCsvRef("msg", "r1", "r2") != ""})
}

func Test_RawErrorType_MsgCsvRefError_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.MsgCsvRefError("msg", "r1") != nil})
}

func Test_RawErrorType_ErrorRefOnly_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.ErrorRefOnly("ref") != nil})
}

func Test_RawErrorType_Expecting_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.Expecting("exp", "act") != nil})
}

func Test_RawErrorType_NoRef_EmptyMsg(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.NoRef("") != ""})
}

func Test_RawErrorType_NoRef_WithMsg_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.InvalidType.NoRef("msg") != ""})
}

func Test_RawErrorType_ErrorNoRefs_WithMsg(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.ErrorNoRefs("msg") != nil})
}

func Test_RawErrorType_ErrorNoRefs_EmptyMsg(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.ErrorNoRefs("") != nil})
}

func Test_RawErrorType_ErrorNoRefsSkip(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[10]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.InvalidType.ErrorNoRefsSkip(0, "msg") != nil})
}

func Test_RawErrorType_HandleUsingPanic_Formatterstypes(t *testing.T) {
	tc := rawErrorTypeMsgTestCases[11]
	panics := callPanicsErrcore(func() { errcore.InvalidType.HandleUsingPanic("msg", "ref") })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"panics": panics})
}

func Test_GetSet_True_Formatterstypes(t *testing.T) {
	tc := getSetTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.GetSet(true, errcore.InvalidType, errcore.NotFound)})
}

func Test_GetSet_False_Formatterstypes(t *testing.T) {
	tc := getSetTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.GetSet(false, errcore.InvalidType, errcore.NotFound)})
}

func Test_GetSetVariant_True_Formatterstypes(t *testing.T) {
	tc := getSetTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.GetSetVariant(true, "a", "b")})
}

func Test_GetSetVariant_False_Formatterstypes(t *testing.T) {
	tc := getSetTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"result": errcore.GetSetVariant(false, "a", "b")})
}

func Test_MeaningfulError_Nil_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.MeaningfulError(errcore.InvalidType, "fn", nil) == nil})
}

func Test_MeaningfulError_WithErr_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.MeaningfulError(errcore.InvalidType, "fn", errors.New("e")) != nil})
}

func Test_MeaningfulErrorHandle_Nil_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[2]
	noPanic := !callPanicsErrcore(func() { errcore.MeaningfulErrorHandle(errcore.InvalidType, "fn", nil) })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"noPanic": noPanic})
}

func Test_MeaningfulErrorHandle_WithErr_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[3]
	panics := callPanicsErrcore(func() { errcore.MeaningfulErrorHandle(errcore.InvalidType, "fn", errors.New("e")) })

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"panics": panics})
}

func Test_MeaningfulErrorWithData_Nil_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.MeaningfulErrorWithData(errcore.InvalidType, "fn", nil, "data") == nil})
}

func Test_MeaningfulErrorWithData_WithErr_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.MeaningfulErrorWithData(errcore.InvalidType, "fn", errors.New("e"), "data") != nil})
}

func Test_MeaningfulMessageError_Nil_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.MeaningfulMessageError(errcore.InvalidType, "fn", nil, "msg") == nil})
}

func Test_MeaningfulMessageError_WithErr_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.MeaningfulMessageError(errcore.InvalidType, "fn", errors.New("e"), "msg") != nil})
}

func Test_PathMeaningfulMessage_Empty_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.PathMeaningfulMessage(errcore.InvalidType, "fn", "/path") == nil})
}

func Test_PathMeaningfulMessage_WithMsgs_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.PathMeaningfulMessage(errcore.InvalidType, "fn", "/path", "a", "b") != nil})
}

func Test_PathMeaningfulError_Nil_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[10]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.PathMeaningfulError(errcore.InvalidType, nil, "/path") == nil})
}

func Test_PathMeaningfulError_WithErr_Formatterstypes(t *testing.T) {
	tc := meaningfulErrorTestCases[11]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.PathMeaningfulError(errcore.InvalidType, errors.New("e"), "/path") != nil})
}

// ══════════════════════════════════════════════════════════════════════════════
// StackEnhance (Coverage08)
// ══════════════════════════════════════════════════════════════════════════════

func Test_StackEnhance_Error_Nil_Formatterstypes(t *testing.T) {
	tc := stackEnhanceTestCases[0]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.StackEnhance.Error(nil) == nil})
}

func Test_StackEnhance_Error_WithErr_Formatterstypes(t *testing.T) {
	tc := stackEnhanceTestCases[1]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.StackEnhance.Error(errors.New("e")) != nil})
}

func Test_StackEnhance_ErrorSkip_Nil(t *testing.T) {
	tc := stackEnhanceTestCases[2]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.StackEnhance.ErrorSkip(0, nil) == nil})
}

func Test_StackEnhance_ErrorSkip_WithErr(t *testing.T) {
	tc := stackEnhanceTestCases[3]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.StackEnhance.ErrorSkip(0, errors.New("e")) != nil})
}

func Test_StackEnhance_Msg_Empty_Formatterstypes(t *testing.T) {
	tc := stackEnhanceTestCases[4]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.StackEnhance.Msg("") == ""})
}

func Test_StackEnhance_Msg_WithMsg(t *testing.T) {
	tc := stackEnhanceTestCases[5]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackEnhance.Msg("hello") != ""})
}

func Test_StackEnhance_MsgSkip_Empty_Formatterstypes(t *testing.T) {
	tc := stackEnhanceTestCases[6]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.StackEnhance.MsgSkip(0, "") == ""})
}

func Test_StackEnhance_MsgSkip_WithMsg(t *testing.T) {
	tc := stackEnhanceTestCases[7]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackEnhance.MsgSkip(0, "hello") != ""})
}

func Test_StackEnhance_MsgSkip_AlreadyHasStackTrace(t *testing.T) {
	tc := stackEnhanceTestCases[8]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackEnhance.MsgSkip(0, "hello Stack Trace: existing") != ""})
}

func Test_StackEnhance_MsgToErrSkip_Empty_Formatterstypes(t *testing.T) {
	tc := stackEnhanceTestCases[9]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.StackEnhance.MsgToErrSkip(0, "") == nil})
}

func Test_StackEnhance_MsgToErrSkip_WithMsg(t *testing.T) {
	tc := stackEnhanceTestCases[10]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.StackEnhance.MsgToErrSkip(0, "hello") != nil})
}

func Test_StackEnhance_FmtSkip_Empty_Formatterstypes(t *testing.T) {
	tc := stackEnhanceTestCases[11]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.StackEnhance.FmtSkip(0, "") == nil})
}

func Test_StackEnhance_FmtSkip_WithFmt(t *testing.T) {
	tc := stackEnhanceTestCases[12]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.StackEnhance.FmtSkip(0, "hello %s", "world") != nil})
}

func Test_StackEnhance_MsgErrorSkip_NilErr_Formatterstypes(t *testing.T) {
	tc := stackEnhanceTestCases[13]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isEmpty": errcore.StackEnhance.MsgErrorSkip(0, "msg", nil) == ""})
}

func Test_StackEnhance_MsgErrorSkip_WithErr_Formatterstypes(t *testing.T) {
	tc := stackEnhanceTestCases[14]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackEnhance.MsgErrorSkip(0, "msg", errors.New("e")) != ""})
}

func Test_StackEnhance_MsgErrorSkip_AlreadyHasStack(t *testing.T) {
	tc := stackEnhanceTestCases[15]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonEmpty": errcore.StackEnhance.MsgErrorSkip(0, "msg Stack Trace: existing", errors.New("e")) != ""})
}

func Test_StackEnhance_MsgErrorToErrSkip_NilErr_Formatterstypes(t *testing.T) {
	tc := stackEnhanceTestCases[16]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"isNil": errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", nil) == nil})
}

func Test_StackEnhance_MsgErrorToErrSkip_WithErr_Formatterstypes(t *testing.T) {
	tc := stackEnhanceTestCases[17]

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{"nonNil": errcore.StackEnhance.MsgErrorToErrSkip(0, "msg", errors.New("e")) != nil})
}

func Test_CountStateChangeTracker_Formatterstypes(t *testing.T) {
	tc := countStateChangeTrackerTestCases[0]
	c := &errcore.RawErrCollection{}
	tracker := errcore.NewCountStateChangeTracker(c)

	// Assert
	tc.ShouldBeEqualMapFirst(t, args.Map{
		"isSame":           tracker.IsSameState(),
		"isValid":          tracker.IsValid(),
		"isSuccess":        tracker.IsSuccess(),
		"noChanges":        !tracker.HasChanges(),
		"notFailed":        !tracker.IsFailed(),
		"sameWithZero":     tracker.IsSameStateUsingCount(0),
		"changedAfterAdd":  func() bool { c.Add(errors.New("e")); return !tracker.IsSameState() }(),
		"hasChanges":       tracker.HasChanges(),
		"isFailed":         tracker.IsFailed(),
	})
}
