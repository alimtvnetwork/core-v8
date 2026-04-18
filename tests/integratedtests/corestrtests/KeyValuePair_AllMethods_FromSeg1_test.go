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

func Test_KVP_BasicAccessors_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_BasicAccessors_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "name", Value: "alice"}

		// Act
		actual := args.Map{
			"key":      kvp.KeyName(),
			"varName":  kvp.VariableName(),
			"valStr":   kvp.ValueString(),
			"isVarEq":  kvp.IsVariableNameEqual("name"),
			"isValEq":  kvp.IsValueEqual("alice"),
			"compile":  kvp.Compile(),
			"hasKey":   kvp.HasKey(),
			"hasValue": kvp.HasValue(),
		}

		// Assert
		expected := args.Map{
			"key":      "name",
			"varName":  "name",
			"valStr":   "alice",
			"isVarEq":  true,
			"isValEq":  true,
			"compile":  kvp.String(),
			"hasKey":   true,
			"hasValue": true,
		}
		expected.ShouldBeEqual(t, 0, "KVP basic accessors -- happy path", actual)
	})
}

func Test_KVP_EmptyChecks_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_EmptyChecks_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{}

		// Act
		actual := args.Map{
			"keyEmpty":    kvp.IsKeyEmpty(),
			"valEmpty":    kvp.IsValueEmpty(),
			"kvEmpty":     kvp.IsKeyValueEmpty(),
			"kvAnyEmpty":  kvp.IsKeyValueAnyEmpty(),
		}

		// Assert
		expected := args.Map{
			"keyEmpty":    true,
			"valEmpty":    true,
			"kvEmpty":     true,
			"kvAnyEmpty":  true,
		}
		expected.ShouldBeEqual(t, 0, "KVP empty checks -- all empty", actual)
	})
}

func Test_KVP_TrimAndConversions_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_TrimAndConversions_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: " name ", Value: " 42 "}

		// Act
		actual := args.Map{
			"trimKey": kvp.TrimKey(),
			"trimVal": kvp.TrimValue(),
		}

		// Assert
		expected := args.Map{
			"trimKey": "name",
			"trimVal": "42",
		}
		expected.ShouldBeEqual(t, 0, "KVP Trim -- whitespace removed", actual)
	})
}

func Test_KVP_ValueBool_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_ValueBool_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "flag", Value: "true"}
		kvpEmpty := corestr.KeyValuePair{Key: "flag", Value: ""}
		kvpBad := corestr.KeyValuePair{Key: "flag", Value: "notabool"}

		// Act
		actual := args.Map{
			"t": kvp.ValueBool(),
			"empty": kvpEmpty.ValueBool(),
			"bad": kvpBad.ValueBool(),
		}

		// Assert
		expected := args.Map{
			"t": true,
			"empty": false,
			"bad": false,
		}
		expected.ShouldBeEqual(t, 0, "KVP ValueBool -- various inputs", actual)
	})
}

func Test_KVP_ValueInt_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_ValueInt_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "n", Value: "42"}
		kvpBad := corestr.KeyValuePair{Key: "n", Value: "abc"}

		// Act
		actual := args.Map{
			"val": kvp.ValueInt(0),
			"def": kvp.ValueDefInt(),
			"bad": kvpBad.ValueInt(99),
		}

		// Assert
		expected := args.Map{
			"val": 42,
			"def": 42,
			"bad": 99,
		}
		expected.ShouldBeEqual(t, 0, "KVP ValueInt -- valid and invalid", actual)
	})
}

func Test_KVP_ValueByte_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_ValueByte_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "b", Value: "200"}
		kvpBad := corestr.KeyValuePair{Key: "b", Value: "abc"}
		kvpOverflow := corestr.KeyValuePair{Key: "b", Value: "999"}

		// Act
		actual := args.Map{
			"val":      int(kvp.ValueByte(0)),
			"def":      int(kvp.ValueDefByte()),
			"bad":      int(kvpBad.ValueByte(5)),
			"overflow": int(kvpOverflow.ValueByte(7)),
		}

		// Assert
		expected := args.Map{
			"val":      200,
			"def":      200,
			"bad":      5,
			"overflow": 7,
		}
		expected.ShouldBeEqual(t, 0, "KVP ValueByte -- valid, invalid, overflow", actual)
	})
}

