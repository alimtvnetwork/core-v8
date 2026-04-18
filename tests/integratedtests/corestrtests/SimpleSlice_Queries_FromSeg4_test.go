package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice — Segment 4a: Add, Insert, Accessors, Contains, Index, Length
// ══════════════════════════════════════════════════════════════════════════════

func Test_SS_AddAsTitleValue_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddAsTitleValue_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddAsTitleValue("Name", "Alice")

		// Act
		actual := args.Map{"nonEmpty": s.First() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AddAsTitleValue -- formatted", actual)
	})
}

func Test_SS_AddAsTitleValueIf_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddAsTitleValueIf_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddAsTitleValueIf(true, "Name", "Alice")
		s.AddAsTitleValueIf(false, "Skip", "Bob")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddAsTitleValueIf -- only true", actual)
	})
}

func Test_SS_AddAsCurlyTitleWrap_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddAsCurlyTitleWrap_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddAsCurlyTitleWrap("Key", "Val")

		// Act
		actual := args.Map{"nonEmpty": s.First() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AddAsCurlyTitleWrap -- formatted", actual)
	})
}

func Test_SS_AddAsCurlyTitleWrapIf_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddAsCurlyTitleWrapIf_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddAsCurlyTitleWrapIf(true, "K", "V")
		s.AddAsCurlyTitleWrapIf(false, "S", "X")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddAsCurlyTitleWrapIf -- only true", actual)
	})
}

func Test_SS_AddStruct_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddStruct_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddStruct(false, struct{ Name string }{"Alice"})

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStruct -- 1 item", actual)
	})
}

func Test_SS_AddStruct_Nil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddStruct_Nil_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddStruct(false, nil)

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStruct nil -- no change", actual)
	})
}

func Test_SS_AddPointer_Nil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddPointer_Nil_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddPointer(false, nil)

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddPointer nil -- no change", actual)
	})
}

func Test_SS_AddsIf_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddsIf_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddsIf(true, "a", "b").AddsIf(false, "c")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsIf -- only true batch", actual)
	})
}

func Test_SS_AddError_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddError_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddError(errors.New("err")).AddError(nil)

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddError -- only non-nil", actual)
	})
}

func Test_SS_First_Last_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_First_Last_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "c"}

		// Act
		actual := args.Map{
			"first": s.First(),
			"last": s.Last(),
			"firstDyn": s.FirstDynamic(),
			"lastDyn": s.LastDynamic(),
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "c",
			"firstDyn": "a",
			"lastDyn": "c",
		}
		expected.ShouldBeEqual(t, 0, "First/Last/Dynamic -- correct", actual)
	})
}

func Test_SS_FirstOrDefault_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_FirstOrDefault_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{
			"val": s.FirstOrDefault(),
			"dyn": s.FirstOrDefaultDynamic(),
		}

		// Assert
		expected := args.Map{
			"val": "",
			"dyn": "",
		}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault empty -- empty string", actual)
	})
}

func Test_SS_LastOrDefault_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_LastOrDefault_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{
			"val": s.LastOrDefault(),
			"dyn": s.LastOrDefaultDynamic(),
		}

		// Assert
		expected := args.Map{
			"val": "",
			"dyn": "",
		}
		expected.ShouldBeEqual(t, 0, "LastOrDefault empty -- empty string", actual)
	})
}

func Test_SS_AsError_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AsError_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"e1", "e2"}
		err := s.AsError("; ")

		// Act
		actual := args.Map{"notNil": err != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AsError -- non-nil", actual)
	})
}

func Test_SS_AsError_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AsError_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{
			"nil": s.AsError("; ") == nil,
			"defNil": s.AsDefaultError() == nil,
		}

		// Assert
		expected := args.Map{
			"nil": true,
			"defNil": true,
		}
		expected.ShouldBeEqual(t, 0, "AsError empty -- nil", actual)
	})
}

func Test_SS_AsError_Nil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AsError_Nil_FromSeg4", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"nil": s.AsError("; ") == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "AsError nil -- nil", actual)
	})
}

