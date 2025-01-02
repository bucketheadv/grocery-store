package mixjson

import (
	"encoding/json"
	"testing"
)

type StringStruct struct {
	// String 类型支持同时解析json中的字段串和数字类型
	Value String `json:"value"`
}

func TestParseString(t *testing.T) {
	var val = `{"value": 1}`
	st := StringStruct{}
	err := json.Unmarshal([]byte(val), &st)
	if err != nil {
		t.Errorf("解析异常, %s", err)
	}
}

type IntStruct struct {
	Value Int `json:"value"`
}

func TestParseInt(t *testing.T) {
	var val = `{"value": "1"}`
	st := IntStruct{}
	err := json.Unmarshal([]byte(val), &st)
	if err != nil {
		t.Errorf("解析异常, %s", err)
	}
}

type FloatStruct struct {
	Value Float `json:"value"`
}

func TestParseFloat(t *testing.T) {
	var val = `{"value": "1.0"}`
	st := FloatStruct{}
	err := json.Unmarshal([]byte(val), &st)
	if err != nil {
		t.Errorf("解析异常, %s", err)
	}
}