func Test_KVP_ValueFloat64_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_ValueFloat64_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "f", Value: "3.14"}
		kvpBad := corestr.KeyValuePair{Key: "f", Value: "abc"}

		// Act
		actual := args.Map{
			"val": kvp.ValueFloat64(0),
			"def": kvp.ValueDefFloat64(),
			"bad": kvpBad.ValueFloat64(1.5),
		}

		// Assert
		expected := args.Map{
			"val": 3.14,
			"def": 3.14,
			"bad": 1.5,
		}
		expected.ShouldBeEqual(t, 0, "KVP ValueFloat64 -- valid and invalid", actual)
	})
}

func Test_KVP_ValueValid_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_ValueValid_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kvp.ValueValid()

		// Act
		actual := args.Map{
			"valid": vv.IsValid,
			"value": vv.Value,
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"value": "v",
		}
		expected.ShouldBeEqual(t, 0, "KVP ValueValid -- default valid", actual)
	})
}

func Test_KVP_ValueValidOptions_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_ValueValidOptions_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kvp.ValueValidOptions(false, "err")

		// Act
		actual := args.Map{
			"valid": vv.IsValid,
			"msg": vv.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "err",
		}
		expected.ShouldBeEqual(t, 0, "KVP ValueValidOptions -- invalid with message", actual)
	})
}

func Test_KVP_Is_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_Is_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"is":    kvp.Is("k", "v"),
			"isKey": kvp.IsKey("k"),
			"isVal": kvp.IsVal("v"),
			"notIs": kvp.Is("x", "y"),
		}

		// Assert
		expected := args.Map{
			"is":    true,
			"isKey": true,
			"isVal": true,
			"notIs": false,
		}
		expected.ShouldBeEqual(t, 0, "KVP Is/IsKey/IsVal -- match and no match", actual)
	})
}

func Test_KVP_NilChecks_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_NilChecks_FromSeg1", func() {
		// Arrange
		var kvp *corestr.KeyValuePair

		// Act
		actual := args.Map{"anyEmpty": kvp.IsKeyValueAnyEmpty()}

		// Assert
		expected := args.Map{"anyEmpty": true}
		expected.ShouldBeEqual(t, 0, "KVP nil IsKeyValueAnyEmpty -- nil receiver", actual)
	})
}

func Test_KVP_FormatString_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_FormatString_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"fmt": kvp.FormatString("%v=%v")}

		// Assert
		expected := args.Map{"fmt": "k=v"}
		expected.ShouldBeEqual(t, 0, "KVP FormatString -- custom format", actual)
	})
}

func Test_KVP_Json_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_Json_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		j := kvp.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KVP Json -- no error", actual)
	})
}

func Test_KVP_Serialize_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_Serialize_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		b, err := kvp.Serialize()

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
		expected.ShouldBeEqual(t, 0, "KVP Serialize -- success", actual)
	})
}

func Test_KVP_SerializeMust_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_SerializeMust_FromSeg1", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		b := kvp.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KVP SerializeMust -- success", actual)
	})
}

func Test_KVP_ClearDispose_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_ClearDispose_FromSeg1", func() {
		// Arrange
		kvp := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kvp.Clear()

		// Act
		actual := args.Map{
			"keyEmpty": kvp.Key == "",
			"valEmpty": kvp.Value == "",
		}

		// Assert
		expected := args.Map{
			"keyEmpty": true,
			"valEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "KVP Clear -- fields emptied", actual)
	})
}