func Test_SS_Skip_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Skip_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "c"}

		// Act
		actual := args.Map{
			"len": len(s.Skip(1)),
			"dynLen": len(s.SkipDynamic(1).([]string)),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"dynLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Skip -- 2 remaining", actual)
	})
}

func Test_SS_Skip_MoreThanLen_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Skip_MoreThanLen_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{"len": len(s.Skip(10))}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Skip more than len -- empty", actual)
	})
}

func Test_SS_Take_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Take_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "c"}

		// Act
		actual := args.Map{"len": len(s.Take(2))}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Take 2 -- 2 items", actual)
	})
}

func Test_SS_Take_MoreThanLen_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Take_MoreThanLen_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{"len": len(s.Take(10))}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Take more than len -- all items", actual)
	})
}

func Test_SS_Limit_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Limit_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "c"}

		// Act
		actual := args.Map{"len": len(s.Limit(2))}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Limit -- delegates to Take", actual)
	})
}

func Test_SS_LimitDynamic_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_LimitDynamic_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "c"}

		// Act
		actual := args.Map{"len": len(s.LimitDynamic(2).([]string))}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LimitDynamic -- delegates to Take", actual)
	})
}

func Test_SS_TakeDynamic_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_TakeDynamic_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		// Fix: TakeDynamic returns []string, not SimpleSlice
		// See issues/corestrtests-takedynamic-type-assertion.md

		// Act
		actual := args.Map{"len": len(s.TakeDynamic(10).([]string))}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "TakeDynamic more -- all items", actual)
	})
}

func Test_SS_Length_Nil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Length_Nil_FromSeg4", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{
			"len": s.Length(),
			"empty": s.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "Length nil -- 0", actual)
	})
}

func Test_SS_Count_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Count_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"count": s.Count()}

		// Assert
		expected := args.Map{"count": 2}
		expected.ShouldBeEqual(t, 0, "Count -- 2", actual)
	})
}

func Test_SS_CountFunc_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_CountFunc_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "bb", "ccc"}

		// Act
		actual := args.Map{"count": s.CountFunc(func(i int, item string) bool { return len(item) > 1 })}

		// Assert
		expected := args.Map{"count": 2}
		expected.ShouldBeEqual(t, 0, "CountFunc -- 2 match", actual)
	})
}

func Test_SS_CountFunc_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_CountFunc_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"count": s.CountFunc(func(i int, item string) bool { return true })}

		// Assert
		expected := args.Map{"count": 0}
		expected.ShouldBeEqual(t, 0, "CountFunc empty -- 0", actual)
	})
}

func Test_SS_HasAnyItem_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_HasAnyItem_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{"has": s.HasAnyItem()}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasAnyItem -- true", actual)
	})
}

func Test_SS_LastIndex_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_LastIndex_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"idx": s.LastIndex()}

		// Assert
		expected := args.Map{"idx": 1}
		expected.ShouldBeEqual(t, 0, "LastIndex -- 1", actual)
	})
}

func Test_SS_HasIndex_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_HasIndex_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{
			"has0": s.HasIndex(0),
			"has5": s.HasIndex(5),
			"hasNeg": s.HasIndex(-1),
		}

		// Assert
		expected := args.Map{
			"has0": true,
			"has5": false,
			"hasNeg": false,
		}
		expected.ShouldBeEqual(t, 0, "HasIndex -- valid and invalid", actual)
	})
}

func Test_SS_IsContains_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsContains_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{
			"has": s.IsContains("a"),
			"miss": s.IsContains("z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContains -- found and missing", actual)
	})
}

func Test_SS_IsContains_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsContains_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"has": s.IsContains("a")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "IsContains empty -- false", actual)
	})
}

func Test_SS_IsContainsFunc_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsContainsFunc_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"hello", "world"}

		// Act
		actual := args.Map{"has": s.IsContainsFunc("hello", func(item, search string) bool { return item == search })}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "IsContainsFunc -- found", actual)
	})
}

