package coreoncetests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage Gaps Iteration 24
//
// Targets remaining reachable branches in coredata/coreonce (99.7% → 100%)
// ══════════════════════════════════════════════════════════════════════════════

// ── AnyOnce: Deserialize with nil toPtr ──

func Test_AnyOnce_Deserialize_NilToPtr(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return "hello" })
	err := ao.Deserialize(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyOnce Deserialize returns error -- nil toPtr", actual)
}

func Test_AnyOnce_Deserialize_EmptyBytes(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return nil })
	var result string
	err := ao.Deserialize(&result)

	// Act — json.Marshal(nil) produces "null", Unmarshal("null", &string) succeeds
	actual := args.Map{
		"hasErr": err != nil,
		"result": result,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"result": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce Deserialize succeeds -- nil value marshals to null", actual)
}

// ── AnyOnce: SerializeSkipExistingError ──

func Test_AnyOnce_SerializeSkipExistingError_ReturnsEarly(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return "test" })
	bytes, err := ao.SerializeSkipExistingError()

	// Act
	actual := args.Map{
		"hasErr":   err != nil,
		"hasBytes": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"hasErr":   false,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce SerializeSkipExistingError works -- valid value", actual)
}

// ── AnyOnce: IsStringEmpty / IsStringEmptyOrWhitespace ──

func Test_AnyOnce_IsStringEmpty_NonEmpty(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return "hello" })

	// Act
	actual := args.Map{
		"isEmpty":    ao.IsStringEmpty(),
		"isWhitespace": ao.IsStringEmptyOrWhitespace(),
	}

	// Assert
	expected := args.Map{
		"isEmpty":    false,
		"isWhitespace": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce IsStringEmpty returns false -- non-empty value", actual)
}

// ── AnyOnce: IsInitialized before and after ──

func Test_AnyOnce_IsInitialized_AfterSet(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return "hello" })
	beforeInit := ao.IsInitialized()
	_ = ao.Value()
	afterInit := ao.IsInitialized()

	// Act
	actual := args.Map{
		"before": beforeInit,
		"after":  afterInit,
	}

	// Assert
	expected := args.Map{
		"before": false,
		"after":  true,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce IsInitialized returns correct state -- before/after", actual)
}

// ── AnyOnce: ValueOnly alias ──

func Test_AnyOnce_ValueOnly_IgnoresError(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return 42 })
	val := ao.ValueOnly()

	// Act
	actual := args.Map{"isNil": val == nil}

	// Assert
	expected := args.Map{"isNil": false}
	expected.ShouldBeEqual(t, 0, "AnyOnce ValueOnly returns value -- non-nil", actual)
}

// ── AnyOnce: CastValueHashmapMap ──

func Test_AnyOnce_CastValueHashmapMap_Valid(t *testing.T) {
	// Arrange
	m := map[string]string{"a": "1"}
	ao := coreonce.NewAnyOncePtr(func() any { return m })
	result, ok := ao.CastValueHashmapMap()

	// Act
	actual := args.Map{
		"ok":  ok,
		"val": result["a"],
	}

	// Assert
	expected := args.Map{
		"ok":  true,
		"val": "1",
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce CastValueHashmapMap returns map -- valid", actual)
}

// ── AnyOnce: CastValueMapStringAnyMap ──

func Test_AnyOnce_CastValueMapStringAnyMap_Valid(t *testing.T) {
	// Arrange
	m := map[string]any{"a": 1}
	ao := coreonce.NewAnyOncePtr(func() any { return m })
	result, ok := ao.CastValueMapStringAnyMap()

	// Act
	actual := args.Map{
		"ok":    ok,
		"hasKey": result["a"] != nil,
	}

	// Assert
	expected := args.Map{
		"ok":    true,
		"hasKey": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce CastValueMapStringAnyMap returns map -- valid", actual)
}

// ── AnyOnce: CastValueBytes ──

func Test_AnyOnce_CastValueBytes_Valid(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return []byte("hello") })
	result, ok := ao.CastValueBytes()

	// Act
	actual := args.Map{
		"ok":  ok,
		"val": string(result),
	}

	// Assert
	expected := args.Map{
		"ok":  true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce CastValueBytes returns bytes -- valid", actual)
}

// ── AnyOnce: CastValueStrings ──

func Test_AnyOnce_CastValueStrings_Valid(t *testing.T) {
	// Arrange
	ao := coreonce.NewAnyOncePtr(func() any { return []string{"a", "b"} })
	result, ok := ao.CastValueStrings()

	// Act
	actual := args.Map{
		"ok":  ok,
		"len": len(result),
	}

	// Assert
	expected := args.Map{
		"ok":  true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "AnyOnce CastValueStrings returns strings -- valid", actual)
}

// ── AnyErrorOnce: IsStringEmpty / IsStringEmptyOrWhitespace ──

func Test_AnyErrorOnce_IsStringEmpty_NonEmpty(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "hello", nil })

	// Act
	actual := args.Map{
		"isEmpty":      aeo.IsStringEmpty(),
		"isWhitespace": aeo.IsStringEmptyOrWhitespace(),
	}

	// Assert
	expected := args.Map{
		"isEmpty":      false,
		"isWhitespace": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce IsStringEmpty returns false -- non-empty", actual)
}