func Test_KVP_DisposeNil_FromSeg1(t *testing.T) {
	safeTest(t, "Test_KVP_DisposeNil_FromSeg1", func() {
		var kvp *corestr.KeyValuePair
		kvp.Dispose() // should not panic
		kvp.Clear()   // should not panic
	})
}


// --- Appended from Seg6 (Batch 2.4) ---

func Test_KeyValuePair_Basic_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Basic", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "name", Value: "hello"}

		// Act
		actual := args.Map{
			"key":      kv.KeyName(),
			"varName":  kv.VariableName(),
			"valStr":   kv.ValueString(),
			"isEqual":  kv.IsVariableNameEqual("name"),
			"isValEq":  kv.IsValueEqual("hello"),
			"isValNeq": kv.IsValueEqual("other"),
		}

		// Assert
		expected := args.Map{
			"key": "name", "varName": "name", "valStr": "hello",
			"isEqual": true, "isValEq": true, "isValNeq": false,
		}
		expected.ShouldBeEqual(t, 0, "Basic accessors -- correct", actual)
	})
}

func Test_KeyValuePair_Empty_Checks_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Empty_Checks", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kvEmpty := &corestr.KeyValuePair{}

		// Act
		actual := args.Map{
			"keyEmpty":     kv.IsKeyEmpty(),
			"valEmpty":     kv.IsValueEmpty(),
			"hasKey":       kv.HasKey(),
			"hasVal":       kv.HasValue(),
			"kvEmpty":      kv.IsKeyValueEmpty(),
			"emptyKeyE":    kvEmpty.IsKeyEmpty(),
			"emptyValE":    kvEmpty.IsValueEmpty(),
			"emptyKVE":     kvEmpty.IsKeyValueEmpty(),
			"anyEmpty":     kv.IsKeyValueAnyEmpty(),
			"anyEmptyFull": kvEmpty.IsKeyValueAnyEmpty(),
		}

		// Assert
		expected := args.Map{
			"keyEmpty": false, "valEmpty": false,
			"hasKey": true, "hasVal": true,
			"kvEmpty": false, "emptyKeyE": true,
			"emptyValE": true, "emptyKVE": true,
			"anyEmpty": false, "anyEmptyFull": true,
		}
		expected.ShouldBeEqual(t, 0, "Empty checks -- correct", actual)
	})
}

func Test_KeyValuePair_Is_IsKey_IsVal_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Is_IsKey_IsVal", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"is":    kv.Is("k", "v"),
			"isNot": kv.Is("k", "x"),
			"isKey": kv.IsKey("k"),
			"isVal": kv.IsVal("v"),
		}

		// Assert
		expected := args.Map{
			"is": true,
			"isNot": false,
			"isKey": true,
			"isVal": true,
		}
		expected.ShouldBeEqual(t, 0, "Is/IsKey/IsVal -- correct", actual)
	})
}

func Test_KeyValuePair_Trim_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Trim", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: " k ", Value: " v "}

		// Act
		actual := args.Map{
			"trimKey": kv.TrimKey(),
			"trimVal": kv.TrimValue(),
		}

		// Assert
		expected := args.Map{
			"trimKey": "k",
			"trimVal": "v",
		}
		expected.ShouldBeEqual(t, 0, "Trim -- trimmed", actual)
	})
}

func Test_KeyValuePair_ValueBool_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueBool", func() {
		// Arrange
		kvTrue := &corestr.KeyValuePair{Key: "k", Value: "true"}
		kvFalse := &corestr.KeyValuePair{Key: "k", Value: "invalid"}
		kvEmpty := &corestr.KeyValuePair{Key: "k", Value: ""}

		// Act
		actual := args.Map{
			"true": kvTrue.ValueBool(),
			"invalid": kvFalse.ValueBool(),
			"empty": kvEmpty.ValueBool(),
		}

		// Assert
		expected := args.Map{
			"true": true,
			"invalid": false,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "ValueBool -- various", actual)
	})
}

