package mixjson

import (
	"encoding/json"
	"strconv"
)

type String string

func (st *String) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = String(strconv.Itoa(v))
	case float64:
		*st = String(strconv.FormatFloat(v, 'f', -1, 64))
	case string:
		*st = String(v)
	}
	return nil
}

type Int int

func (st *Int) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = Int(v)
	case float64:
		*st = Int(int64(v))
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*st = Int(i)
	}
	return nil
}

type Float float64

func (st *Float) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = Float(v)
	case float64:
		*st = Float(v)
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		*st = Float(f)
	}
	return nil
}