func Test_SS_IsContainsFunc_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsContainsFunc_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"has": s.IsContainsFunc("x", func(item, search string) bool { return true })}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "IsContainsFunc empty -- false", actual)
	})
}

func Test_SS_IndexOf_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IndexOf_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "c"}

		// Act
		actual := args.Map{
			"idx": s.IndexOf("b"),
			"miss": s.IndexOf("z"),
		}

		// Assert
		expected := args.Map{
			"idx": 1,
			"miss": -1,
		}
		expected.ShouldBeEqual(t, 0, "IndexOf -- found and missing", actual)
	})
}

func Test_SS_IndexOf_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IndexOf_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"idx": s.IndexOf("a")}

		// Assert
		expected := args.Map{"idx": -1}
		expected.ShouldBeEqual(t, 0, "IndexOf empty -- -1", actual)
	})
}

func Test_SS_IndexOfFunc_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IndexOfFunc_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"aa", "bb"}

		// Act
		actual := args.Map{"idx": s.IndexOfFunc("bb", func(item, search string) bool { return item == search })}

		// Assert
		expected := args.Map{"idx": 1}
		expected.ShouldBeEqual(t, 0, "IndexOfFunc -- found", actual)
	})
}

func Test_SS_IndexOfFunc_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IndexOfFunc_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"idx": s.IndexOfFunc("x", func(item, search string) bool { return true })}

		// Assert
		expected := args.Map{"idx": -1}
		expected.ShouldBeEqual(t, 0, "IndexOfFunc empty -- -1", actual)
	})
}

func Test_SS_StringsList_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_StringsList_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{
			"strLen": len(s.Strings()),
			"listLen": len(s.List()),
		}

		// Assert
		expected := args.Map{
			"strLen": 2,
			"listLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Strings/List -- same", actual)
	})
}

func Test_SS_Hashset_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Hashset_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "a"}

		// Act
		actual := args.Map{"len": s.Hashset().Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset -- 2 unique", actual)
	})
}

func Test_SS_WrapDoubleQuote_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_WrapDoubleQuote_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		w := s.WrapDoubleQuote()

		// Act
		actual := args.Map{"len": w.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "WrapDoubleQuote -- 2 items", actual)
	})
}

func Test_SS_WrapSingleQuote_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_WrapSingleQuote_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		w := s.WrapSingleQuote()

		// Act
		actual := args.Map{"len": w.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapSingleQuote -- 1 item", actual)
	})
}

func Test_SS_WrapTildaQuote_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_WrapTildaQuote_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		w := s.WrapTildaQuote()

		// Act
		actual := args.Map{"len": w.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapTildaQuote -- 1 item", actual)
	})
}

func Test_SS_WrapDoubleQuoteIfMissing_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_WrapDoubleQuoteIfMissing_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		w := s.WrapDoubleQuoteIfMissing()

		// Act
		actual := args.Map{"len": w.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapDoubleQuoteIfMissing -- 1 item", actual)
	})
}

func Test_SS_WrapSingleQuoteIfMissing_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_WrapSingleQuoteIfMissing_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		w := s.WrapSingleQuoteIfMissing()

		// Act
		actual := args.Map{"len": w.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapSingleQuoteIfMissing -- 1 item", actual)
	})
}

func Test_SS_Transpile_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Transpile_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		result := s.Transpile(func(x string) string { return x + "!" })

		// Act
		actual := args.Map{"first": (*result)[0]}

		// Assert
		expected := args.Map{"first": "a!"}
		expected.ShouldBeEqual(t, 0, "Transpile -- modified", actual)
	})
}

func Test_SS_Transpile_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Transpile_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		result := s.Transpile(func(x string) string { return x })

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Transpile empty -- empty", actual)
	})
}

func Test_SS_TranspileJoin_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_TranspileJoin_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		result := s.TranspileJoin(func(x string) string { return x + "!" }, ",")

		// Act
		actual := args.Map{"val": result}

		// Assert
		expected := args.Map{"val": "a!,b!"}
		expected.ShouldBeEqual(t, 0, "TranspileJoin -- joined", actual)
	})
}

func Test_SS_Join_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Join_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"val": s.Join(",")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Join -- comma", actual)
	})
}