// ── AnyErrorOnce: CastValueHashmapMap ──

func Test_AnyErrorOnce_CastValueHashmapMap_Valid(t *testing.T) {
	// Arrange
	m := map[string]string{"k": "v"}
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return m, nil })
	result, err, ok := aeo.CastValueHashmapMap()

	// Act
	actual := args.Map{
		"ok":    ok,
		"noErr": err == nil,
		"val":   result["k"],
	}

	// Assert
	expected := args.Map{
		"ok":    true,
		"noErr": true,
		"val":   "v",
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce CastValueHashmapMap returns map -- valid", actual)
}

// ── AnyErrorOnce: CastValueMapStringAnyMap ──

func Test_AnyErrorOnce_CastValueMapStringAnyMap_Valid(t *testing.T) {
	// Arrange
	m := map[string]any{"k": 1}
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return m, nil })
	result, err, ok := aeo.CastValueMapStringAnyMap()

	// Act
	actual := args.Map{
		"ok":    ok,
		"noErr": err == nil,
		"hasK":  result["k"] != nil,
	}

	// Assert
	expected := args.Map{
		"ok":    true,
		"noErr": true,
		"hasK":  true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce CastValueMapStringAnyMap returns map -- valid", actual)
}

// ── AnyErrorOnce: CastValueBytes ──

func Test_AnyErrorOnce_CastValueBytes_Valid(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return []byte("data"), nil })
	result, err, ok := aeo.CastValueBytes()

	// Act
	actual := args.Map{
		"ok":    ok,
		"noErr": err == nil,
		"val":   string(result),
	}

	// Assert
	expected := args.Map{
		"ok":    true,
		"noErr": true,
		"val":   "data",
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce CastValueBytes returns bytes -- valid", actual)
}

// ── AnyErrorOnce: CastValueStrings ──

func Test_AnyErrorOnce_CastValueStrings_Valid(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return []string{"x"}, nil })
	result, err, ok := aeo.CastValueStrings()

	// Act
	actual := args.Map{
		"ok":    ok,
		"noErr": err == nil,
		"len":   len(result),
	}

	// Assert
	expected := args.Map{
		"ok":    true,
		"noErr": true,
		"len":   1,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce CastValueStrings returns strings -- valid", actual)
}

// ── AnyErrorOnce: SerializeSkipExistingError ──

func Test_AnyErrorOnce_SerializeSkipExistingError_ReturnsEarly(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "test", nil })
	bytes, err := aeo.SerializeSkipExistingError()

	// Act
	actual := args.Map{
		"hasErr":   err != nil,
		"hasBytes": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"hasErr":   false,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce SerializeSkipExistingError works -- valid", actual)
}

// ── AnyErrorOnce: Serialize with existing error ──

func Test_AnyErrorOnce_Serialize_ExistingError_Skips(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, errors.New("fail") })
	_, err := aeo.Serialize()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce Serialize returns error -- existing error", actual)
}

// ── AnyErrorOnce: ValueString cached ──

func Test_AnyErrorOnce_ValueString_Cached_ReturnsSame(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return "hello", nil })
	v1, _ := aeo.ValueString()
	v2, _ := aeo.ValueString()

	// Act
	actual := args.Map{
		"v1": v1,
		"v2": v2,
	}

	// Assert
	expected := args.Map{
		"v1": "hello",
		"v2": "hello",
	}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ValueString returns cached -- second call", actual)
}

// ── AnyErrorOnce: ValueString nil value ──

