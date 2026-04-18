package corestrtests

import (
	"errors"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ══════════════════════════════════════════════════════════════════════════════

func Test_KAVP_BasicAccessors_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_BasicAccessors_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "name", Value: 42}

		// Act
		actual := args.Map{
			"key":      kav.KeyName(),
			"varName":  kav.VariableName(),
			"valAny":   kav.ValueAny(),
			"isVarEq":  kav.IsVariableNameEqual("name"),
			"isNull":   kav.IsValueNull(),
			"hasValue": kav.HasValue(),
			"hasNon":   kav.HasNonNull(),
		}

		// Assert
		expected := args.Map{
			"key":      "name",
			"varName":  "name",
			"valAny":   42,
			"isVarEq":  true,
			"isNull":   false,
			"hasValue": true,
			"hasNon":   true,
		}
		expected.ShouldBeEqual(t, 0, "KAVP basic accessors -- happy path", actual)
	})
}

func Test_KAVP_NilValue_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_NilValue_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act
		actual := args.Map{
			"isNull":  kav.IsValueNull(),
			"emptyStr": kav.IsValueEmptyString(),
			"ws":       kav.IsValueWhitespace(),
		}

		// Assert
		expected := args.Map{
			"isNull":  true,
			"emptyStr": true,
			"ws":       true,
		}
		expected.ShouldBeEqual(t, 0, "KAVP nil value -- all empty checks true", actual)
	})
}

func Test_KAVP_ValueString_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_ValueString_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "hello"}

		// Act
		actual := args.Map{"valStr": kav.ValueString()}

		// Assert
		expected := args.Map{"valStr": "hello"}
		expected.ShouldBeEqual(t, 0, "KAVP ValueString -- string value", actual)
	})
}

func Test_KAVP_ValueStringCached_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_ValueStringCached_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: 99}
		// First call initializes, second returns cached
		v1 := kav.ValueString()
		v2 := kav.ValueString()

		// Act
		actual := args.Map{"same": v1 == v2}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "KAVP ValueString cached -- same on second call", actual)
	})
}

func Test_KAVP_Json_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_Json_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		j := kav.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KAVP Json -- no error", actual)
	})
}

func Test_KAVP_Serialize_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_Serialize_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b, err := kav.Serialize()

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
		expected.ShouldBeEqual(t, 0, "KAVP Serialize -- success", actual)
	})
}

func Test_KAVP_SerializeMust_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_SerializeMust_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b := kav.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KAVP SerializeMust -- success", actual)
	})
}

func Test_KAVP_ParseInjectUsingJson_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_ParseInjectUsingJson_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kav.JsonPtr()
		kav2 := &corestr.KeyAnyValuePair{}
		result, err := kav2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"notNil": result != nil,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "KAVP ParseInjectUsingJson -- round trip", actual)
	})
}

func Test_KAVP_ParseInjectUsingJsonMust_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_ParseInjectUsingJsonMust_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kav.JsonPtr()
		kav2 := &corestr.KeyAnyValuePair{}
		result := kav2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "KAVP ParseInjectUsingJsonMust -- no panic", actual)
	})
}

func Test_KAVP_ParseInjectUsingJsonMust_Panic_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_ParseInjectUsingJsonMust_Panic_FromSeg1", func() {
		defer func() { recover() }()
		kav := &corestr.KeyAnyValuePair{}
		badJson := &corejson.Result{}
		_ = kav.ParseInjectUsingJsonMust(badJson)
	})
}

func Test_KAVP_JsonParseSelfInject_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_JsonParseSelfInject_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kav.JsonPtr()
		kav2 := &corestr.KeyAnyValuePair{}
		err := kav2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KAVP JsonParseSelfInject -- success", actual)
	})
}

func Test_KAVP_Interfaces_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_Interfaces_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"binder":   kav.AsJsonContractsBinder() != nil,
			"jsoner":   kav.AsJsoner() != nil,
			"injector": kav.AsJsonParseSelfInjector() != nil,
		}

		// Assert
		expected := args.Map{
			"binder":   true,
			"jsoner":   true,
			"injector": true,
		}
		expected.ShouldBeEqual(t, 0, "KAVP interface casts -- all non-nil", actual)
	})
}

func Test_KAVP_String_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_String_FromSeg1", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"nonEmpty": kav.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "KAVP String -- non-empty", actual)
	})
}

func Test_KAVP_Compile_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_Compile_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"eq": kav.Compile() == kav.String()}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "KAVP Compile equals String -- same result", actual)
	})
}

func Test_KAVP_ClearDispose_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_ClearDispose_FromSeg1", func() {
		// Arrange
		kav := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kav.Clear()

		// Act
		actual := args.Map{
			"keyEmpty": kav.Key == "",
			"valNil": kav.Value == nil,
		}

		// Assert
		expected := args.Map{
			"keyEmpty": true,
			"valNil": true,
		}
		expected.ShouldBeEqual(t, 0, "KAVP Clear -- fields emptied", actual)
	})
}

func Test_KAVP_DisposeNil_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_DisposeNil_FromSeg1", func() {
		var kav *corestr.KeyAnyValuePair
		kav.Dispose() // should not panic
		kav.Clear()   // should not panic
	})
}

func Test_KAVP_NilIsValueNull_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KAVP_NilIsValueNull_FromSeg1", func() {
		// Arrange
		var kav *corestr.KeyAnyValuePair

		// Act
		actual := args.Map{"isNull": kav.IsValueNull()}

		// Assert
		expected := args.Map{"isNull": true}
		expected.ShouldBeEqual(t, 0, "KAVP nil receiver IsValueNull -- true", actual)
	})
}