func Test_SS_Join_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Join_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"val": s.Join(",")}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Join empty -- empty", actual)
	})
}

func Test_SS_JoinLine_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinLine_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"nonEmpty": s.JoinLine() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinLine -- non-empty", actual)
	})
}

func Test_SS_JoinLine_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinLine_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"val": s.JoinLine()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinLine empty -- empty", actual)
	})
}

func Test_SS_JoinLineEofLine_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinLineEofLine_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		result := s.JoinLineEofLine()

		// Act
		actual := args.Map{"nonEmpty": result != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinLineEofLine -- non-empty", actual)
	})
}

func Test_SS_JoinLineEofLine_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinLineEofLine_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"val": s.JoinLineEofLine()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinLineEofLine empty -- empty", actual)
	})
}

func Test_SS_JoinSpace_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinSpace_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"val": s.JoinSpace()}

		// Assert
		expected := args.Map{"val": "a b"}
		expected.ShouldBeEqual(t, 0, "JoinSpace -- space separated", actual)
	})
}

func Test_SS_JoinComma_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinComma_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"val": s.JoinComma()}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "JoinComma -- comma separated", actual)
	})
}

func Test_SS_JoinWith_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinWith_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"nonEmpty": s.JoinWith(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinWith -- non-empty", actual)
	})
}

func Test_SS_JoinWith_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinWith_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"val": s.JoinWith(",")}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinWith empty -- empty", actual)
	})
}

func Test_SS_JoinCsv_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinCsv_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"nonEmpty": s.JoinCsv() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinCsv -- non-empty", actual)
	})
}

func Test_SS_JoinCsvLine_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinCsvLine_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"nonEmpty": s.JoinCsvLine() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinCsvLine -- non-empty", actual)
	})
}

func Test_SS_JoinCsvString_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinCsvString_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"nonEmpty": s.JoinCsvString(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinCsvString -- non-empty", actual)
	})
}

func Test_SS_JoinCsvString_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JoinCsvString_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"val": s.JoinCsvString(",")}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinCsvString empty -- empty", actual)
	})
}

func Test_SS_EachItemSplitBy_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_EachItemSplitBy_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a,b", "c,d"}
		result := s.EachItemSplitBy(",")

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "EachItemSplitBy -- 4 items", actual)
	})
}

func Test_SS_PrependJoin_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_PrependJoin_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"c", "d"}
		result := s.PrependJoin(",", "a", "b")

		// Act
		actual := args.Map{"nonEmpty": result != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "PrependJoin -- non-empty", actual)
	})
}

func Test_SS_AppendJoin_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AppendJoin_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		result := s.AppendJoin(",", "c", "d")

		// Act
		actual := args.Map{"nonEmpty": result != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "AppendJoin -- non-empty", actual)
	})
}

func Test_SS_PrependAppend_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_PrependAppend_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"b"}
		s.PrependAppend([]string{"a"}, []string{"c"})

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "PrependAppend -- 3 items", actual)
	})
}

func Test_SS_PrependAppend_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_PrependAppend_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"b"}
		s.PrependAppend(nil, nil)

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "PrependAppend empty -- no change", actual)
	})
}

func Test_SS_IsEqual_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqual_FromSeg4", func() {
		// Arrange
		s1 := corestr.SimpleSlice{"a", "b"}
		s2 := corestr.SimpleSlice{"a", "b"}
		s3 := corestr.SimpleSlice{"x"}

		// Act
		actual := args.Map{
			"eq":       s1.IsEqual(&s2),
			"neq":      s1.IsEqual(&s3),
			"nilBoth":  (*corestr.SimpleSlice)(nil).IsEqual(nil),
			"nilOne":   s1.IsEqual(nil),
		}

		// Assert
		expected := args.Map{
			"eq":       true,
			"neq":      false,
			"nilBoth":  true,
			"nilOne":   false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqual -- various cases", actual)
	})
}

func Test_SS_IsEqual_EmptyBoth_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqual_EmptyBoth_FromSeg4", func() {
		// Arrange
		s1 := corestr.SimpleSlice{}
		s2 := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"eq": s1.IsEqual(&s2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqual empty both -- true", actual)
	})
}