func Test_AnyErrorOnce_ValueString_NilValue_ReturnsEmpty(t *testing.T) {
	// Arrange
	aeo := coreonce.NewAnyErrorOncePtr(func() (any, error) { return nil, nil })
	v, _ := aeo.ValueString()

	// Act
	actual := args.Map{"notEmpty": v != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyErrorOnce ValueString returns nil bracket -- nil value", actual)
}

// ── BytesErrorOnce: HasIssuesOrEmpty / HasSafeItems ──

func Test_BytesErrorOnce_HasIssuesOrEmpty_NoData(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })

	// Act
	actual := args.Map{
		"hasIssues": beo.HasIssuesOrEmpty(),
		"hasSafe":   beo.HasSafeItems(),
	}

	// Assert
	expected := args.Map{
		"hasIssues": true,
		"hasSafe":   false,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce HasIssuesOrEmpty true -- no data", actual)
}

func Test_BytesErrorOnce_HasIssuesOrEmpty_WithError_ReturnsTrue(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("data"), errors.New("err") })

	// Act
	actual := args.Map{
		"hasIssues": beo.HasIssuesOrEmpty(),
		"hasSafe":   beo.HasSafeItems(),
	}

	// Assert
	expected := args.Map{
		"hasIssues": true,
		"hasSafe":   false,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce HasIssuesOrEmpty true -- has error", actual)
}

func Test_BytesErrorOnce_HasSafeItems_Valid(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("data"), nil })

	// Act
	actual := args.Map{
		"hasIssues": beo.HasIssuesOrEmpty(),
		"hasSafe":   beo.HasSafeItems(),
	}

	// Assert
	expected := args.Map{
		"hasIssues": false,
		"hasSafe":   true,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce HasSafeItems true -- valid data", actual)
}

// ── BytesErrorOnce: IsBytesEmpty / IsEmptyBytes ──

func Test_BytesErrorOnce_IsBytesEmpty(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, nil })

	// Act
	actual := args.Map{
		"isBytesEmpty": beo.IsBytesEmpty(),
		"isEmptyBytes": beo.IsEmptyBytes(),
	}

	// Assert
	expected := args.Map{
		"isBytesEmpty": true,
		"isEmptyBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce IsBytesEmpty true -- nil bytes", actual)
}

// ── BytesErrorOnce: IsStringEmpty / IsStringEmptyOrWhitespace ──

func Test_BytesErrorOnce_IsStringEmpty_NonEmpty(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("hello"), nil })

	// Act
	actual := args.Map{
		"isEmpty":      beo.IsStringEmpty(),
		"isWhitespace": beo.IsStringEmptyOrWhitespace(),
	}

	// Assert
	expected := args.Map{
		"isEmpty":      false,
		"isWhitespace": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce IsStringEmpty false -- has data", actual)
}

// ── BytesErrorOnce: MarshalJSON ──

func Test_BytesErrorOnce_MarshalJSON_Valid(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte(`"hello"`), nil })
	bytes, err := beo.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce MarshalJSON works -- valid data", actual)
}

// ── BytesErrorOnce: ValueWithError alias ──

func Test_BytesErrorOnce_ValueWithError_ReturnsBoth(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("ok"), nil })
	bytes, err := beo.ValueWithError()

	// Act
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce ValueWithError works -- valid", actual)
}

// ── BytesErrorOnce: Deserialize with valid toPtr but existing error ──

func Test_BytesErrorOnce_Deserialize_ExistingError_Skips(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return nil, errors.New("fail") })
	var result string
	err := beo.Deserialize(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce Deserialize returns error -- existing error", actual)
}

// ── BytesErrorOnce: Deserialize with invalid JSON ──

func Test_BytesErrorOnce_Deserialize_InvalidJson(t *testing.T) {
	// Arrange
	beo := coreonce.NewBytesErrorOncePtr(func() ([]byte, error) { return []byte("not-json"), nil })
	var result string
	err := beo.Deserialize(&result)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesErrorOnce Deserialize returns error -- invalid json", actual)
}

// ── ErrorOnce: Serialize with nil error ──

func Test_ErrorOnce_Serialize_NilError(t *testing.T) {
	// Arrange
	eo := coreonce.NewErrorOncePtr(func() error { return nil })
	bytes, err := eo.Serialize()

	// Act
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "ErrorOnce Serialize works -- nil error", actual)
}

// ── MapStringStringOnce: GetValue / GetValueWithStatus ──

func Test_MapStringStringOnce_GetValue_Found(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})
	v := mo.GetValue("a")
	missing := mo.GetValue("z")

	// Act
	actual := args.Map{
		"found":   v,
		"missing": missing,
	}

	// Assert
	expected := args.Map{
		"found":   "1",
		"missing": "",
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce GetValue returns value -- found and missing", actual)
}

