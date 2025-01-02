package core

import (
	"cmp"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
)

func ConvertStringTo[T cmp.Ordered | bool](v string, t *T) {
	typ := reflect.TypeOf(t)
	val := reflect.ValueOf(t)
	eleType := typ.Elem().Kind()
	switch eleType {
	case reflect.String:
		val.Elem().SetString(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			logrus.Error(err.Error())
		}
		val.Elem().SetInt(i)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			logrus.Error(err.Error())
		}
		val.Elem().SetFloat(f)
	case reflect.Bool:
		b, err := strconv.ParseBool(v)
		if err != nil {
			logrus.Error(err.Error())
		}
		val.Elem().SetBool(b)
	default:
		logrus.Warnf("不支持的数据类型: %s", eleType)
	}
}