func Test_SS_IsEqualLines_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualLines_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{
			"eq":       s.IsEqualLines([]string{"a", "b"}),
			"neq":      s.IsEqualLines([]string{"a", "c"}),
			"diffLen":  s.IsEqualLines([]string{"a"}),
			"nilBoth":  (*corestr.SimpleSlice)(nil).IsEqualLines(nil),
			"nilLeft":  (*corestr.SimpleSlice)(nil).IsEqualLines([]string{"a"}),
		}

		// Assert
		expected := args.Map{
			"eq":       true,
			"neq":      false,
			"diffLen":  false,
			"nilBoth":  true,
			"nilLeft":  false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualLines -- various", actual)
	})
}

func Test_SS_IsEqualUnorderedLines_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualUnorderedLines_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"b", "a"}

		// Act
		actual := args.Map{
			"eq":      s.IsEqualUnorderedLines([]string{"a", "b"}),
			"diffLen": s.IsEqualUnorderedLines([]string{"a"}),
		}

		// Assert
		expected := args.Map{
			"eq":      true,
			"diffLen": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLines -- sorted comparison", actual)
	})
}

func Test_SS_IsEqualUnorderedLines_NilBoth_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualUnorderedLines_NilBoth_FromSeg4", func() {
		// Act
		actual := args.Map{"eq": (*corestr.SimpleSlice)(nil).IsEqualUnorderedLines(nil)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLines nil -- true", actual)
	})
}

func Test_SS_IsEqualUnorderedLinesClone_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualUnorderedLinesClone_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"b", "a"}

		// Act
		actual := args.Map{"eq": s.IsEqualUnorderedLinesClone([]string{"a", "b"})}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLinesClone -- sorted clone comparison", actual)
	})
}

func Test_SS_IsEqualUnorderedLinesClone_NilBoth_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualUnorderedLinesClone_NilBoth_FromSeg4", func() {
		// Act
		actual := args.Map{"eq": (*corestr.SimpleSlice)(nil).IsEqualUnorderedLinesClone(nil)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLinesClone nil -- true", actual)
	})
}

func Test_SS_IsEqualUnorderedLinesClone_DiffLen_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualUnorderedLinesClone_DiffLen_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"eq": s.IsEqualUnorderedLinesClone([]string{"a"})}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLinesClone diff len -- false", actual)
	})
}

func Test_SS_IsEqualUnorderedLinesClone_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualUnorderedLinesClone_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"eq": s.IsEqualUnorderedLinesClone([]string{})}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLinesClone empty -- true", actual)
	})
}

func Test_SS_IsUnorderedEqualRaw_Clone_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsUnorderedEqualRaw_Clone_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"b", "a"}

		// Act
		actual := args.Map{"eq": s.IsUnorderedEqualRaw(true, "a", "b")}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqualRaw clone -- equal", actual)
	})
}

func Test_SS_IsUnorderedEqualRaw_NoClone_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsUnorderedEqualRaw_NoClone_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"b", "a"}

		// Act
		actual := args.Map{"eq": s.IsUnorderedEqualRaw(false, "a", "b")}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqualRaw no clone -- equal", actual)
	})
}

func Test_SS_IsUnorderedEqualRaw_DiffLen_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsUnorderedEqualRaw_DiffLen_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{"eq": s.IsUnorderedEqualRaw(true, "a", "b")}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqualRaw diff len -- false", actual)
	})
}

func Test_SS_IsUnorderedEqualRaw_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsUnorderedEqualRaw_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"eq": s.IsUnorderedEqualRaw(true)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqualRaw empty -- true", actual)
	})
}

func Test_SS_IsUnorderedEqual_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsUnorderedEqual_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"b", "a"}
		r := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"eq": s.IsUnorderedEqual(true, &r)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqual -- equal", actual)
	})
}

func Test_SS_IsUnorderedEqual_BothEmpty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsUnorderedEqual_BothEmpty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		r := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"eq": s.IsUnorderedEqual(true, &r)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqual both empty -- true", actual)
	})
}

