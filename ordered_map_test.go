package generic

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestOrderedMapAny(t *testing.T) {
	var (
		om     = NewOrderedMap[string, any]()
		expect = []byte(`{"c":true,"a":{"test":[1,2,3],"other":[0,"string",null]},"b":["Hello","World"]}`)
	)

	om.Set("c", true)

	second := NewOrderedMap[string, any]()
	second.Set("test", []int{1, 2, 3})
	second.Set("other", []any{0, "string", nil})
	om.Set("a", second)

	om.Set("b", []string{"Hello", "World"})

	data, err := json.Marshal(om)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, expect) {
		t.Errorf("expected OrderedMap.Marshal to be %q, got %q", expect, data)
	}
}

func TestOrderedMapUnmarshalMap(t *testing.T) {
	var (
		om *OrderedMap[string, string]

		js = []byte(`{
			"a": "123",
			"b": "456",
			"c": "789"
		}`)

		keys = []string{"a", "b", "c"}
		exp  = map[string]string{"a": "123", "b": "456", "c": "789"}
	)

	if err := json.Unmarshal(js, &om); err != nil {
		t.Error(err)
	}

	if len(om.keys) != len(keys) {
		t.Errorf("expected OrderedMap to have len %d, got %d", len(keys), len(om.keys))
	}

	if !EqualsSlice(om.keys, keys) {
		t.Errorf("expected OrderedMap to have keys %+v, got %+v", keys, om.keys)
	}

	if !EqualsMap(om.umap, exp) {
		t.Errorf("expected OrderedMap to have map %+v, got %+v", exp, om.umap)
	}
}

func TestOrderedMapUnmarshalList(t *testing.T) {
	var (
		om *OrderedMap[string, string]

		js = []byte(`[
			["a", "123"],
			["b", "456"],
			["c", "789"]
		]`)

		keys   = []string{"a", "b", "c"}
		expect = map[string]string{"a": "123", "b": "456", "c": "789"}
	)

	if err := json.Unmarshal(js, &om); err != nil {
		t.Error(err)
	}

	if len(om.keys) != len(keys) {
		t.Errorf("expected OrderedMap to have len %d, got %d", len(keys), len(om.keys))
	}

	if !EqualsSlice(om.keys, keys) {
		t.Errorf("expected OrderedMap to have keys %+v, got %+v", keys, om.keys)
	}

	if !EqualsMap(om.umap, expect) {
		t.Errorf("expected OrderedMap to have map %+v, got %+v", expect, om.umap)
	}
}

func TestOrderedMapMarshalMap(t *testing.T) {
	var (
		om = NewOrderedMap[string, string]()

		expect = []byte(`{"a":"123","b":"456","c":"789"}`)
	)
	om.Set("a", "123")
	om.Set("b", "456")
	om.Set("c", "789")

	data, err := json.Marshal(om)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, expect) {
		t.Errorf("expected OrderedMap.Marshal to be %q, got %q", expect, data)
	}
}

func TestOrderedMapMarshalList(t *testing.T) {
	var (
		om = NewOrderedMap[string, string]().AsList()

		expect = []byte(`[["a","123"],["b","456"],["c","789"]]`)
	)
	om.Set("a", "123")
	om.Set("b", "456")
	om.Set("c", "789")

	data, err := json.Marshal(om)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, expect) {
		t.Errorf("expected OrderedMap.Marshal to be %q, got %q", expect, data)
	}
}
