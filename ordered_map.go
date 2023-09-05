package generic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type OrderedMap[K comparable, V any] struct {
	// Backing unordered map
	umap map[K]V

	// Ordered keys
	keys []K

	// Should marshal as list?
	list bool
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		umap: make(map[K]V),
		keys: nil,
	}
}

func (om *OrderedMap[K, V]) AsList() *OrderedMap[K, V] {
	om.list = true
	return om
}

func (om *OrderedMap[K, V]) SortKeys(less func(a, b K) bool) {
	SortFunc(om.keys, less)
}

func (om *OrderedMap[K, V]) Set(key K, value V) {
	if _, ok := om.umap[key]; ok {
		om.umap[key] = value
		return
	}

	om.umap[key] = value
	om.keys = append(om.keys, key)

}

func (om *OrderedMap[K, V]) Get(key K) V {
	return om.umap[key]
}

func (om *OrderedMap[K, V]) GetOk(key K) (V, bool) {
	v, ok := om.umap[key]

	return v, ok
}

func (om *OrderedMap[K, V]) GetOr(key K, default_ V) V {
	if v, ok := om.umap[key]; ok {
		return v
	}

	return default_
}

func (om *OrderedMap[K, V]) Exists(key K) bool {
	_, ok := om.umap[key]
	return ok
}

func (om *OrderedMap[K, V]) ToMap() map[K]V {
	return CopyMap(om.umap)
}

func (om *OrderedMap[K, V]) Delete(key K) V {
	var (
		v   V
		ok  bool
		idx int
	)

	if v, ok = om.umap[key]; !ok {
		return v
	}

	for k, v := range om.keys {
		if v == key {
			idx = k
			break
		}
	}

	copy(om.keys[idx:], om.keys[idx+1:])
	om.keys = om.keys[:len(om.keys)-1]
	delete(om.umap, key)

	return v
}

func (om *OrderedMap[K, V]) Keys() []K {
	return CopySlice(om.keys)
}

func (om *OrderedMap[K, V]) Iter(fn func(K, V)) {
	for _, k := range om.keys {
		fn(k, om.umap[k])
	}
}

func (om *OrderedMap[K, V]) IterBreak(fn func(K, V) bool) {
	for _, k := range om.keys {
		if !fn(k, om.umap[k]) {
			break
		}
	}
}

func (om *OrderedMap[K, V]) String() string {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("ordered-map[")
	for i, key := range om.keys {
		if i != 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(buf, "%+v:%+v", key, om.umap[key])
	}
	buf.WriteString("]")

	return buf.String()
}

func (om *OrderedMap[K, V]) MarshalJSON() ([]byte, error) {
	var (
		res = bytes.NewBuffer(nil)
		enc = json.NewEncoder(res)
	)

	if om.list {
		res.WriteByte('[')
		for i, k := range om.keys {
			if i > 0 {
				res.WriteByte(',')
			}

			res.WriteByte('[')

			if err := enc.Encode(k); err != nil {
				return nil, err
			}

			res.WriteByte(',')

			if err := enc.Encode(om.umap[k]); err != nil {
				return nil, err
			}

			res.WriteByte(']')
		}
		res.WriteByte(']')
	} else {
		res.WriteByte('{')
		for i, k := range om.keys {
			if i > 0 {
				res.WriteByte(',')
			}

			if err := enc.Encode(k); err != nil {
				return nil, err
			}

			res.WriteByte(':')

			if err := enc.Encode(om.umap[k]); err != nil {
				return nil, err
			}
		}
		res.WriteByte('}')
	}

	return res.Bytes(), nil
}

var ErrUnmarshalError = errors.New("unmarshal error")

// not the prettiest solution ever
func consume(data *[]byte, dec **json.Decoder, until ...byte) (byte, error) {
	var cur int64
	if dec != nil {
		cur = (*dec).InputOffset()
	}
	*data = (*data)[cur:]

	for i, chr := range *data {
		for _, expected := range until {
			if chr != expected {
				continue
			}

			*data = (*data)[i+1:]
			if dec != nil {
				*dec = json.NewDecoder(bytes.NewReader(*data))
			}
			return chr, nil
		}

		switch chr {
		case ' ', '\t', '\n', '\r', '\v':
		default:
			return 0, &json.SyntaxError{}
		}
	}

	return 0, &json.SyntaxError{}
}

func (om *OrderedMap[K, V]) UnmarshalJSON(b []byte) error {
	chr, err := consume(&b, nil, '{', '[')
	if err != nil {
		return err
	}

	if om.umap == nil || len(om.umap) > 0 {
		om.umap = make(map[K]V)
		om.keys = nil
	}

	if chr == '{' {
		return om.unmarshalMap(b)
	}

	return om.unmarshalList(b)
}

func (om *OrderedMap[K, V]) unmarshalMap(b []byte) error {
	var (
		key K
		val V
		dec = json.NewDecoder(bytes.NewReader(b))
	)

	for {
		if err := dec.Decode(&key); err != nil {
			return err
		}

		if _, err := consume(&b, &dec, ':'); err != nil {
			return err
		}

		if err := dec.Decode(&val); err != nil {
			return err
		}

		om.Set(key, val)

		if c, err := consume(&b, &dec, ',', '}'); err != nil {
			return err
		} else if c == '}' {
			break
		}
	}

	return nil
}

func (om *OrderedMap[K, V]) unmarshalList(b []byte) error {
	var (
		key K
		val V
		dec = json.NewDecoder(bytes.NewReader(b))
	)

	for {
		if _, err := consume(&b, &dec, '['); err != nil {
			return err
		}

		if err := dec.Decode(&key); err != nil {
			return err
		}

		if _, err := consume(&b, &dec, ','); err != nil {
			return err
		}

		if err := dec.Decode(&val); err != nil {
			return err
		}

		om.Set(key, val)

		if _, err := consume(&b, &dec, ']'); err != nil {
			return err
		}

		if c, err := consume(&b, &dec, ',', ']'); err != nil {
			return err
		} else if c == ']' {
			break
		}
	}

	return nil
}