func Test_SS_IsUnorderedEqual_NilRight_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsUnorderedEqual_NilRight_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{"eq": s.IsUnorderedEqual(true, nil)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsUnorderedEqual nil right -- false", actual)
	})
}

func Test_SS_IsEqualByFunc_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualByFunc_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{
			"eq":  s.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b"),
			"neq": s.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "c"),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"neq": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualByFunc -- match and mismatch", actual)
	})
}

func Test_SS_IsEqualByFunc_DiffLen_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualByFunc_DiffLen_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{"eq": s.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a", "b")}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualByFunc diff len -- false", actual)
	})
}

func Test_SS_IsEqualByFunc_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualByFunc_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"eq": s.IsEqualByFunc(func(i int, l, r string) bool { return true })}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualByFunc empty -- true", actual)
	})
}

func Test_SS_IsEqualByFuncLinesSplit_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualByFuncLinesSplit_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{
			"eq": s.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return l == r }),
		}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualByFuncLinesSplit -- match", actual)
	})
}

func Test_SS_IsEqualByFuncLinesSplit_Trim_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualByFuncLinesSplit_Trim_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{" a ", " b "}

		// Act
		actual := args.Map{
			"eq": s.IsEqualByFuncLinesSplit(true, ",", "a,b", func(i int, l, r string) bool { return l == r }),
		}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualByFuncLinesSplit trim -- match", actual)
	})
}

func Test_SS_IsEqualByFuncLinesSplit_DiffLen_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualByFuncLinesSplit_DiffLen_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{
			"eq": s.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return true }),
		}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualByFuncLinesSplit diff len -- false", actual)
	})
}

func Test_SS_IsEqualByFuncLinesSplit_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsEqualByFuncLinesSplit_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{
			"eq": s.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true }),
		}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualByFuncLinesSplit empty -- mismatch on split", actual)
	})
}

func Test_SS_IsDistinctEqualRaw_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsDistinctEqualRaw_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "a"}

		// Act
		actual := args.Map{"eq": s.IsDistinctEqualRaw("a", "b")}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsDistinctEqualRaw -- distinct equal", actual)
	})
}

func Test_SS_IsDistinctEqual_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_IsDistinctEqual_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		r := corestr.SimpleSlice{"b", "a"}

		// Act
		actual := args.Map{"eq": s.IsDistinctEqual(&r)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsDistinctEqual -- distinct equal", actual)
	})
}

func Test_SS_Collection_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Collection_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		c := s.Collection(true)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection -- 2 items", actual)
	})
}

func Test_SS_ToCollection_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_ToCollection_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		c := s.ToCollection(false)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ToCollection -- 1 item", actual)
	})
}

func Test_SS_PtrNonPtr_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_PtrNonPtr_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{
			"nonPtrLen": len(s.NonPtr()),
			"ptrLen":    s.Ptr().Length(),
			"toPtrLen":  s.ToPtr().Length(),
			"toNonLen":  len(s.ToNonPtr()),
		}

		// Assert
		expected := args.Map{
			"nonPtrLen": 1,
			"ptrLen":    1,
			"toPtrLen":  1,
			"toNonLen":  1,
		}
		expected.ShouldBeEqual(t, 0, "Ptr/NonPtr -- all same", actual)
	})
}

func Test_SS_String_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_String_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"nonEmpty": s.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_SS_String_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_String_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"val": s.String()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "String empty -- empty", actual)
	})
}

func Test_SS_ConcatNew_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_ConcatNew_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		result := s.ConcatNew("b", "c")

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ConcatNew -- 3 items", actual)
	})
}

func Test_SS_ConcatNewStrings_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_ConcatNewStrings_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		result := s.ConcatNewStrings("b")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNewStrings -- 2 items", actual)
	})
}

func Test_SS_ConcatNewStrings_Nil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_ConcatNewStrings_Nil_FromSeg4", func() {
		// Arrange
		var s *corestr.SimpleSlice
		result := s.ConcatNewStrings("a")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ConcatNewStrings nil -- clone of items", actual)
	})
}

