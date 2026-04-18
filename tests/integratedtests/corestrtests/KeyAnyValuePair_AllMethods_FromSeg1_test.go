package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args")

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


// --- Appended from Seg6 (Batch 2.4) ---

func Test_KeyAnyValuePair_Basic_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Basic", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "name", Value: "hello"}

		// Act
		actual := args.Map{
			"key":     kv.KeyName(),
			"varName": kv.VariableName(),
			"valAny":  kv.ValueAny() != nil,
			"isEqual": kv.IsVariableNameEqual("name"),
			"notEq":   kv.IsVariableNameEqual("other"),
		}

		// Assert
		expected := args.Map{
			"key": "name",
			"varName": "name",
			"valAny": true,
			"isEqual": true,
			"notEq": false,
		}
		expected.ShouldBeEqual(t, 0, "Basic accessors -- correct", actual)
	})
}

func Test_KeyAnyValuePair_ValueChecks_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_ValueChecks", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "hello"}
		kvNil := corestr.KeyAnyValuePair{Key: "k", Value: nil}

		// Act
		actual := args.Map{
			"hasVal":        kv.HasValue(),
			"hasNonNull":    kv.HasNonNull(),
			"isNull":        kv.IsValueNull(),
			"nilIsNull":     kvNil.IsValueNull(),
			"nilHasVal":     kvNil.HasValue(),
			"emptyStr":      kv.IsValueEmptyString(),
			"whitespace":    kv.IsValueWhitespace(),
			"nilEmptyStr":   kvNil.IsValueEmptyString(),
			"nilWhitespace": kvNil.IsValueWhitespace(),
		}

		// Assert
		expected := args.Map{
			"hasVal": true, "hasNonNull": true, "isNull": false,
			"nilIsNull": true, "nilHasVal": false,
			"emptyStr": false, "whitespace": false,
			"nilEmptyStr": true, "nilWhitespace": true,
		}
		expected.ShouldBeEqual(t, 0, "Value checks -- correct", actual)
	})
}

func Test_KeyAnyValuePair_ValueString_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_ValueString", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{"val": kv.ValueString()}

		// Assert
		expected := args.Map{"val": "42"}
		expected.ShouldBeEqual(t, 0, "ValueString -- formatted", actual)
	})
}

func Test_KeyAnyValuePair_ValueString_Cached_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_ValueString_Cached", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		_ = kv.ValueString() // init

		// Act
		actual := args.Map{"val": kv.ValueString()}

		// Assert
		expected := args.Map{"val": "42"}
		expected.ShouldBeEqual(t, 0, "ValueString cached -- same value", actual)
	})
}

func Test_KeyAnyValuePair_String_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_String", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"nonEmpty": kv.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_KeyAnyValuePair_Compile_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Compile", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"nonEmpty": kv.Compile() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Compile -- delegates to String", actual)
	})
}

func Test_KeyAnyValuePair_SerializeMust_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_SerializeMust", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b := kv.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "SerializeMust -- success", actual)
	})
}

func Test_KeyAnyValuePair_Serialize_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Serialize", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		b, err := kv.Serialize()

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

func Test_KeyAnyValuePair_Json_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Json", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		j := kv.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_KeyAnyValuePair_InterfaceCasts_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_InterfaceCasts", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"jsoner":   kv.AsJsoner() != nil,
			"binder":   kv.AsJsonContractsBinder() != nil,
			"injector": kv.AsJsonParseSelfInjector() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner": true,
			"binder": true,
			"injector": true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_KeyAnyValuePair_ParseInjectUsingJson_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_ParseInjectUsingJson", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		_, err := kv2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_KeyAnyValuePair_ParseInjectUsingJsonMust_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_ParseInjectUsingJsonMust", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		result := kv2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_KeyAnyValuePair_JsonParseSelfInject_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_JsonParseSelfInject", func() {
		// Arrange
		kv := corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		jr := kv.JsonPtr()
		kv2 := &corestr.KeyAnyValuePair{}
		err := kv2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

func Test_KeyAnyValuePair_Clear_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Clear", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kv.Clear()

		// Act
		actual := args.Map{
			"key": kv.Key,
			"isNull": kv.IsValueNull(),
		}

		// Assert
		expected := args.Map{
			"key": "",
			"isNull": true,
		}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_KeyAnyValuePair_Dispose_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Dispose", func() {
		// Arrange
		kv := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kv.Dispose()

		// Act
		actual := args.Map{"key": kv.Key}

		// Assert
		expected := args.Map{"key": ""}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

func Test_KeyAnyValuePair_Clear_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Clear_Nil", func() {
		var kv *corestr.KeyAnyValuePair
		kv.Clear() // should not panic
	})
}

func Test_KeyAnyValuePair_Dispose_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KAVP_Dispose_Nil", func() {
		var kv *corestr.KeyAnyValuePair
		kv.Dispose() // should not panic
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair — Segment 6f
// ══════════════════════════════════════════════════════════════════════════════