func Test_KeyValuePair_ValueInt_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueInt", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "42"}
		kvBad := &corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act
		actual := args.Map{
			"val": kv.ValueInt(0),
			"def": kvBad.ValueInt(99),
			"defInt": kv.ValueDefInt(),
			"badDef": kvBad.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"val": 42,
			"def": 99,
			"defInt": 42,
			"badDef": 0,
		}
		expected.ShouldBeEqual(t, 0, "ValueInt -- various", actual)
	})
}

func Test_KeyValuePair_ValueByte_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueByte", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "65"}
		kvBad := &corestr.KeyValuePair{Key: "k", Value: "abc"}
		kvBig := &corestr.KeyValuePair{Key: "k", Value: "999"}

		// Act
		actual := args.Map{
			"val":    kv.ValueByte(0),
			"bad":    kvBad.ValueByte(1),
			"big":    kvBig.ValueByte(2),
			"defVal": kv.ValueDefByte(),
			"defBad": kvBad.ValueDefByte(),
			"defBig": kvBig.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"val": byte(65),
			"bad": byte(1),
			"big": byte(2),
			"defVal": byte(65),
			"defBad": byte(0),
			"defBig": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "ValueByte -- various", actual)
	})
}

func Test_KeyValuePair_ValueFloat64_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueFloat64", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "1.5"}
		kvBad := &corestr.KeyValuePair{Key: "k", Value: "abc"}

		// Act
		actual := args.Map{
			"val": kv.ValueFloat64(0),
			"bad": kvBad.ValueFloat64(9.9),
			"def": kv.ValueDefFloat64(),
		}

		// Assert
		expected := args.Map{
			"val": 1.5,
			"bad": 9.9,
			"def": 1.5,
		}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 -- various", actual)
	})
}

func Test_KeyValuePair_ValueValid_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueValid", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "hello"}
		vv := kv.ValueValid()

		// Act
		actual := args.Map{
			"val": vv.Value,
			"valid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueValid -- valid", actual)
	})
}

func Test_KeyValuePair_ValueValidOptions_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_ValueValidOptions", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "hello"}
		vv := kv.ValueValidOptions(false, "err")

		// Act
		actual := args.Map{
			"val": vv.Value,
			"valid": vv.IsValid,
			"msg": vv.Message,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": false,
			"msg": "err",
		}
		expected.ShouldBeEqual(t, 0, "ValueValidOptions -- custom", actual)
	})
}

func Test_KeyValuePair_FormatString_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_FormatString", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"val": kv.FormatString("%v=%v")}

		// Assert
		expected := args.Map{"val": "k=v"}
		expected.ShouldBeEqual(t, 0, "FormatString -- formatted", actual)
	})
}

func Test_KeyValuePair_String_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_String", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"nonEmpty": kv.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_KeyValuePair_Compile_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Compile", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"nonEmpty": kv.Compile() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Compile -- delegates to String", actual)
	})
}

func Test_KeyValuePair_Json_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Json", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		j := kv.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_KeyValuePair_Serialize_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Serialize", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
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

func Test_KeyValuePair_SerializeMust_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_SerializeMust", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		b := kv.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "SerializeMust -- success", actual)
	})
}

func Test_KeyValuePair_Clear_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Clear", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()

		// Act
		actual := args.Map{
			"key": kv.Key,
			"val": kv.Value,
		}

		// Assert
		expected := args.Map{
			"key": "",
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_KeyValuePair_Dispose_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Dispose", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Dispose()

		// Act
		actual := args.Map{"key": kv.Key}

		// Assert
		expected := args.Map{"key": ""}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

func Test_KeyValuePair_Clear_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Clear_Nil", func() {
		var kv *corestr.KeyValuePair
		kv.Clear() // should not panic
	})
}

func Test_KeyValuePair_Dispose_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVP_Dispose_Nil", func() {
		var kv *corestr.KeyValuePair
		kv.Dispose() // should not panic
	})
}