func Test_SS_ConcatNewSimpleSlices_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_ConcatNewSimpleSlices_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		s2 := corestr.SimpleSlice{"b", "c"}
		result := s.ConcatNewSimpleSlices(&s2)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ConcatNewSimpleSlices -- 3 items", actual)
	})
}

func Test_SS_CsvStrings_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_CsvStrings_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}

		// Act
		actual := args.Map{"len": len(s.CsvStrings())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CsvStrings -- 2 quoted", actual)
	})
}

func Test_SS_CsvStrings_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_CsvStrings_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"len": len(s.CsvStrings())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CsvStrings empty -- empty", actual)
	})
}

func Test_SS_Json_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Json_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		j := s.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_SS_MarshalJSON_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_MarshalJSON_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		b, err := s.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_SS_UnmarshalJSON_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_UnmarshalJSON_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		err := s.UnmarshalJSON([]byte(`["a","b"]`))

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": s.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_SS_UnmarshalJSON_Invalid_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_UnmarshalJSON_Invalid_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		err := s.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_SS_Serialize_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Serialize_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		b, err := s.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_SS_Deserialize_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Deserialize_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		var dest []string
		err := s.Deserialize(&dest)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": len(dest),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_SS_ParseInjectUsingJson_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_ParseInjectUsingJson_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		jr := s.JsonPtr()
		s2 := corestr.SimpleSlice{}
		result, err := s2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": result.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- round trip", actual)
	})
}

func Test_SS_ParseInjectUsingJsonMust_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_ParseInjectUsingJsonMust_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		jr := s.JsonPtr()
		s2 := corestr.SimpleSlice{}
		result := s2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_SS_ParseInjectUsingJsonMust_Panic_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_ParseInjectUsingJsonMust_Panic_FromSeg4", func() {
		defer func() { recover() }()
		s := corestr.SimpleSlice{}
		_ = s.ParseInjectUsingJsonMust(&corejson.Result{})
	})
}

func Test_SS_JsonParseSelfInject_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JsonParseSelfInject_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		jr := s.JsonPtr()
		s2 := corestr.SimpleSlice{}
		err := s2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

func Test_SS_JsonModel_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JsonModel_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{"len": len(s.JsonModel())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- returns items", actual)
	})
}

func Test_SS_JsonModelAny_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_JsonModelAny_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{"notNil": s.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_SS_InterfaceCasts_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_InterfaceCasts_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{
			"binder":   s.AsJsonContractsBinder() != nil,
			"jsoner":   s.AsJsoner() != nil,
			"injector": s.AsJsonParseSelfInjector() != nil,
			"marsh":    s.AsJsonMarshaller() != nil,
		}

		// Assert
		expected := args.Map{
			"binder":   true,
			"jsoner":   true,
			"injector": true,
			"marsh":    true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_SS_DeepClone_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_DeepClone_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{"len": s.DeepClone().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DeepClone -- 1 item", actual)
	})
}

func Test_SS_ShadowClone_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_ShadowClone_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{"len": s.ShadowClone().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ShadowClone -- 1 item", actual)
	})
}

func Test_SS_Dispose_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Dispose_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		s.Dispose()

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleared", actual)
	})
}

func Test_SS_Dispose_Nil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Dispose_Nil_FromSeg4", func() {
		var s *corestr.SimpleSlice
		s.Dispose() // should not panic
	})
}

func Test_SS_DistinctDiffRaw_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_DistinctDiffRaw_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "c"}
		diff := s.DistinctDiffRaw("a", "d")

		// Act
		actual := args.Map{"nonEmpty": len(diff) > 0}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "DistinctDiffRaw -- has diff", actual)
	})
}

func Test_SS_DistinctDiffRaw_BothNil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_DistinctDiffRaw_BothNil_FromSeg4", func() {
		// Arrange
		var s *corestr.SimpleSlice
		diff := s.DistinctDiffRaw(nil...)

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "DistinctDiffRaw both nil -- empty", actual)
	})
}