func Test_MapStringStringOnce_GetValueWithStatus_Found(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	v, hasA := mo.GetValueWithStatus("a")
	_, hasZ := mo.GetValueWithStatus("z")

	// Act
	actual := args.Map{
		"val":  v,
		"hasA": hasA,
		"hasZ": hasZ,
	}

	// Assert
	expected := args.Map{
		"val":  "1",
		"hasA": true,
		"hasZ": false,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce GetValueWithStatus returns correct -- found and missing", actual)
}

// ── MapStringStringOnce: IsMissing ──

func Test_MapStringStringOnce_IsMissing_ReturnsTrue(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})

	// Act
	actual := args.Map{
		"missingZ": mo.IsMissing("z"),
		"missingA": mo.IsMissing("a"),
	}

	// Assert
	expected := args.Map{
		"missingZ": true,
		"missingA": false,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce IsMissing returns correct -- has and missing", actual)
}

// ── MapStringStringOnce: Length with nil map ──

func Test_MapStringStringOnce_Length_NilMap(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return nil
	})

	// Act
	actual := args.Map{"len": mo.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce Length returns 0 -- nil map", actual)
}

// ── MapStringStringOnce: AllKeys / AllValues empty ──

func Test_MapStringStringOnce_AllKeys_Empty_ReturnsNil(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{}
	})

	// Act
	actual := args.Map{
		"keysLen":   len(mo.AllKeys()),
		"valuesLen": len(mo.AllValues()),
	}

	// Assert
	expected := args.Map{
		"keysLen":   0,
		"valuesLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce AllKeys/AllValues return empty -- empty map", actual)
}

// ── MapStringStringOnce: AllKeys / AllValues cached ──

func Test_MapStringStringOnce_AllKeys_Cached_ReturnsSame(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})
	_ = mo.AllKeys()
	k := mo.AllKeys()
	_ = mo.AllValues()
	v := mo.AllValues()

	// Act
	actual := args.Map{
		"keysLen":   len(k),
		"valuesLen": len(v),
	}

	// Assert
	expected := args.Map{
		"keysLen":   2,
		"valuesLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce AllKeys/AllValues cached -- second call", actual)
}

// ── MapStringStringOnce: Strings empty map ──

func Test_MapStringStringOnce_Strings_EmptyMap(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{}
	})

	// Act
	actual := args.Map{"len": len(mo.Strings())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce Strings returns empty -- empty map", actual)
}

// ── MapStringStringOnce: Strings cached ──

func Test_MapStringStringOnce_Strings_Cached_ReturnsSame(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})
	_ = mo.Strings()
	s := mo.Strings()

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce Strings cached -- second call", actual)
}

// ── MapStringStringOnce: HasAll ──

func Test_MapStringStringOnce_HasAll_AllPresent(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1", "b": "2"}
	})

	// Act
	actual := args.Map{
		"hasAll":     mo.HasAll("a", "b"),
		"hasPartial": mo.HasAll("a", "z"),
	}

	// Assert
	expected := args.Map{
		"hasAll":     true,
		"hasPartial": false,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce HasAll returns correct -- all and partial", actual)
}

// ── MapStringStringOnce: IsEqual nil cases ──

func Test_MapStringStringOnce_IsEqual_BothNil_ReturnsTrue(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return nil
	})

	// Act
	actual := args.Map{
		"equalNil":    mo.IsEqual(nil),
		"equalNotNil": mo.IsEqual(map[string]string{"a": "1"}),
	}

	// Assert
	expected := args.Map{
		"equalNil":    true,
		"equalNotNil": false,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce IsEqual handles nil -- nil map", actual)
}

func Test_MapStringStringOnce_IsEqual_DiffLen_ReturnsFalse(t *testing.T) {
	// Arrange
	mo := coreonce.NewMapStringStringOncePtr(func() map[string]string {
		return map[string]string{"a": "1"}
	})

	// Act
	actual := args.Map{
		"result": mo.IsEqual(map[string]string{"a": "1", "b": "2"}),
	}

	// Assert
	expected := args.Map{
		"result": false,
	}
	expected.ShouldBeEqual(t, 0, "MapStringStringOnce IsEqual returns false -- diff len", actual)
}

// ── IntegersOnce: Sorted cached ──

func Test_IntegersOnce_Sorted_Cached_ReturnsSame(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{3, 1, 2} })
	_ = io.Sorted()
	s := io.Sorted()

	// Act
	actual := args.Map{
		"first": s[0],
		"last":  s[len(s)-1],
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"last":  3,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce Sorted cached -- second call", actual)
}

