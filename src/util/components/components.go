package components

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ValiantChip/gopck/src/util/nbt"
	"github.com/ValiantChip/gopck/src/util/parsing"
)

func Parse(d []struct {
	Key   string
	Value any
}) (string, error) {
	vals := make([]string, 0)
	for _, v := range d {
		s, err := ParseValue(v.Value)
		if err != nil {
			return "", err
		}
		if s != "" {
			vals = append(vals, fmt.Sprintf("%s=%s", v.Key, s))
		}
	}

	return fmt.Sprintf("[%s]", strings.Join(vals, ",")), nil
}

func ParseValue(d any) (string, error) {
	switch d.(type) {
	case bool:
		return strconv.FormatBool(d.(bool)), nil
	case float64:
		return strconv.FormatFloat(d.(float64), 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(d.(float32)), 'f', -1, 32), nil
	case int32:
		return strconv.FormatInt(int64(d.(int32)), 10), nil
	case int:
		return strconv.FormatInt(int64(d.(int)), 10), nil
	case int64:
		return strconv.FormatInt(int64(d.(int64)), 10), nil
	case string:
		return fmt.Sprintf("\"%s\"", d.(string)), nil
	case []any:
		vals := make([]string, 0)
		for _, v := range d.([]any) {
			s, err := ParseValue(v)
			if err != nil {
				return "", err
			}
			if s != "" {
				vals = append(vals, s)
			}
		}

		return fmt.Sprintf("[%s]", strings.Join(vals, ",")), nil
	case map[string]any:
		return nbt.Parse(d.(map[string]any))
	case parsing.Parsable:
		return d.(parsing.Parsable).String(), nil
	}

	return "", fmt.Errorf("components.ParseValue: Unsupported type %T", d)
}
