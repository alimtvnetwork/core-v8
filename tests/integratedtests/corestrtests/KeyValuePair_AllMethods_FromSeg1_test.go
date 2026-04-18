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

