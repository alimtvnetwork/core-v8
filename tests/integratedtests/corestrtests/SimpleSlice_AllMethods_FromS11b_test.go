package corestrtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// S11b — SimpleSlice.go Lines 600-1317 — Equal, Clone, Diff, JSON
// ══════════════════════════════════════════════════════════════

func Test_SimpleSlice_66_SimpleSlice_IsEqual_FromS11b(t *testing.T) {
	safeTest(t, "Test_66_SimpleSlice_IsEqual", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		b := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": a.IsEqual(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_67_SimpleSlice_IsEqual_BothNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_67_SimpleSlice_IsEqual_BothNil", func() {
		// Arrange
		var a *corestr.SimpleSlice
		var b *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_68_SimpleSlice_IsEqual_OneNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_68_SimpleSlice_IsEqual_OneNil", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		var b *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_SimpleSlice_69_SimpleSlice_IsEqual_DiffLength_FromS11b(t *testing.T) {
	safeTest(t, "Test_69_SimpleSlice_IsEqual_DiffLength", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		b := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_SimpleSlice_70_SimpleSlice_IsEqual_BothEmpty_FromS11b(t *testing.T) {
	safeTest(t, "Test_70_SimpleSlice_IsEqual_BothEmpty", func() {
		// Arrange
		a := corestr.Empty.SimpleSlice()
		b := corestr.Empty.SimpleSlice()

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_71_SimpleSlice_IsEqualLines_FromS11b(t *testing.T) {
	safeTest(t, "Test_71_SimpleSlice_IsEqualLines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.IsEqualLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		actual = args.Map{"result": ss.IsEqualLines([]string{"a", "c"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_SimpleSlice_72_SimpleSlice_IsEqualLines_BothNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_72_SimpleSlice_IsEqualLines_BothNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": ss.IsEqualLines(nil)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_73_SimpleSlice_IsEqualLines_OneNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_73_SimpleSlice_IsEqualLines_OneNil", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.IsEqualLines(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_SimpleSlice_74_SimpleSlice_IsEqualUnorderedLines_FromS11b(t *testing.T) {
	safeTest(t, "Test_74_SimpleSlice_IsEqualUnorderedLines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_75_SimpleSlice_IsEqualUnorderedLines_BothNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_75_SimpleSlice_IsEqualUnorderedLines_BothNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLines(nil)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_76_SimpleSlice_IsEqualUnorderedLines_DiffLength_FromS11b(t *testing.T) {
	safeTest(t, "Test_76_SimpleSlice_IsEqualUnorderedLines_DiffLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_SimpleSlice_77_SimpleSlice_IsEqualUnorderedLines_BothEmpty_FromS11b(t *testing.T) {
	safeTest(t, "Test_77_SimpleSlice_IsEqualUnorderedLines_BothEmpty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLines([]string{})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_78_SimpleSlice_IsEqualUnorderedLines_Mismatch_FromS11b(t *testing.T) {
	safeTest(t, "Test_78_SimpleSlice_IsEqualUnorderedLines_Mismatch", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLines([]string{"b"})}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_SimpleSlice_79_SimpleSlice_IsEqualUnorderedLinesClone_FromS11b(t *testing.T) {
	safeTest(t, "Test_79_SimpleSlice_IsEqualUnorderedLinesClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_80_SimpleSlice_IsEqualUnorderedLinesClone_BothNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_80_SimpleSlice_IsEqualUnorderedLinesClone_BothNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone(nil)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_81_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength_FromS11b(t *testing.T) {
	safeTest(t, "Test_81_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_SimpleSlice_82_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty_FromS11b(t *testing.T) {
	safeTest(t, "Test_82_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_83_SimpleSlice_IsEqualUnorderedLinesClone_Mismatch_FromS11b(t *testing.T) {
	safeTest(t, "Test_83_SimpleSlice_IsEqualUnorderedLinesClone_Mismatch", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{"b"})}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_SimpleSlice_84_SimpleSlice_Collection_FromS11b(t *testing.T) {
	safeTest(t, "Test_84_SimpleSlice_Collection", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.Collection(true).Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_85_SimpleSlice_NonPtr_Ptr_FromS11b(t *testing.T) {
	safeTest(t, "Test_85_SimpleSlice_NonPtr_Ptr", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		_ = ss.NonPtr()

		// Act
		actual := args.Map{"result": ss.Ptr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_SimpleSlice_86_SimpleSlice_String_FromS11b(t *testing.T) {
	safeTest(t, "Test_86_SimpleSlice_String", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().String() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SimpleSlice_87_SimpleSlice_ConcatNewSimpleSlices_FromS11b(t *testing.T) {
	safeTest(t, "Test_87_SimpleSlice_ConcatNewSimpleSlices", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		b := corestr.New.SimpleSlice.Strings([]string{"b"})
		result := a.ConcatNewSimpleSlices(b)

		// Act
		actual := args.Map{"result": result.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_SimpleSlice_88_SimpleSlice_ConcatNewStrings_FromS11b(t *testing.T) {
	safeTest(t, "Test_88_SimpleSlice_ConcatNewStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.ConcatNewStrings("b")

		// Act
		actual := args.Map{"result": len(result) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_89_SimpleSlice_ConcatNewStrings_Nil_FromS11b(t *testing.T) {
	safeTest(t, "Test_89_SimpleSlice_ConcatNewStrings_Nil", func() {
		// Arrange
		var ss *corestr.SimpleSlice
		result := ss.ConcatNewStrings("b")

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_90_SimpleSlice_ConcatNew_FromS11b(t *testing.T) {
	safeTest(t, "Test_90_SimpleSlice_ConcatNew", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.ConcatNew("b")

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_91_SimpleSlice_ToCollection_FromS11b(t *testing.T) {
	safeTest(t, "Test_91_SimpleSlice_ToCollection", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.ToCollection(false).Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_92_SimpleSlice_CsvStrings_FromS11b(t *testing.T) {
	safeTest(t, "Test_92_SimpleSlice_CsvStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		csv := ss.CsvStrings()

		// Act
		actual := args.Map{"result": len(csv) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().CsvStrings() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_SimpleSlice_93_SimpleSlice_JoinCsvString_FromS11b(t *testing.T) {
	safeTest(t, "Test_93_SimpleSlice_JoinCsvString", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.JoinCsvString(",") == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().JoinCsvString(",") != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SimpleSlice_94_SimpleSlice_JoinWith_FromS11b(t *testing.T) {
	safeTest(t, "Test_94_SimpleSlice_JoinWith", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.JoinWith("|")

		// Act
		actual := args.Map{"result": strings.HasPrefix(result, "|")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected prefix |", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().JoinWith("|") != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SimpleSlice_95_SimpleSlice_JsonModel_FromS11b(t *testing.T) {
	safeTest(t, "Test_95_SimpleSlice_JsonModel", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": len(ss.JsonModel()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_96_SimpleSlice_Sort_FromS11b(t *testing.T) {
	safeTest(t, "Test_96_SimpleSlice_Sort", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"c", "a", "b"})
		ss.Sort()

		// Act
		actual := args.Map{"result": ss.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
	})
}

func Test_SimpleSlice_97_SimpleSlice_Reverse_FromS11b(t *testing.T) {
	safeTest(t, "Test_97_SimpleSlice_Reverse", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		ss.Reverse()

		// Act
		actual := args.Map{"result": ss.First() != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c first", actual)
	})
}

func Test_SimpleSlice_98_SimpleSlice_Reverse_Two_FromS11b(t *testing.T) {
	safeTest(t, "Test_98_SimpleSlice_Reverse_Two", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		ss.Reverse()

		// Act
		actual := args.Map{"result": ss.First() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_SimpleSlice_99_SimpleSlice_Reverse_Single_FromS11b(t *testing.T) {
	safeTest(t, "Test_99_SimpleSlice_Reverse_Single", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		ss.Reverse()

		// Act
		actual := args.Map{"result": ss.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_SimpleSlice_100_SimpleSlice_MarshalJSON_FromS11b(t *testing.T) {
	safeTest(t, "Test_100_SimpleSlice_MarshalJSON", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		data, err := ss.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid JSON", actual)
	})
}

func Test_SimpleSlice_101_SimpleSlice_UnmarshalJSON_FromS11b(t *testing.T) {
	safeTest(t, "Test_101_SimpleSlice_UnmarshalJSON", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()
		err := ss.UnmarshalJSON([]byte(`["a","b"]`))

		// Act
		actual := args.Map{"result": err != nil || ss.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_102_SimpleSlice_UnmarshalJSON_Invalid_FromS11b(t *testing.T) {
	safeTest(t, "Test_102_SimpleSlice_UnmarshalJSON_Invalid", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()
		err := ss.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_SimpleSlice_103_SimpleSlice_Json_FromS11b(t *testing.T) {
	safeTest(t, "Test_103_SimpleSlice_Json", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.Json()

		// Act
		actual := args.Map{"result": jsonResult.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_SimpleSlice_104_SimpleSlice_JsonPtr_FromS11b(t *testing.T) {
	safeTest(t, "Test_104_SimpleSlice_JsonPtr", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_SimpleSlice_105_SimpleSlice_ParseInjectUsingJson_FromS11b(t *testing.T) {
	safeTest(t, "Test_105_SimpleSlice_ParseInjectUsingJson", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.JsonPtr()
		target := corestr.Empty.SimpleSlice()
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Act
		actual := args.Map{"result": err != nil || result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_106_SimpleSlice_ParseInjectUsingJsonMust_FromS11b(t *testing.T) {
	safeTest(t, "Test_106_SimpleSlice_ParseInjectUsingJsonMust", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.JsonPtr()
		target := corestr.Empty.SimpleSlice()
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_107_SimpleSlice_AsJsonContractsBinder_FromS11b(t *testing.T) {
	safeTest(t, "Test_107_SimpleSlice_AsJsonContractsBinder", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.AsJsonContractsBinder() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_SimpleSlice_108_SimpleSlice_AsJsoner_FromS11b(t *testing.T) {
	safeTest(t, "Test_108_SimpleSlice_AsJsoner", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.AsJsoner() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_SimpleSlice_109_SimpleSlice_ToPtr_ToNonPtr_FromS11b(t *testing.T) {
	safeTest(t, "Test_109_SimpleSlice_ToPtr_ToNonPtr", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.ToPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		_ = ss.ToNonPtr()
	})
}

func Test_SimpleSlice_110_SimpleSlice_JsonParseSelfInject_FromS11b(t *testing.T) {
	safeTest(t, "Test_110_SimpleSlice_JsonParseSelfInject", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		jsonResult := ss.JsonPtr()
		target := corestr.Empty.SimpleSlice()
		err := target.JsonParseSelfInject(jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_SimpleSlice_111_SimpleSlice_AsJsonParseSelfInjector_FromS11b(t *testing.T) {
	safeTest(t, "Test_111_SimpleSlice_AsJsonParseSelfInjector", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.AsJsonParseSelfInjector() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_SimpleSlice_112_SimpleSlice_AsJsonMarshaller_FromS11b(t *testing.T) {
	safeTest(t, "Test_112_SimpleSlice_AsJsonMarshaller", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.AsJsonMarshaller() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_SimpleSlice_113_SimpleSlice_JsonModelAny_FromS11b(t *testing.T) {
	safeTest(t, "Test_113_SimpleSlice_JsonModelAny", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_SimpleSlice_114_SimpleSlice_Clear_FromS11b(t *testing.T) {
	safeTest(t, "Test_114_SimpleSlice_Clear", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		ss.Clear()

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_115_SimpleSlice_Clear_Nil_FromS11b(t *testing.T) {
	safeTest(t, "Test_115_SimpleSlice_Clear_Nil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": ss.Clear() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_SimpleSlice_116_SimpleSlice_Dispose_FromS11b(t *testing.T) {
	safeTest(t, "Test_116_SimpleSlice_Dispose", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		ss.Dispose()
	})
}

func Test_SimpleSlice_117_SimpleSlice_Dispose_Nil_FromS11b(t *testing.T) {
	safeTest(t, "Test_117_SimpleSlice_Dispose_Nil", func() {
		var ss *corestr.SimpleSlice
		ss.Dispose()
	})
}

func Test_SimpleSlice_118_SimpleSlice_Clone_FromS11b(t *testing.T) {
	safeTest(t, "Test_118_SimpleSlice_Clone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		cloned := ss.Clone(true)

		// Act
		actual := args.Map{"result": cloned.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_119_SimpleSlice_ClonePtr_FromS11b(t *testing.T) {
	safeTest(t, "Test_119_SimpleSlice_ClonePtr", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.ClonePtr(true).Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_120_SimpleSlice_ClonePtr_Nil_FromS11b(t *testing.T) {
	safeTest(t, "Test_120_SimpleSlice_ClonePtr_Nil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		actual := args.Map{"result": ss.ClonePtr(true) != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_SimpleSlice_121_SimpleSlice_DeepClone_FromS11b(t *testing.T) {
	safeTest(t, "Test_121_SimpleSlice_DeepClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.DeepClone().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_122_SimpleSlice_ShadowClone_FromS11b(t *testing.T) {
	safeTest(t, "Test_122_SimpleSlice_ShadowClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.ShadowClone().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_123_SimpleSlice_IsDistinctEqualRaw_FromS11b(t *testing.T) {
	safeTest(t, "Test_123_SimpleSlice_IsDistinctEqualRaw", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.IsDistinctEqualRaw("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_124_SimpleSlice_IsDistinctEqual_FromS11b(t *testing.T) {
	safeTest(t, "Test_124_SimpleSlice_IsDistinctEqual", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"a"})
		b := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": a.IsDistinctEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_125_SimpleSlice_IsUnorderedEqualRaw_FromS11b(t *testing.T) {
	safeTest(t, "Test_125_SimpleSlice_IsUnorderedEqualRaw", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})

		// Act
		actual := args.Map{"result": ss.IsUnorderedEqualRaw(true, "a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal with clone", actual)
		ss2 := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		actual = args.Map{"result": ss2.IsUnorderedEqualRaw(false, "a", "b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal without clone", actual)
	})
}

func Test_SimpleSlice_126_SimpleSlice_IsUnorderedEqualRaw_DiffLength_FromS11b(t *testing.T) {
	safeTest(t, "Test_126_SimpleSlice_IsUnorderedEqualRaw_DiffLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.IsUnorderedEqualRaw(false, "a", "b")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_SimpleSlice_127_SimpleSlice_IsUnorderedEqualRaw_BothEmpty_FromS11b(t *testing.T) {
	safeTest(t, "Test_127_SimpleSlice_IsUnorderedEqualRaw_BothEmpty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		actual := args.Map{"result": ss.IsUnorderedEqualRaw(false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_128_SimpleSlice_IsUnorderedEqual_FromS11b(t *testing.T) {
	safeTest(t, "Test_128_SimpleSlice_IsUnorderedEqual", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		b := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": a.IsUnorderedEqual(true, b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_129_SimpleSlice_IsUnorderedEqual_BothEmpty_FromS11b(t *testing.T) {
	safeTest(t, "Test_129_SimpleSlice_IsUnorderedEqual_BothEmpty", func() {
		// Arrange
		a := corestr.Empty.SimpleSlice()
		b := corestr.Empty.SimpleSlice()

		// Act
		actual := args.Map{"result": a.IsUnorderedEqual(false, b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_130_SimpleSlice_IsUnorderedEqual_NilRight_FromS11b(t *testing.T) {
	safeTest(t, "Test_130_SimpleSlice_IsUnorderedEqual_NilRight", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": a.IsUnorderedEqual(false, nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_SimpleSlice_131_SimpleSlice_IsEqualByFunc_FromS11b(t *testing.T) {
	safeTest(t, "Test_131_SimpleSlice_IsEqualByFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		result := ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b")

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SimpleSlice_132_SimpleSlice_IsEqualByFunc_DiffLength_FromS11b(t *testing.T) {
	safeTest(t, "Test_132_SimpleSlice_IsEqualByFunc_DiffLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a", "b")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleSlice_133_SimpleSlice_IsEqualByFunc_Empty_FromS11b(t *testing.T) {
	safeTest(t, "Test_133_SimpleSlice_IsEqualByFunc_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		actual := args.Map{"result": ss.IsEqualByFunc(func(i int, l, r string) bool { return true })}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for both empty", actual)
	})
}

func Test_SimpleSlice_134_SimpleSlice_IsEqualByFunc_Mismatch_FromS11b(t *testing.T) {
	safeTest(t, "Test_134_SimpleSlice_IsEqualByFunc_Mismatch", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.IsEqualByFunc(func(i int, l, r string) bool { return false }, "a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleSlice_135_SimpleSlice_IsEqualByFuncLinesSplit_FromS11b(t *testing.T) {
	safeTest(t, "Test_135_SimpleSlice_IsEqualByFuncLinesSplit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		result := ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return l == r })

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SimpleSlice_136_SimpleSlice_IsEqualByFuncLinesSplit_Trim_FromS11b(t *testing.T) {
	safeTest(t, "Test_136_SimpleSlice_IsEqualByFuncLinesSplit_Trim", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{" a ", " b "})
		result := ss.IsEqualByFuncLinesSplit(true, ",", "a,b", func(i int, l, r string) bool { return l == r })

		// Act
		actual := args.Map{"result": result}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true with trim", actual)
	})
}

func Test_SimpleSlice_137_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength_FromS11b(t *testing.T) {
	safeTest(t, "Test_137_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return true })}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleSlice_138_SimpleSlice_IsEqualByFuncLinesSplit_Empty_FromS11b(t *testing.T) {
	safeTest(t, "Test_138_SimpleSlice_IsEqualByFuncLinesSplit_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()
		// strings.Split("", ",") returns [""] (length 1) which != 0, so returns false

		// Act
		actual := args.Map{"result": ss.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true })}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty vs split-empty mismatch", actual)
	})
}

func Test_SimpleSlice_139_SimpleSlice_IsEqualByFuncLinesSplit_Mismatch_FromS11b(t *testing.T) {
	safeTest(t, "Test_139_SimpleSlice_IsEqualByFuncLinesSplit_Mismatch", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.IsEqualByFuncLinesSplit(false, ",", "b", func(i int, l, r string) bool { return l == r })}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleSlice_140_SimpleSlice_DistinctDiffRaw_FromS11b(t *testing.T) {
	safeTest(t, "Test_140_SimpleSlice_DistinctDiffRaw", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		diff := ss.DistinctDiffRaw("b", "c")

		// Act
		actual := args.Map{"result": len(diff) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_141_SimpleSlice_DistinctDiffRaw_BothNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_141_SimpleSlice_DistinctDiffRaw_BothNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice
		diff := ss.DistinctDiffRaw()

		// Act
		actual := args.Map{"result": len(diff) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_142_SimpleSlice_DistinctDiffRaw_LeftNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_142_SimpleSlice_DistinctDiffRaw_LeftNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice
		diff := ss.DistinctDiffRaw("a")

		// Act
		actual := args.Map{"result": len(diff) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_143_SimpleSlice_DistinctDiffRaw_RightNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_143_SimpleSlice_DistinctDiffRaw_RightNil", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		diff := ss.DistinctDiffRaw(nil...)

		// Act
		actual := args.Map{"result": len(diff) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_144_SimpleSlice_AddedRemovedLinesDiff_FromS11b(t *testing.T) {
	safeTest(t, "Test_144_SimpleSlice_AddedRemovedLinesDiff", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		added, removed := ss.AddedRemovedLinesDiff("b", "c")

		// Act
		actual := args.Map{"result": len(added) != 1 || len(removed) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 added 1 removed", actual)
	})
}

func Test_SimpleSlice_145_SimpleSlice_AddedRemovedLinesDiff_BothNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_145_SimpleSlice_AddedRemovedLinesDiff_BothNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice
		added, removed := ss.AddedRemovedLinesDiff()

		// Act
		actual := args.Map{"result": added != nil || removed != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_SimpleSlice_146_SimpleSlice_RemoveIndexes_FromS11b(t *testing.T) {
	safeTest(t, "Test_146_SimpleSlice_RemoveIndexes", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		result, err := ss.RemoveIndexes(1)

		// Act
		actual := args.Map{"result": err != nil || result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_147_SimpleSlice_RemoveIndexes_Empty_FromS11b(t *testing.T) {
	safeTest(t, "Test_147_SimpleSlice_RemoveIndexes_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()
		_, err := ss.RemoveIndexes(0)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_SimpleSlice_148_SimpleSlice_RemoveIndexes_InvalidIndex_FromS11b(t *testing.T) {
	safeTest(t, "Test_148_SimpleSlice_RemoveIndexes_InvalidIndex", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		_, err := ss.RemoveIndexes(5)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for invalid index", actual)
	})
}

func Test_SimpleSlice_149_SimpleSlice_RemoveIndexes_AllRemoved_FromS11b(t *testing.T) {
	safeTest(t, "Test_149_SimpleSlice_RemoveIndexes_AllRemoved", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result, err := ss.RemoveIndexes(0)

		// Act
		actual := args.Map{"result": err != nil || result.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_150_SimpleSlice_DistinctDiff_FromS11b(t *testing.T) {
	safeTest(t, "Test_150_SimpleSlice_DistinctDiff", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		b := corestr.New.SimpleSlice.Strings([]string{"b", "c"})
		diff := a.DistinctDiff(b)

		// Act
		actual := args.Map{"result": len(diff) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_151_SimpleSlice_DistinctDiff_BothNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_151_SimpleSlice_DistinctDiff_BothNil", func() {
		// Arrange
		var a *corestr.SimpleSlice
		diff := a.DistinctDiff(nil)

		// Act
		actual := args.Map{"result": len(diff) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_152_SimpleSlice_DistinctDiff_LeftNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_152_SimpleSlice_DistinctDiff_LeftNil", func() {
		// Arrange
		var a *corestr.SimpleSlice
		b := corestr.New.SimpleSlice.Strings([]string{"x"})
		diff := a.DistinctDiff(b)

		// Act
		actual := args.Map{"result": len(diff) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_153_SimpleSlice_DistinctDiff_RightNil_FromS11b(t *testing.T) {
	safeTest(t, "Test_153_SimpleSlice_DistinctDiff_RightNil", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"x"})
		diff := a.DistinctDiff(nil)

		// Act
		actual := args.Map{"result": len(diff) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_154_SimpleSlice_Serialize_FromS11b(t *testing.T) {
	safeTest(t, "Test_154_SimpleSlice_Serialize", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		data, err := ss.Serialize()

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid bytes", actual)
	})
}

func Test_SimpleSlice_155_SimpleSlice_Deserialize_FromS11b(t *testing.T) {
	safeTest(t, "Test_155_SimpleSlice_Deserialize", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		var target []string
		err := ss.Deserialize(&target)

		// Act
		actual := args.Map{"result": err != nil || len(target) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_156_SimpleSlice_SafeStrings_FromS11b(t *testing.T) {
	safeTest(t, "Test_156_SimpleSlice_SafeStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": len(ss.SafeStrings()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(corestr.Empty.SimpleSlice().SafeStrings()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}
