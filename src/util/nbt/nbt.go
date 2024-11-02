package nbt

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Decode(d any) (string, error) {
	if d == nil {
		return "", nil
	}

	switch d.(type) {
	case int8:
		return strconv.FormatInt(int64(d.(int8)), 10) + `b`, nil
	case bool:
		return strconv.FormatBool(d.(bool)), nil
	case int16:
		return strconv.FormatInt(int64(d.(int16)), 10) + `s`, nil
	case int32:
		return strconv.FormatInt(int64(d.(int32)), 10), nil
	case int:
		return strconv.FormatInt(int64(d.(int)), 10), nil
	case int64:
		return strconv.FormatInt(int64(d.(int64)), 10) + `l`, nil
	case float32:
		return strconv.FormatFloat(float64(d.(float32)), 'f', -1, 32) + `f`, nil
	case float64:
		return strconv.FormatFloat(d.(float64), 'f', -1, 64), nil
	case string:
		return fmt.Sprintf("\"%s\"", d.(string)), nil
	case []any:
		vals := make([]string, 0)
		for _, v := range d.([]any) {
			s, err := Decode(v)
			if err != nil {
				return "", err
			}
			if s != "" {
				vals = append(vals, s)
			}
		}
		return "[" + strings.Join(vals, ",") + "]", nil
	case map[string]any:
		vals := make([]string, 0)
		for k, v := range d.(map[string]any) {
			s, err := Decode(v)
			if err != nil {
				return "", err
			}
			if s != "" {
				vals = append(vals, fmt.Sprintf("\"%s\":%s", k, s))
			}
		}
		return "{" + strings.Join(vals, ",") + "}", nil
	case ByteArray:
		result := "[B;"
		vals := make([]string, 0)
		for _, v := range d.(ByteArray) {
			vals = append(vals, strconv.FormatInt(int64(v), 10)+`b`)
		}
		result += strings.Join(vals, ",")
		result += "]"
		return result, nil
	case IntArray:
		result := "[I;"
		vals := make([]string, 0)
		for _, v := range d.(IntArray) {
			vals = append(vals, strconv.FormatInt(int64(v), 10))
		}
		result += strings.Join(vals, ",")
		result += "]"
		return result, nil
	case LongArray:
		result := "[L;"
		vals := make([]string, 0)
		for _, v := range d.(LongArray) {
			vals = append(vals, strconv.FormatInt(int64(v), 10)+`l`)
		}
		result += strings.Join(vals, ",")
		result += "]"
		return result, nil
	case Decodable:
		return d.(Decodable).String(), nil
	}

	return "", UnsupportedTypeError{Err: fmt.Sprintf("decoder.Decode(): Error, type %s of d is not supported and does not implement Decodable", reflect.TypeOf(d).String())}
}

type ByteArray []int8

type IntArray []int32

type LongArray []int64

type Decodable interface {
	String() string
}

type UnsupportedTypeError struct {
	Err string
}

func (e UnsupportedTypeError) Error() string {
	return e.Err
}