// ── IntegersOnce: String ──

func Test_IntegersOnce_String(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2} })

	// Act
	actual := args.Map{"notEmpty": io.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "IntegersOnce String returns non-empty -- has items", actual)
}

// ── IntegersOnce: alias methods ──

func Test_IntegersOnce_Aliases_ValuesListExecute(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{1, 2} })

	// Act
	actual := args.Map{
		"valuesLen":  len(io.Values()),
		"executeLen": len(io.Execute()),
		"listLen":    len(io.List()),
	}

	// Assert
	expected := args.Map{
		"valuesLen":  2,
		"executeLen": 2,
		"listLen":    2,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce alias methods return correct len -- 2 items", actual)
}

// ── IntegersOnce: RangesMap / RangesBoolMap / UniqueMap with data ──

func Test_IntegersOnce_Maps_WithData(t *testing.T) {
	// Arrange
	io := coreonce.NewIntegersOncePtr(func() []int { return []int{10, 20, 30} })

	// Act
	actual := args.Map{
		"rangesLen":     len(io.RangesMap()),
		"rangesBoolLen": len(io.RangesBoolMap()),
		"uniqueLen":     len(io.UniqueMap()),
	}

	// Assert
	expected := args.Map{
		"rangesLen":     3,
		"rangesBoolLen": 3,
		"uniqueLen":     3,
	}
	expected.ShouldBeEqual(t, 0, "IntegersOnce maps return correct len -- 3 items", actual)
}

// ── StringsOnce: Value cached (isInitialized == true) ──

func Test_StringsOnce_Value_Cached(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return []string{"a"} })
	// Trigger initialization via UnmarshalJSON to set isInitialized=true
	err := so.UnmarshalJSON([]byte(`["x","y"]`))
	val := so.Value()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len":   len(val),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len":   2,
	}
	expected.ShouldBeEqual(t, 0, "StringsOnce Value returns unmarshalled data -- after UnmarshalJSON", actual)
}

// ── StringsOnce: Length with nil values ──

func Test_StringsOnce_Length_Nil_ReturnsZero(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return nil })

	// Act
	actual := args.Map{"len": so.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringsOnce Length returns 0 -- nil", actual)
}

// ── StringsOnce: UniqueMap with nil values ──

func Test_StringsOnce_UniqueMap_Nil_ReturnsEmpty(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return nil })

	// Act
	actual := args.Map{"len": len(so.UniqueMap())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringsOnce UniqueMap returns empty -- nil values", actual)
}

// ── StringsOnce: RangesMap empty ──

func Test_StringsOnce_RangesMap_Empty_ReturnsNil(t *testing.T) {
	// Arrange
	so := coreonce.NewStringsOncePtr(func() []string { return []string{} })

	// Act
	actual := args.Map{"len": len(so.RangesMap())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringsOnce RangesMap returns empty -- empty", actual)
}

// ── ByteOnce: UnmarshalJSON ──

func Test_ByteOnce_UnmarshalJSON_Valid(t *testing.T) {
	// Arrange
	bo := coreonce.NewByteOncePtr(func() byte { return 0 })
	err := bo.UnmarshalJSON([]byte("42"))
	v := bo.Value()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val":   v,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val":   byte(42),
	}
	expected.ShouldBeEqual(t, 0, "ByteOnce UnmarshalJSON sets value -- valid byte", actual)
}

// ── BytesOnce: Value with nil initializerFunc ──

func Test_BytesOnce_UnmarshalJSON_Valid(t *testing.T) {
	// Arrange
	bo := coreonce.NewBytesOncePtr(func() []byte { return nil })
	jsonBytes, _ := json.Marshal([]byte("hello"))
	err := bo.UnmarshalJSON(jsonBytes)
	v := bo.Value()

	// Act
	actual := args.Map{
		"noErr":  err == nil,
		"hasVal": len(v) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr":  true,
		"hasVal": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesOnce UnmarshalJSON sets value -- valid", actual)
}

// ── BytesOnce: nil initializerFunc path ──

func Test_BytesOnce_NilInitializer_ReturnsNil(t *testing.T) {
	// Arrange
	bo := &coreonce.BytesOnce{}
	v := bo.Value()

	// Act
	actual := args.Map{
		"isNil": v == nil,
	}

	// Assert
	expected := args.Map{
		"isNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesOnce Value returns nil -- nil initializer", actual)
}