func Test_SS_DistinctDiffRaw_LeftNil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_DistinctDiffRaw_LeftNil_FromSeg4", func() {
		// Arrange
		var s *corestr.SimpleSlice
		diff := s.DistinctDiffRaw("a")

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffRaw left nil -- returns right", actual)
	})
}

func Test_SS_DistinctDiffRaw_RightNil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_DistinctDiffRaw_RightNil_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		diff := s.DistinctDiffRaw(nil...)

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffRaw right nil -- returns left", actual)
	})
}

func Test_SS_DistinctDiff_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_DistinctDiff_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		r := corestr.SimpleSlice{"a", "c"}
		diff := s.DistinctDiff(&r)

		// Act
		actual := args.Map{"nonEmpty": len(diff) > 0}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "DistinctDiff -- has diff", actual)
	})
}

func Test_SS_DistinctDiff_BothNil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_DistinctDiff_BothNil_FromSeg4", func() {
		// Arrange
		var s *corestr.SimpleSlice
		diff := s.DistinctDiff(nil)

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "DistinctDiff both nil -- empty", actual)
	})
}

func Test_SS_DistinctDiff_LeftNil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_DistinctDiff_LeftNil_FromSeg4", func() {
		// Arrange
		var s *corestr.SimpleSlice
		r := corestr.SimpleSlice{"a"}
		diff := s.DistinctDiff(&r)

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiff left nil -- right", actual)
	})
}

func Test_SS_DistinctDiff_RightNil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_DistinctDiff_RightNil_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		diff := s.DistinctDiff(nil)

		// Act
		actual := args.Map{"len": len(diff)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiff right nil -- left", actual)
	})
}

func Test_SS_AddedRemovedLinesDiff_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddedRemovedLinesDiff_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		added, removed := s.AddedRemovedLinesDiff("b", "c")

		// Act
		actual := args.Map{
			"addedLen": len(added),
			"removedLen": len(removed),
		}

		// Assert
		expected := args.Map{
			"addedLen": 1,
			"removedLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "AddedRemovedLinesDiff -- 1 added 1 removed", actual)
	})
}

func Test_SS_AddedRemovedLinesDiff_BothNil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddedRemovedLinesDiff_BothNil_FromSeg4", func() {
		// Arrange
		var s *corestr.SimpleSlice
		added, removed := s.AddedRemovedLinesDiff(nil...)

		// Act
		actual := args.Map{
			"addedNil": added == nil,
			"removedNil": removed == nil,
		}

		// Assert
		expected := args.Map{
			"addedNil": true,
			"removedNil": true,
		}
		expected.ShouldBeEqual(t, 0, "AddedRemovedLinesDiff both nil -- nil", actual)
	})
}

func Test_SS_RemoveIndexes_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_RemoveIndexes_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "c", "d"}
		result, err := s.RemoveIndexes(1, 3)

		// Act
		actual := args.Map{
			"len": result.Length(),
			"noErr": err == nil,
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"noErr": true,
		}
		expected.ShouldBeEqual(t, 0, "RemoveIndexes -- 2 remaining", actual)
	})
}

func Test_SS_RemoveIndexes_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_RemoveIndexes_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		_, err := s.RemoveIndexes(0)

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "RemoveIndexes empty -- error", actual)
	})
}

func Test_SS_RemoveIndexes_OutOfBounds_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_RemoveIndexes_OutOfBounds_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		result, err := s.RemoveIndexes(5)

		// Act
		actual := args.Map{
			"hasErr": err != nil,
			"len": result.Length(),
		}

		// Assert
		expected := args.Map{
			"hasErr": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "RemoveIndexes out of bounds -- error with all items", actual)
	})
}

func Test_SS_SafeStrings_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_SafeStrings_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}

		// Act
		actual := args.Map{"len": len(s.SafeStrings())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "SafeStrings -- 1 item", actual)
	})
}

func Test_SS_SafeStrings_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_SafeStrings_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}

		// Act
		actual := args.Map{"len": len(s.SafeStrings())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeStrings empty -- empty", actual)
	})
}

